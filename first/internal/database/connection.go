package database

import (
	"fmt"
	"log"
	"module/internal/models"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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

	// открытие соединения. вызывать опен нужно только один раз. он и пингует и соединяется с базой
	db, err := gorm.Open(postgres.Open(connectStr), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	// установка пула соединений. без пула начинает пропускать если больше 10 запросов одновременно. с пулом обрабатывает больше 100 запросов одновременно
	// можно в одном запросе делать несколько действий в горутинах и это будет работать
	sqlDB, err := db.DB()
	if err != nil {
		// control error
	}

	// Find в gorm работает очень медленно, и при куче find запросов будут просадки. при просадках менять find на raw sql query
	// idle connection Открываются при работе на N время. а во время запросов на милисекунду могут открываться ещё и временные подключения.
	// если idle уже открыты то запросы выполняются очень быстро, но временные подключения так и будут создаваться, поэтому их лучше делать +2-4 от количества idle запросов

	// если не справляется старое количество подключений ( например 5000 тяжёлых запросов в момент времени) просадки до секунды при создании новых соединений
	// sqlDB.SetMaxIdleConns(10)
	// sqlDB.SetMaxOpenConns(15)
	// sqlDB.SetConnMaxIdleTime(time.Minute * 10)
	// sqlDB.SetConnMaxLifetime(time.Minute * 10)

	// вроде самые эффективные значения для небольшого количества запросов (меньше 1000 тяжёлых запросов в момент времени) небольшие просадки при создании новых соединений
	sqlDB.SetMaxIdleConns(4)
	sqlDB.SetMaxOpenConns(6)
	sqlDB.SetConnMaxLifetime(time.Minute)

	sqlDB.Ping()

	// миграция
	db.AutoMigrate(models.Test_Car{})

	return db
}
