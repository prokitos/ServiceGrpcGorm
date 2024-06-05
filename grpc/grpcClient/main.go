package main

import (
	"log"
	server "module/internal"
)

func main() {

	log.Println("Client running ...")

	server.MainServer()

}
