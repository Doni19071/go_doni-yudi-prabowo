package main

import (
	"praktikum/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	// create a new echo instance
	e := echo.New()
	// Route / to handler function
	e.GET("/users", controllers.GetBooksController)
	e.GET("/users/:id", controllers.GetBookController)
	e.POST("/users", controllers.CreateBookController)
	e.DELETE("/users/:id", controllers.DeleteBookController)
	e.PUT("/users/:id", controllers.UpdateBookController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
