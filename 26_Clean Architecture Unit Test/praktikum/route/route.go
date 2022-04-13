package route

import (
	"belajar-go-echo/config"
	"belajar-go-echo/controller"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	// Route / to handler function
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	err = config.MigrateDB(db)
	if err != nil {
		panic(err)
	}

	app := echo.New()
	app.GET("/users", controller.GetAllUsers(db))
	app.POST("/users", controller.CreateUser(db))
	return app
}
