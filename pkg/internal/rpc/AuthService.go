package rpc

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/PostService/pkg/internal/rpc/schema"
	"github.com/PostService/pkg/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"strings"
	"time"
)

func VerifyToken(token string, userAgent string) error {
	parsedToken := strings.Split(token, " ")[1]
	var (
		port = utils.GetValue("PORT_AUTH_RPG")
		host = utils.GetValue("HOST_AUTH_RPG")
	)
	conn, err := grpc.NewClient(fmt.Sprintf("%v:%v", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf(err.Error())
		}
	}(conn)

	client := pb.NewAuthorizationServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := client.VerifyToken(ctx, &pb.VerifyTokenRequest{
		Token: parsedToken,
		Id:    userAgent,
	})

	if err != nil {
		return err
	}

	if !res.IsValid {
		return errors.New(res.Message)
	}

	return nil
}
