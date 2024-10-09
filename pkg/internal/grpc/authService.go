package grpc

import (
	"context"
	"errors"
	pb "github.com/SocService/pkg/internal/grpc/schema"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func CheckAuth(id string, token string) error {
	//clientOpts := alts.DefaultClientOptions()
	//clientOpts.TargetServiceAccounts = []string{expectedServerSA}

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Printf(err.Error())
		}
	}(conn)

	client := pb.NewAuthorizationServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.VerifyToken(ctx, &pb.VerifyTokenRequest{Id: id, Token: token})
	if err != nil {
		return err
	}

	if !res.IsValid {
		return errors.New(res.Message)
	}

	return nil
}
