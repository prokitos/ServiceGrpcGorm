package app

import (
	"fmt"
	server "module/internal/grpc"
	"net"

	"google.golang.org/grpc"
)

type App struct {
	GRPCserver *grpc.Server
}

func (a *App) NewServer() {
	gRPCserver := grpc.NewServer()
	server.Register(gRPCserver)
	a.GRPCserver = gRPCserver
}

func (a *App) Run() error {

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 888))
	if err != nil {
		return fmt.Errorf("%s: %w", "run", err)
	}

	if err := a.GRPCserver.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", "run", err)
	}

	return nil
}

func (a *App) Stop() {
	a.GRPCserver.GracefulStop()
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}
