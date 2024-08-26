package rpc

import (
	"context"
	"errors"
	pb "github.com/AuthService/pkg/internal/rpc/schema"
	modelhttp "github.com/AuthService/pkg/models/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func CreateNode(user modelhttp.RegisterResponse) error {
	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}(conn)

	client := pb.NewAuth2SocServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := client.CreateNode(ctx, &pb.MakeNodeRequest{
		Id:   user.Id,
		Name: user.UserName,
	})

	if err != nil {
		return err
	}

	if !res.Success {
		return errors.New(res.Message)
	}

	return nil
}
