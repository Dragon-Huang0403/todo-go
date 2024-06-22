package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary		Health Check
// @Description	Health Check
// @Tags			Health
// @Accept			json
// @Produce		json
// @Success		200 {object} handler.HealthCheck.response "OK"
// @Router			/health [get]
func (Handler) HealthCheck() echo.HandlerFunc {
	type response struct {
		Status string `json:"status" example:"OK" validate:"required"`
	}
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, response{Status: "OK"})
	}
}
