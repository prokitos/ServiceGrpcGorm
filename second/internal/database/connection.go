package database

import (
	"database/sql"
	"fmt"
	"module/internal/models"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// путь до .env миграций
var envConnvertion string = "internal/config/postgres.env"
var migrationRoute string = "internal/database/migrations"

// хранение соединения с базой данных
var curDbRef *sql.DB

// функция для получения базы данных из переменной
func GetConnection() *sql.DB {
	return curDbRef
}

// установка соединения с базой данных
func ConnectToDb() {

	godotenv.Load(envConnvertion)

	envUser := os.Getenv("User")
	envPass := os.Getenv("Pass")
	envHost := os.Getenv("Host")
	envPort := os.Getenv("Port")
	envName := os.Getenv("Name")

	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", envUser, envPass, envHost, envPort, envName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Error("database connection error")
		log.Debug("there is not connection with database")
		models.CheckError(err)
	}

	db.Begin()

	curDbRef = db
}

func CloseConnectToDb() {
	curDbRef.Close()
}

// получить адрес внешнего сервера
func GetExternalRoutes(address *string) {
	godotenv.Load(envConnvertion)
	*address = os.Getenv("ExtAddress")
}

// начать миграцию
func MigrateStart() {
	db := Init()
	GlobalHandler = New(db)
}

var GlobalHandler Handler

type Handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) Handler {
	return Handler{db}
}

func Init() *gorm.DB {
	dbURL := "postgres://postgres:root@localhost:8092/postgres"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(models.Carer{})

	return db
}
