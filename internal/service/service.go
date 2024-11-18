package service

import (
	"fmt"
	"github.com/Golang-Mentor-Education/gateway/internal/model"
)

type Service struct {
	dbR DbRepo
	c   ClientC
}

func NewService(dbR DbRepo, c ClientC) *Service {
	return &Service{dbR: dbR, c: c}
}

func (s *Service) SayHello(data *model.In) {
	err := s.dbR.SaveToDB(data.InString, data.OutString)
	if err != nil {
		fmt.Println("save to db error:", err)
	}
	result, err := s.c.SendMessage("")
	if err != nil {
		fmt.Println("send message error:", err)
	}
	data.Result = fmt.Sprintf("hello numbers: %s, %s = %s", data.InString, data.OutString, result)
}
