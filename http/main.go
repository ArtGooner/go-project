package main

import (
	"github.com/ArtGooner/go-project/http/handler"
	"github.com/ArtGooner/go-project/user/config"
	pb "github.com/ArtGooner/go-project/user/proto"
	"github.com/labstack/echo"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "localhost:50051"
	port    = ":8888"
)

func main() {
	cfg, err := config.New()

	if err != nil {
		log.Fatal(err)
	}

	/* grpc Setup */
	conn, err := grpc.Dial(cfg.GrpcAddress, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewUserServiceClient(conn)

	/* Echo Setup*/
	e := echo.New()
	e.HideBanner = true

	uh := handler.New(c)
	e.GET("/Login", uh.Login)

	// has to be fixed, after vendor update
	//err = e.Start(cfg.HttpPort)
	err = e.Start(port)

	if err != nil {
		log.Fatal(err)
	}
}
