package main

import (
	"context"
	"fmt"
	"log"
	"module/internal/generated"
	"time"

	"google.golang.org/grpc"
)

func main() {

	log.Println("Client running ...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := generated.NewSellerClient(conn)
	response, err := client.Register(ctx, &generated.RegisterRequest{Email: "tyqw", Password: "1234"})
	if err != nil {
		fmt.Println("too long. context time expired. more than 1 second.")
		//panic(err)
		return
	}

	fmt.Println(response.UserId)

}
