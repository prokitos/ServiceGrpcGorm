package server

import (
	"context"
	"fmt"
	"module/internal/generated"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

func LoginRoute(c *fiber.Ctx) error {

	email := c.Query("email", "")
	pass := c.Query("password", "")

	return loginSend(c, email, pass)
}

func RegisterRoute(c *fiber.Ctx) error {

	email := c.Query("email", "")
	pass := c.Query("password", "")

	return registerSend(c, email, pass)
}

func loginSend(c *fiber.Ctx, email string, pass string) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		return c.SendString("connecting error")
	}
	defer conn.Close()
	client := generated.NewSellerClient(conn)

	response, err := client.Login(ctx, &generated.LoginRequest{Email: email, Password: pass})
	if err != nil {
		fmt.Println("too long. context time expired. more than 1 second.")
		return c.SendString("long execution")
	}
	fmt.Println(response.Response)

	return c.SendString(response.Response)
}

func registerSend(c *fiber.Ctx, email string, pass string) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		return c.SendString("connecting error")
	}
	defer conn.Close()
	client := generated.NewSellerClient(conn)

	response, err := client.Register(ctx, &generated.RegisterRequest{Email: email, Password: pass})
	if err != nil {
		fmt.Println("too long. context time expired. more than 1 second.")
		return c.SendString("long execution")
	}
	fmt.Println(response.UserId)

	return c.SendString(strconv.Itoa(int(response.UserId)))
}
