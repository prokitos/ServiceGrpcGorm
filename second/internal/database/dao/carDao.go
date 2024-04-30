package dao

import (
	"encoding/json"
	"fmt"
	"module/internal/database"
	"module/internal/models"
	"net/http"
	"strconv"
)

// Обновление записи
func UpdateData(w *http.ResponseWriter, curModel *models.Test_Car) {

	var curCar models.Test_Car
	id := curModel.Car_id

	if result := database.GlobalHandler.DB.First(&curCar, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	curCar.Mark = curModel.Mark
	curCar.Model = curModel.Model
	curCar.RegNum = curModel.RegNum
	curCar.Year = curModel.Year

	database.GlobalHandler.DB.Save(&curCar)

	// Send a 201 created response
	json.NewEncoder(*w).Encode("Updated")
}

// Показать записи
func ShowData(w *http.ResponseWriter, curModel *models.Test_Car, curSettings *models.Searcher) {

	var finded []models.Test_Car

	tempConstructor := database.GlobalHandler.DB

	// добавление оффсета и лимита если указан
	limit, err := strconv.Atoi(curSettings.Limit)
	if err == nil {
		tempConstructor = tempConstructor.Limit(limit)
		fmt.Println("offset " + curSettings.Offset)
	}
	offset, err := strconv.Atoi(curSettings.Offset)
	if err == nil {
		tempConstructor = tempConstructor.Offset(offset)
		fmt.Println("limit " + curSettings.Limit)
	}

	// добавление сортировки
	//tempConstructor.Order("name")

	// подгрузка зависимостей. показывать машины и их владельцев
	tempConstructor = tempConstructor.Preload("Owner")

	if result := tempConstructor.Find(&finded, curModel); result.Error != nil {
		fmt.Println(result.Error)
	}

	json.NewEncoder(*w).Encode(finded)
}

// Создать новую запись
func CreateData(w *http.ResponseWriter, curCar *models.Test_Car) {

	// покачто тестово сам добавляю владельцев
	var curOnwer []models.Test_Owner
	owner1 := models.Test_Owner{}
	owner1.Name = "johan"
	owner1.Surname = "newbies"
	owner2 := models.Test_Owner{}
	owner2.Name = "anton"
	owner2.Surname = "gimov"
	curOnwer = append(curOnwer, owner1)
	curOnwer = append(curOnwer, owner2)
	curCar.Owner = curOnwer

	// Append to the Books table
	if result := database.GlobalHandler.DB.Create(&curCar); result.Error != nil {
		fmt.Println(result.Error)
		return
	}

	// Send a 201 created response
	json.NewEncoder(*w).Encode("Created")

}

// Удаление записи по айди
func DeleteData(w *http.ResponseWriter, id string) {

	var curCar models.Test_Car

	if result := database.GlobalHandler.DB.First(&curCar, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Delete that book
	database.GlobalHandler.DB.Delete(&curCar)

	// Send a 201 created response
	json.NewEncoder(*w).Encode("Deleted")
}
