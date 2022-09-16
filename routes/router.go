package routes

import (
	"main/constants"
	"main/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	// unauth
	e.POST("/users", controllers.CreateUserController)
	// Login
	e.POST("/login", controllers.LoginUserController)
	e.GET("/users", controllers.GetUsersController)
	
	// implementasi middleware with group routing / jwt group
	eAuth := e.Group("/v1")
	// eAuth.Use(middleware.BasicAuth(m.BasicAuthDB))

	eAuth.Use(middleware.JWT([]byte(constants.SecretKey())))
	eAuth.GET("/users", controllers.GetUsersController)
	eAuth.GET("/users/:id", controllers.GetSingleUserController)
	eAuth.PUT("/users/:id", controllers.EditUserController)
	eAuth.DELETE("/users/:id", controllers.DeleteUserController)

	return e
}
