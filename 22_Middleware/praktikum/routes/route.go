package routes

import (
	"praktikum/controllers"
	"praktikum/middleware"

	mid "github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	// Route / to handler function
	e := echo.New()
	eJWT := e.Group("/")
	e.POST("/login", controllers.LoginUserController)
	e.POST("/users", controllers.CreateUserController)
	// Need Authentication
	eJWT.Use(mid.JWT([]byte(middleware.SECRET_JWT)))
	eJWT.GET("/users", controllers.GetUsersController)
	eJWT.GET("/users/:id", controllers.GetUserController)
	eJWT.DELETE("/users/:id", controllers.DeleteUserController)
	eJWT.PUT("/users/:id", controllers.UpdateUserController)
	eJWT.POST("/login", controllers.LoginUserController)

	return e
}
