package auth

import (
	"context"
	"fmt"
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

func (c *Client) Login(username, password, email string) (string, error) {
	resp, err := c.client.Login(context.Background(), &auth.LoginIn{
		Username: username,
		Password: password,
		Email:    email,
	})
	if err != nil {
		return "", fmt.Errorf("failed to login: %w", err)
	}
	return resp.Token, nil
}

func (c *Client) Signup(username, email, password string) error {
	_, err := c.client.Signup(context.Background(), &auth.SignupIn{
		Username: username,
		Email:    email,
		Password: password,
	})
	if err != nil {
		return fmt.Errorf("failed to signup: %w", err)
	}
	return nil
}
