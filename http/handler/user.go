package handler

import (
	"context"
	"fmt"
	pb "github.com/ArtGooner/go-project/user/proto"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"time"
)

type UserHandler struct {
	cl pb.userServiceClient
}

func New(cl pb.userServiceClient) *UserHandler {
	return &UserHandler{cl: cl}
}

func (u UserHandler) Login(c echo.Context) error {
	email := c.QueryParam("email")
	password := c.QueryParam("password")

	acc := &pb.Account{Email: email, Password: password}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Authenticate(ctx, acc)

	if err != nil {
		log.Fatalf("auth error: %v", err)
	}

	log.Printf("Authenticated: %v", r)

	return c.String(http.StatusOK, fmt.Sprintf("Authenticated: %v", r))
}
