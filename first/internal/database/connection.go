package database

import (
	"fmt"
	"module/internal/models"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// путь до env файла
var envConnvertion string = "internal/config/postgres.env"

// начать миграцию
func MigrateStart() {
	db := Init()
	GlobalHandler = New(db)
}

// глобальный открытый коннект
var GlobalHandler Handler

type Handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) Handler {
	return Handler{db}
}

func Init() *gorm.DB {

	// получение строки соединения с бд из env
	godotenv.Load(envConnvertion)
	envUser := os.Getenv("User")
	envPass := os.Getenv("Pass")
	envHost := os.Getenv("Host")
	envPort := os.Getenv("Port")
	envName := os.Getenv("Name")
	connectStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", envUser, envPass, envHost, envPort, envName)

	// открытие соединения
	db, err := gorm.Open(postgres.Open(connectStr), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	// установка пула соединений. без пула начинает пропускать если больше 10 запросов одновременно. с пулом обрабатывает больше 100 запросов одновременно
	sqlDB, err := db.DB()
	if err != nil {
		// control error
	}

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(8)
	sqlDB.SetConnMaxLifetime(time.Minute)

	// миграция
	db.AutoMigrate(models.Test_Car{})

	return db
}
