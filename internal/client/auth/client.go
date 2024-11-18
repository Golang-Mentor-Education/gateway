package auth

import (
	"context"
	"github.com/Golang-Mentor-Education/auth/pkg/auth"
	"google.golang.org/grpc"
)

type Client struct {
	client auth.AuthServiceClient
}

func New() *Client {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := auth.NewAuthServiceClient(conn)
	return &Client{client: client}
}

func (c *Client) DoLogin() string {
	token, err := c.client.Login(context.Background(), &auth.LoginIn{
		Username: "qwe",
		Password: "123",
	})
	if err != nil {
		panic(err)
	}
	return token.Token
}
