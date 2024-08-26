package rpc

import (
	"context"
	"fmt"
	pb "github.com/AuthService/pkg/handlers/rpc/schema"
	"github.com/AuthService/utils"
	"google.golang.org/grpc"
	"log"
	"net"
)

type AuthHandler struct {
	pb.UnimplementedAuthorizationServiceServer
	pb.UnimplementedAuthenticationServiceServer
}

func NewAuthHandler() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err.Error())
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterAuthorizationServiceServer(grpcServer, initAuthorizationHandler())

	fmt.Println("start to listen on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err.Error())
	}
}

func initAuthorizationHandler() *AuthHandler {
	return &AuthHandler{}
}

func (s *AuthHandler) VerifyToken(ctx context.Context, in *pb.VerifyTokenRequest) (*pb.TokenResponse, error) {
	if len(in.Token) == 0 || len(in.Id) == 0 {
		return &pb.TokenResponse{IsValid: false, Message: "token or id is invalid!"}, nil
	}
	err := utils.ValidateTokenWithId(in.Token, in.Id)
	if err != nil {
		return &pb.TokenResponse{IsValid: false, Message: err.Error()}, nil
	}
	return &pb.TokenResponse{IsValid: true, Message: "success"}, nil
}

func (s *AuthHandler) IsValidToken(ctx context.Context, in *pb.TokenRequest) (*pb.TokenResponse, error) {
	if len(in.Token) == 0 {
		return &pb.TokenResponse{IsValid: false, Message: "token or id is invalid!"}, nil
	}
	err := utils.IsValidToken(in.Token)
	if err != nil {
		return &pb.TokenResponse{IsValid: false, Message: err.Error()}, nil
	}
	return &pb.TokenResponse{IsValid: true, Message: "success"}, nil
}
