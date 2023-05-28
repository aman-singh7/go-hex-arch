package grpc

import (
	"log"
	"net"

	"github.com/aman-singh7/go-hex-arch/internal/adapters/framework/driver/grpc/pb"
	"github.com/aman-singh7/go-hex-arch/internal/ports"
	"google.golang.org/grpc"
)

type GRPCAdapter struct {
	api ports.APIPort
}

func NewAdapter(api ports.APIPort) *GRPCAdapter {
	return &GRPCAdapter{
		api: api,
	}
}

func (grpca *GRPCAdapter) Run() {
	var err error

	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v", err)
	}

	arithmeticServiceServer := grpca
	grpcServer := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(grpcServer, arithmeticServiceServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRPC server over port 9000: %v", err)
	}
}
