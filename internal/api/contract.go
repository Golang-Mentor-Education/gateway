package api

import "github.com/Golang-Mentor-Education/gateway/internal/model"

type Srv interface {
	SayHello(data *model.In)
}
