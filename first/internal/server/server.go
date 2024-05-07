package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// запуск сервера. localhost:8888
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

// все пути
func routers() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/delete", deleteRequest).Methods(http.MethodDelete)
	router.HandleFunc("/insert", insertRequest).Methods(http.MethodPost)
	router.HandleFunc("/update", updateRequest).Methods(http.MethodPut)
	router.HandleFunc("/show", showsSpecRequest).Methods(http.MethodGet)

	// стресс тест и проверка транзакций
	router.HandleFunc("/stress", stressRequest).Methods(http.MethodGet)
	router.HandleFunc("/test", insertTransaction).Methods(http.MethodPost)

	return router
}
