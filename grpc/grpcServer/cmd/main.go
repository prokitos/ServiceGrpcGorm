package main

import (
	"fmt"
	"module/internal/app"
	"module/internal/config"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// go run cmd/main.go --config=config/local.yaml
	cfg := config.MustLoad()
	fmt.Println(cfg)

	var application app.App
	application.NewServer(cfg.GRPC.Port)
	go application.MustRun() // запуск сервера в горутине, чтобы потом нормально звершать приложение

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	application.Stop() // безопасное выключение сервера

	// проверить есть ли получение пути через строку без флагов. 1:27
	// https://www.youtube.com/watch?v=EURjTg5fw-E

}
