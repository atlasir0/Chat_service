package main

import (
	"context"
	"log"
	"time"

	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/protobuf/types/known/timestamppb"

	desc "github.com/atlasir0/Chat_service/Auth_chat/pkg/note_v1"
)

const (
	address = "localhost:50051"
)

type server struct {
	desc.UnimplementedUserServiceServer
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("User id: %d", req.GetId())

	return &desc.GetResponse{
		Info: &desc.UserInfo{
			Id:        req.GetId(),
			Name:      gofakeit.Name(),
			Email:     gofakeit.Email(),
			Role:      desc.UserRole(gofakeit.Int32()),
			CreatedAt: timestamppb.New(gofakeit.Date()),
			UpdatedAt: timestamppb.New(gofakeit.Date()),
		},
	}, nil
}

func main() {
	creds, err := credentials.NewClientTLSFromFile("./certificate/service.pem", "")
	if err != nil {
		log.Fatalf("failed to load TLS certificate: %v", err)
	}
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := desc.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Get(ctx, &desc.GetRequest{Id: 5})
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}
	log.Printf("User: %v", r.GetInfo())
}
