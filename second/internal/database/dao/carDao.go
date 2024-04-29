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
func UpdateData(w *http.ResponseWriter, curModel *models.Carer) {

	var curCar models.Carer
	id := curModel.Id

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
func ShowData(w *http.ResponseWriter, curModel *models.Carer, curSettings *models.Searcher) {

	var finded []models.Carer

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

	if result := tempConstructor.Find(&finded, curModel); result.Error != nil {
		fmt.Println(result.Error)
	}

	json.NewEncoder(*w).Encode(finded)
}

// Создать новую запись
func CreateData(w *http.ResponseWriter, curCar *models.Carer) {

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

	var curCar models.Carer

	if result := database.GlobalHandler.DB.First(&curCar, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Delete that book
	database.GlobalHandler.DB.Delete(&curCar)

	// Send a 201 created response
	json.NewEncoder(*w).Encode("Deleted")
}
