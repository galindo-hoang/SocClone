package rpc

import (
	"context"
	"fmt"
	pb "github.com/SocService/pkg/handlers/rpc/schema"
	repository "github.com/SocService/pkg/repositories/impl"
	"github.com/SocService/pkg/service"
	"github.com/SocService/pkg/service/impl"
	"github.com/SocService/pkg/service/models"
	"google.golang.org/grpc"
	"log"
	"net"
)

type ServerHandler struct {
	pb.UnimplementedAuth2SocServiceServer
}

func NewServerHandler() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen on port 50052: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterAuth2SocServiceServer(grpcServer, initServer())
	fmt.Println("start to listen on port 50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

func initServer() *ServerHandler {
	return &ServerHandler{}
}

func (s *ServerHandler) CreateNode(ctx context.Context, person *pb.MakeNodeRequest) (*pb.MakeNodeResponse, error) {
	var nodeRequest = models.PersonDto{
		Id:    person.Id,
		Name:  person.Name,
		Image: person.Image,
	}
	var socService service.ISocService = impl.NewSocService(&repository.SocRepository{})

	if err := socService.CreateUser(nodeRequest); err != nil {
		return &pb.MakeNodeResponse{
			Success: false,
			Message: err.Error(),
		}, err
	}

	return &pb.MakeNodeResponse{
		Success: true,
		Message: "",
	}, nil
}
