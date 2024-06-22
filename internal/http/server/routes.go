package httpserver

import (
	"github.com/dragon-huang0403/todo-go/internal/http/server/handler"
	"github.com/labstack/echo/v4"
)

func addRoutes(e *echo.Group, h *handler.Handler) {
	e.GET("/health", h.HealthCheck())

	// Task
	task := e.Group("/tasks")
	task.GET("", h.ListTasks())
	task.POST("", h.CreateTask())
	task.PUT("/:taskId", h.UpdateTask())
	task.DELETE("/:taskId", h.DeleteTask())
}
