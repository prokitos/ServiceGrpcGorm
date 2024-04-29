package server

import (
	"module/internal/services"
	"net/http"
)

func showsSpecGetRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// здесь провекра доступа

	services.CarShow(&w, r)
}
