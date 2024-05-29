package services

import (
	"module/internal/database/dao"
	"module/internal/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

// Удаление записи по айди
func CarDelete(c *fiber.Ctx) error {

	// получение айди, и если айди не может конвертится в число, то выдаем ошибку
	id := c.Query("id", "0")
	if _, err := strconv.Atoi(id); err != nil {
		log.Debug("id couldn't convert to a number: " + id)
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "Dont convert id into number"})
	}

	// запрос к базе данных
	return dao.DeleteData(c, id)
}

// Создать новую запись
func CarCreate(c *fiber.Ctx) error {

	var curCar models.Test_Car
	if err := c.BodyParser(&curCar); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	// временно. получется из внешнего сервиса
	var users []models.Test_Owner
	var user models.Test_Owner
	user.Name = "Anton"
	user.Patronymic = "Igorevich"
	user.Surname = "Pavlov"
	users = append(users, user)
	curCar.Owner = users

	// запрос к базе данных
	return dao.CreateData(c, &curCar)
}

// Показать записи
func CarShow(c *fiber.Ctx) error {

	var curCar models.Test_Car
	tempID := c.Query("id", "0")
	curCar.Car_id, _ = strconv.Atoi(tempID)
	curCar.RegNum = c.Query("regNum", "")
	curCar.Mark = c.Query("mark", "")
	curCar.Model = c.Query("model", "")
	curCar.Year = c.Query("year", "")

	var curSettings models.Searcher
	curSettings.Sort = "0"
	curSettings.Offset = c.Query("offset", "0")
	curSettings.Limit = c.Query("limit", "10")

	// запрос к базе данных
	return dao.ShowData(c, &curCar, &curSettings)
}
