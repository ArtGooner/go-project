package main

import (
	"context"
	"github.com/ArtGooner/go-project/user/config"
	pb "github.com/ArtGooner/go-project/user/proto"
	"github.com/ArtGooner/go-project/user/repository"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) Authenticate(ctx context.Context, in *pb.Account) (*pb.User, error) {
	rps, err := repository.NewRepository()

	if err != nil {
		log.Fatal(err)
	}

	res, err := rps.Get(in)

	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return res, nil
}

func main() {
	cfg, err := config.New()

	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", cfg.GrpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
