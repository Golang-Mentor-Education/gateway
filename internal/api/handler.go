package api

import (
	"fmt"
	"github.com/Golang-Mentor-Education/gateway/internal/model"
)

type Handler struct {
	srv Srv
}

func NewHandler(srv Srv) *Handler {
	return &Handler{srv: srv}
}

func (h *Handler) Do(in string, out string) error {
	data := &model.In{
		InString:  in,
		OutString: out,
	}
	h.srv.SayHello(data)
	fmt.Println(data.Result)
	return nil
}
