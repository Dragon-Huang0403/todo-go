package handler

import (
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestHealthCheck(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		m := setup(t)
		// prepare
		c, rec := m.prepareContext(nil)

		// assert
		err := m.handler.HealthCheck()(c)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, rec.Code)

		contentType := rec.Header().Get(echo.HeaderContentType)
		require.Equal(t, echo.MIMEApplicationJSON, contentType)

		require.JSONEq(t, `{"status":"OK"}`, rec.Body.String())
	})
}
