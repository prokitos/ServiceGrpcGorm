package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
)

func main() {

	MainServer()

}

func MainServer() {

	router := routers()

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8888",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}

func routers() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/people", peopleAdd).Methods(http.MethodPost) // добавление пользователя в бд
	router.HandleFunc("/car", carAdd).Methods(http.MethodPost)       // запрос на другой микросервис для верификации машины на пользователя

	return router
}

func peopleAdd(w http.ResponseWriter, r *http.Request) {

	newPerson := Person{}
	newPerson.Name = r.FormValue("name")
	newPerson.Surname = r.FormValue("surname")

	DatabaseProducing(newPerson)

	json.NewEncoder(w).Encode("send: " + newPerson.Name + " " + newPerson.Surname)
}

func carAdd(w http.ResponseWriter, r *http.Request) {

	newPerson := Person{}
	newCar := Car{}
	newPerson.Name = r.FormValue("name")
	newPerson.Surname = r.FormValue("surname")
	newCar.Mark = r.FormValue("mark")
	newCar.RegNum = r.FormValue("regNum")
	newCar.Owner = newPerson

	RegisterProdicing(newCar)

	json.NewEncoder(w).Encode("send: " + newCar.Mark + " and " + newPerson.Name)
}

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}

}

// очередь в database
func DatabaseProducing(TempPerson Person) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	handleError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	amqpChannel, err := conn.Channel()
	handleError(err, "Can't create a amqpChannel")

	defer amqpChannel.Close()

	queue, err := amqpChannel.QueueDeclare("database", true, false, false, false, nil)
	handleError(err, "Could not declare `database` queue")

	rand.Seed(time.Now().UnixNano())

	addTask := TempPerson
	body, err := json.Marshal(addTask)
	if err != nil {
		handleError(err, "Error encoding JSON")
	}

	err = amqpChannel.Publish("", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})

	if err != nil {
		log.Fatalf("Error publishing message: %s", err)
	}

	log.Printf("AddTask: %s %s", addTask.Name, addTask.Surname)
}

// очередь в register
func RegisterProdicing(TempPerson Car) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	handleError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	amqpChannel, err := conn.Channel()
	handleError(err, "Can't create a amqpChannel")

	defer amqpChannel.Close()

	queue, err := amqpChannel.QueueDeclare("register", true, false, false, false, nil)
	handleError(err, "Could not declare `register` queue")

	rand.Seed(time.Now().UnixNano())

	addTask := TempPerson
	body, err := json.Marshal(addTask)
	if err != nil {
		handleError(err, "Error encoding JSON")
	}

	err = amqpChannel.Publish("", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         body,
	})

	if err != nil {
		log.Fatalf("Error publishing message: %s", err)
	}

	log.Printf("AddTask: %s %s", addTask.Mark, addTask.RegNum)
}

type Car struct {
	Mark   string
	RegNum string
	Owner  Person
}

type Person struct {
	Name    string
	Surname string
}
