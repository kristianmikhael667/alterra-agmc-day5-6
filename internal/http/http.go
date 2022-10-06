package http

import (
	"main/internal/app/auth"
	"main/internal/app/users"
	"main/internal/factory"
	"main/pkg/util"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	e.Validator = &util.CustomValidator{
		Validator: validator.New(),
	}

	e.GET("/status", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "OK"})
	})

	v1 := e.Group("/api/v1")
	users.NewHandler(f).Route(v1.Group("/users"))
	auth.NewHandler(f).Route(v1.Group("/auth"))
}
