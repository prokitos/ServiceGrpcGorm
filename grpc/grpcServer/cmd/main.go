package main

import (
	"fmt"
	"module/internal/app"
	"module/internal/config"
)

func main() {

	// go run cmd/main.go --config=config/local.yaml
	cfg := config.MustLoad()
	fmt.Println(cfg)

	var application app.App
	application.NewServer()
	application.MustRun()

	// проверить есть ли получение пути через строку без флагов. 1:09
	// https://www.youtube.com/watch?v=EURjTg5fw-E

}
