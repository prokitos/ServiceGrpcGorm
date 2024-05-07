package server

import (
	"module/internal/services"
	"net/http"
)

func insertTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// здесь провекра доступа

	services.CarTransact(&w, r)
}
