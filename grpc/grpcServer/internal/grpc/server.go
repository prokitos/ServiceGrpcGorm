package grpc

import (
	"context"
	"module/internal/grpc/generated"

	"google.golang.org/grpc"
)

type serverApi struct {
	generated.UnimplementedSellerServer
}

func Register(gRPC *grpc.Server) {
	generated.RegisterSellerServer(gRPC, &serverApi{})
}

func (s *serverApi) Register(ctx context.Context, req *generated.RegisterRequest) (*generated.RegisterResponse, error) {
	panic("not implemented")
}
