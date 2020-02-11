package main

import (
	"context"
	pb "github.com/ArtGooner/go-project/user"
	"github.com/ArtGooner/go-project/user/repository"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) Authenticate(ctx context.Context, in *pb.Account) (*pb.User, error) {
	rps, err := user.NewRepository()

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
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
