package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var secureCode string = "gpt45"

func main() {
	serverStart()
}

// запуск сервера
func serverStart() {

	router := mux.NewRouter()
	router.HandleFunc("/getter", getterRoute).Methods(http.MethodPost)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8112",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
	srv.ListenAndServe()
}

// подключение пути для получения данных с 1 сервера
func getterRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	secret := r.FormValue("key")

	if secret != secureCode {
		log.Info("wrong secure code !!!")
		json.NewEncoder(w).Encode("wrong code !!!!")
		return
	}

	var user send_Owner
	user.Name = "Anton"
	user.Patronymic = "Igorevich"
	user.Surname = "Pavlov"

	json.NewEncoder(w).Encode(user)
}

// структура, с помощью которой идёт обмен внутри сервисов, и потом в базу
type send_Owner struct {
	Owner_id   int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Secret     string `json:"key"`
}
