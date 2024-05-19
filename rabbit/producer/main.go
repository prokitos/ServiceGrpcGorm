package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	gopher_and_rabbit "github.com/masnun/gopher-and-rabbit"
	"github.com/streadway/amqp"
)

func main() {

	MainServer()

}

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}

}

func producing(fnum int, snum int) {
	conn, err := amqp.Dial(gopher_and_rabbit.Config.AMQPConnectionURL)
	handleError(err, "Can't connect to AMQP")
	defer conn.Close()

	amqpChannel, err := conn.Channel()
	handleError(err, "Can't create a amqpChannel")

	defer amqpChannel.Close()

	queue, err := amqpChannel.QueueDeclare("add", true, false, false, false, nil)
	handleError(err, "Could not declare `add` queue")

	rand.Seed(time.Now().UnixNano())

	addTask := gopher_and_rabbit.AddTask{Number1: fnum, Number2: snum}
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

	log.Printf("AddTask: %d+%d", addTask.Number1, addTask.Number2)
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

	router.HandleFunc("/math", testingMath).Methods(http.MethodGet)

	return router
}

func testingMath(w http.ResponseWriter, r *http.Request) {

	fNum := r.FormValue("first")
	sNum := r.FormValue("second")

	first, _ := strconv.Atoi(fNum)
	second, _ := strconv.Atoi(sNum)

	producing(first, second)

	json.NewEncoder(w).Encode("send: " + fNum + " + " + sNum)
}
