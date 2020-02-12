package main

import (
	//"github.com/ArtGooner/go-project/user/config"
	"github.com/ArtGooner/go-project/http/handler"
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
	/* grpc Setup */
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := pb.NewUserServiceClient(conn)

	/* Echo Setup*/
	e := echo.New()
	e.HideBanner = true

	eh := handler.New(c)
	//eh := New(c)
	e.GET("/Login", eh.Login)

	err = e.Start(port)

	if err != nil {
		log.Fatal(err)
	}
}

//type UserHandler struct {
//	cl pb.UserServiceClient
//}
//
//func New(cl pb.UserServiceClient) *UserHandler {
//	return &UserHandler{cl: cl}
//}
//
//func (u UserHandler) Login(c echo.Context) error {
//	email := c.QueryParam("email")
//	password := c.QueryParam("password")
//
//	acc := &pb.Account{Email: email, Password: password}
//
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
//	defer cancel()
//
//	r, err := u.cl.Authenticate(ctx, acc)
//
//	if err != nil {
//		log.Fatalf("auth error: %v", err)
//	}
//
//	log.Printf("Authenticated: %v", r)
//
//	return c.String(http.StatusOK, fmt.Sprintf("Authenticated: %v", r))
//}
