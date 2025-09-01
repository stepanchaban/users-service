package grpc

import (
	"fmt"
	"net"

	userpb "github.com/stepanchaban/project-prot/proto/user"
	"github.com/stepanchaban/users-service/internal/user"
	"google.golang.org/grpc"
)

func RunGRPC(svc user.UserService) error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, NewHandler(svc))

	fmt.Println("gRPC server is running on port 50051")
	return grpcServer.Serve(lis)
}