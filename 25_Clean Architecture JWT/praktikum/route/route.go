package route

import (
	"belajar-go-echo/config"
	"belajar-go-echo/controller"
	"belajar-go-echo/middleware"

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
	app.GET("/users", controller.GetAllUsers(db), middleware.JWT([]byte(middleware.secretJwt)))
	app.POST("/users", controller.CreateUser(db), middleware.JWT([]byte(middleware.secretJwt)))
	return app
}
