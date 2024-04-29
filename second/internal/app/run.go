package app

import (
	"module/internal/database"
	"module/internal/server"

	log "github.com/sirupsen/logrus"
)

func RunApp() {

	// установка логов. установка чтобы показывать логи debug уровня
	log.SetLevel(log.DebugLevel)
	log.Info("the server is starting")

	// подключение к базе данных и закрытие после закрытия приложения или краша.
	database.ConnectToDb()
	defer database.CloseConnectToDb()

	// миграция и подключение к бд
	go database.MigrateStart()

	// запуск сервера. внутри указан порт 8888
	server.MainServer()

}
