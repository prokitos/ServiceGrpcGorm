package services

import (
	"encoding/json"
	"fmt"
	"io"
	"module/internal/database"
	"module/internal/database/dao"
	"module/internal/models"
	"net/http"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

// обновление записи
func CarUpdate(w *http.ResponseWriter, r *http.Request) {

	reqBody, _ := io.ReadAll(r.Body)
	var curCar models.Test_Car
	json.Unmarshal(reqBody, &curCar)

	// если пользователь не ввёл айди изменяемой записи, то ошибка
	if curCar.Car_id < 1 {
		log.Debug("don't correct id of updated car")
		models.BadClientResponse400(w)
		return
	}

	// запрос к базе данных
	dao.UpdateData(w, &curCar)

}

// Удаление записи по айди
func CarDelete(w *http.ResponseWriter, r *http.Request) {

	// получение айди, и если айди не может конвертится в число, то выдаем ошибку
	id := r.FormValue("id")
	if _, err := strconv.Atoi(id); err != nil {
		log.Debug("id couldn't convert to a number: " + id)
		models.BadClientResponse400(w)
		return
	}

	// запрос к базе данных
	dao.DeleteData(w, id)

}

// Показать записи
func CarShow(w *http.ResponseWriter, r *http.Request) {

	var curCar models.Test_Car
	tempID := r.FormValue("id")
	curCar.Car_id, _ = strconv.Atoi(tempID)
	curCar.RegNum = r.FormValue("regNum")
	curCar.Mark = r.FormValue("mark")
	curCar.Model = r.FormValue("model")
	curCar.Year = r.FormValue("year")

	var curSettings models.Searcher
	curSettings.Sort = r.FormValue("sort")
	curSettings.Offset = r.FormValue("offset")
	curSettings.Limit = r.FormValue("limit")

	// проверка что в offset,limit либо пустота либо числа.
	if _, err := strconv.Atoi(curSettings.Offset); curSettings.Offset != "" && err != nil {
		log.Debug("offset params couldn't convert to a number, offset = " + curSettings.Offset)
		models.BadClientResponse400(w)
		return
	}
	if _, err := strconv.Atoi(curSettings.Limit); curSettings.Limit != "" && err != nil {
		log.Debug("limit params couldn't convert to a number, limit = " + curSettings.Limit)
		models.BadClientResponse400(w)
		return
	}

	// запрос к базе данных
	dao.ShowData(w, &curCar, &curSettings)

}

// Создать новую запись
func CarCreate(w *http.ResponseWriter, r *http.Request) {

	reqBody, _ := io.ReadAll(r.Body)
	var curCar models.Test_Car
	json.Unmarshal(reqBody, &curCar)

	// запрос к базе данных
	dao.CreateData(w, &curCar)

}

// Создать новые записи через транзакцию
func CarTransact(w *http.ResponseWriter, r *http.Request) {

	// создание пустой структуры, в дальнейшем заполнится программно
	var curCar models.Test_Car
	CreateDataTransaction(w, &curCar)

}

// Создать новую запись через транзакцию
func CreateDataTransaction(w *http.ResponseWriter, curCar *models.Test_Car) {

	// создание транзакции
	curCon := database.GlobalHandler.DB
	tx := curCon.Begin()

	// обработка нестандартных ошибок. например краш программы. не добавлять ничего если случилось.
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			json.NewEncoder(*w).Encode("Error")
			return
		}
	}()

	// типо делаем пакет запросов (транзакция) если один из запросов не выполнен, то не выполнятся все
	for i := 0; i < 3; i++ {

		var tempCar models.Test_Car
		tempCar.Mark = "marks"
		tempCar.Model = "model s"
		tempCar.RegNum = "xx100xx"
		tempCar.Year = "2000"

		var curOnwer []models.Test_Owner
		owner1 := models.Test_Owner{}
		owner1.Name = "johan"
		owner1.Surname = "newbies"
		curOnwer = append(curOnwer, owner1)
		tempCar.Owner = curOnwer

		// создать здесь типо подключение к внешнему серверу, и получать оттуда данные о владельце.
		// если нет подключения к внешнему серверу то откат
		time.Sleep(time.Second * 1)

		// если нет соединения с базой то не добавлять ничего
		if result := tx.Create(&tempCar); result.Error != nil {
			tx.Rollback()
			fmt.Println(result.Error)
			json.NewEncoder(*w).Encode("Error")
			return
		}

	}

	// создать данные если всё прошло хорошо
	tx.Commit()
	json.NewEncoder(*w).Encode("Created")

}
