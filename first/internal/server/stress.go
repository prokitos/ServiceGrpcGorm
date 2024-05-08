package server

import (
	"module/internal/services"
	"net/http"
)

func stressRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// стресс тест. 50000 одновременных запросов
	for i := 0; i < 50000; i++ {
		services.CarShow(&w, r)
	}

}
