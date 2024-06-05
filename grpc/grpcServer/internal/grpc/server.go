package server

import (
	"context"
	"fmt"
	"module/internal/grpc/generated"

	"google.golang.org/grpc"
)

type serverApi struct {
	generated.UnimplementedSellerServer
}

func Register(gRPC *grpc.Server) {
	generated.RegisterSellerServer(gRPC, &serverApi{})
}

// типо регистрация пользователя, выводит id который генерируется по символам пароля и имени.
func (s *serverApi) Register(ctx context.Context, req *generated.RegisterRequest) (*generated.RegisterResponse, error) {

	// попытаюсь получить данные и вывод их в консоль
	var temp string = req.GetEmail() + " " + req.GetPassword()
	fmt.Println(temp)

	// потом какие то вычисления, связь с бд

	var result int = 0
	result = len(temp) * len(req.GetEmail())

	// вывод пользователю его юзер айди. по идее сюда должен возвращаться registerResponse, который заполнится внутри DAO
	return &generated.RegisterResponse{UserId: (int64(result))}, nil

}

// выводит good если правильный пароль и имя, и bad если неправильный
func (s *serverApi) Login(ctx context.Context, req *generated.LoginRequest) (*generated.LoginResponse, error) {

	// проверка данных
	if req.GetEmail() == "admin" && req.GetPassword() == "123" {
		fmt.Println("good")
		return &generated.LoginResponse{Response: "good"}, nil
	} else {
		fmt.Println("bad")
		return &generated.LoginResponse{Response: "bad"}, nil
	}

}
