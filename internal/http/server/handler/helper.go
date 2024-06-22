package handler

import "github.com/labstack/echo/v4"

func bindAndValidate[T any](c echo.Context) (*T, error) {
	t := new(T)
	if err := c.Bind(t); err != nil {
		return nil, err
	}
	if err := c.Validate(t); err != nil {
		return nil, err
	}

	return t, nil
}
