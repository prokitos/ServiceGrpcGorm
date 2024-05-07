package server

import (
	"module/internal/services"
	"net/http"
)

func stressRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// стресс тест. 500 одновременных запросов
	for i := 0; i < 500; i++ {
		services.CarShow(&w, r)
	}

}
