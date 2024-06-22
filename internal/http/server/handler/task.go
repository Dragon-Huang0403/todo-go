package handler

import (
	"errors"
	"net/http"

	"github.com/dragon-huang0403/todo-go/internal/controller"
	"github.com/dragon-huang0403/todo-go/internal/models"
	httpserver "github.com/dragon-huang0403/todo-go/pkg/http/server"
	"github.com/dragon-huang0403/todo-go/pkg/logger"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// @Summary		List Tasks
// @Description	List Tasks
// @Tags			Task
// @Accept			json
// @Produce		json
// @Success		200 {object} handler.ListTasks.response "OK"
// @Router			/tasks [get]
func (h *Handler) ListTasks() echo.HandlerFunc {
	type response struct {
		Data []*models.Task `json:"data" validate:"required"`
	}
	return func(c echo.Context) error {
		ctx := httpserver.TransformContext(c)
		task, err := h.controller.Task.List(ctx)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.ErrInternalServerError)
		}

		return c.JSON(http.StatusOK, response{Data: task})
	}
}

// @Summary		Create Task
// @Description	Create Task
// @Tags			Task
// @Accept			json
// @Produce		json
// @Param			request	body		handler.CreateTask.request	true	"request body"
// @Success		200		{object}	handler.CreateTask.response	"OK"
// @Failure		400		{object}	Failure				"Bad Request"
// @Router			/tasks [post]
func (h *Handler) CreateTask() echo.HandlerFunc {
	type request struct {
		Name   string            `json:"name" validate:"required"`
		Status models.TaskStatus `json:"status" validate:"required,oneof=0 1"`
	}
	type response struct {
		Data models.Task `json:"data" validate:"required"`
	}
	return func(c echo.Context) error {
		ctx := httpserver.TransformContext(c)

		req, err := bindAndValidate[request](c)
		if err != nil {
			logger.Debug(ctx, "failed to bind and validate request", zap.Error(err))
			return c.JSON(http.StatusBadRequest, Failure{Message: err.Error()})
		}

		task, err := h.controller.Task.Create(ctx, controller.CreateTaskParams{
			Name:   req.Name,
			Status: req.Status,
		})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.ErrInternalServerError)
		}

		return c.JSON(http.StatusOK, response{Data: *task})
	}
}

// @Summary		Update Task
// @Description	Update Task
// @Tags			Task
// @Accept			json
// @Produce		json
// @Param			taskId	path		string						true	"task id"
// @Param			request	body		handler.UpdateTask.request	true	"request body"
// @Success		200		{object}	handler.UpdateTask.response	"OK"
// @Failure		400		{object}	Failure				"Bad Request"
// @Failure		404		{object}	Failure				"Not Found"
// @Router			/tasks/{taskId} [put]
func (h *Handler) UpdateTask() echo.HandlerFunc {
	type request struct {
		Name   string            `json:"name" validate:"required"`
		Status models.TaskStatus `json:"status" validate:"required,oneof=0 1"`
	}
	type response struct {
		Data models.Task `json:"data" validate:"required"`
	}
	return func(c echo.Context) error {
		ctx := httpserver.TransformContext(c)

		taskId, err := uuid.Parse(c.Param("taskId"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, Failure{Message: "invalid task id"})
		}

		req, err := bindAndValidate[request](c)
		if err != nil {
			logger.Debug(ctx, "failed to bind and validate request", zap.Error(err))
			return c.JSON(http.StatusBadRequest, Failure{Message: err.Error()})
		}

		task, err := h.controller.Task.Update(ctx, controller.UpdateTaskParams{
			ID:     taskId,
			Name:   req.Name,
			Status: req.Status,
		})
		if err != nil {
			if errors.Is(err, controller.ErrNotFound) {
				return c.JSON(http.StatusNotFound, echo.ErrNotFound)
			}
			return c.JSON(http.StatusInternalServerError, echo.ErrInternalServerError)
		}

		return c.JSON(http.StatusOK, response{Data: *task})
	}
}

// @Summary		Delete Task
// @Description	Delete Task
// @Tags			Task
// @Accept			json
// @Produce		json
// @Param			taskId	path		string						true	"task id"
// @Success		200		{object}	Success	"OK"
// @Failure		400		{object}	Failure				"Bad Request"
// @Failure		404		{object}	Failure				"Not Found"
// @Router			/tasks/{taskId} [delete]
func (h *Handler) DeleteTask() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := httpserver.TransformContext(c)

		taskId, err := uuid.Parse(c.Param("taskId"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, Failure{Message: "invalid task id"})
		}

		err = h.controller.Task.Delete(ctx, taskId)
		if err != nil {
			if errors.Is(err, controller.ErrNotFound) {
				return c.JSON(http.StatusNotFound, echo.ErrNotFound)
			}
			return c.JSON(http.StatusInternalServerError, echo.ErrInternalServerError)
		}

		return c.JSON(http.StatusOK, Success{Success: true})
	}
}
