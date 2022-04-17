package main

import (
	"belajar-go-echo/controller"
	"belajar-go-echo/model"
	"belajar-go-echo/service"
	"os"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	connectionString := os.Getenv("APP_DB_CONNECTION_STRING")
	if connectionString == "" {
		connectionString = "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	e := echo.New()
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		e.Logger.Fatal(err)
	}
	db.AutoMigrate(&model.User{})

	ps := service.NewDBUserService(db)
	pc := controller.NewUserController(ps)
	e.GET("/users", pc.Get)  // mendapatkan users
	e.POST("/users", pc.Add) // menambah users
	e.Logger.Fatal(e.Start(":" + port))
}
