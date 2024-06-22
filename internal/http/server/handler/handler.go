package handler

import "github.com/dragon-huang0403/todo-go/internal/controller"

type Handler struct {
	controller *controller.Controller
}

func New(controller *controller.Controller) *Handler {
	return &Handler{
		controller: controller,
	}
}

type Failure struct {
	Message string `json:"message" validate:"required"`
}

type Success struct {
	Success bool `json:"success" validate:"required"`
}
