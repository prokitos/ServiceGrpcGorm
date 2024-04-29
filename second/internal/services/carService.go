package services

import (
	"encoding/json"
	"io"
	"module/internal/database/dao"
	"module/internal/models"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// обновление записи
func CarUpdate(w *http.ResponseWriter, r *http.Request) {

	reqBody, _ := io.ReadAll(r.Body)
	var curCar models.Carer
	json.Unmarshal(reqBody, &curCar)

	// если пользователь не ввёл айди изменяемой записи, то ошибка
	if curCar.Id < 1 {
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

	var curCar models.Carer
	tempID := r.FormValue("id")
	curCar.Id, _ = strconv.Atoi(tempID)
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
	var curCar models.Carer
	json.Unmarshal(reqBody, &curCar)

	// запрос к базе данных
	dao.CreateData(w, &curCar)

}
