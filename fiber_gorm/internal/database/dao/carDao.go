package dao

import (
	"module/internal/database"
	"module/internal/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Удаление записи по айди
func DeleteData(c *fiber.Ctx, id string) error {

	var curCar models.Test_Car

	// сначала проверить что запись есть, а потом удалить
	// if result := database.GlobalHandler.DB.First(&curCar, id); result.Error != nil {
	// 	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No note with that Id exists"})
	// }
	// database.GlobalHandler.DB.Delete(&curCar)

	result := database.GlobalHandler.DB.Delete(&curCar, "car_id = ?", id)

	if result.RowsAffected == 0 || result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No note with that Id exists"})
	}

	// Send a 201 created response
	return c.SendStatus(fiber.StatusAccepted)
}

// Показать записи
func ShowData(c *fiber.Ctx, curModel *models.Test_Car, curSettings *models.Searcher) error {

	var finded []models.Test_Car

	// tempConstructor := database.GlobalHandler.DB
	// // добавление оффсета и лимита если указан
	// limit, _ := strconv.Atoi(curSettings.Limit)
	// tempConstructor = tempConstructor.Limit(limit)

	// offset, _ := strconv.Atoi(curSettings.Offset)
	// tempConstructor = tempConstructor.Offset(offset)

	// // подгрузка зависимостей. показывать машины и их владельцев
	// tempConstructor = tempConstructor.Preload("Owner")

	// // найти запись
	// if result := tempConstructor.Find(&finded, curModel); result.Error != nil {
	// 	return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	// }

	offset, _ := strconv.Atoi(curSettings.Offset)
	limit, _ := strconv.Atoi(curSettings.Limit)
	results := database.GlobalHandler.DB.Limit(limit).Offset(offset).Find(&finded, curModel)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(finded), "notes": finded})
}

// Создать новую запись
func CreateData(c *fiber.Ctx, curCar *models.Test_Car) error {

	// создать запись
	if result := database.GlobalHandler.DB.Create(&curCar); result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	// Send a 201 created response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"note": curCar}})
}
