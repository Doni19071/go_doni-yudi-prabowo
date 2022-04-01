package main

import (
	"praktikum/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	// create a new echo instance
	e := echo.New()
	// Route / to handler function
	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}

// var (
// 	DB *gorm.DB
// )

// func init() {
// 	InitDB()
// 	InitialMigration()
// }

// type Config struct {
// 	DB_Username string
// 	DB_Password string
// 	DB_Port     string
// 	DB_Host     string
// 	DB_Name     string
// }

// func InitDB() {

// 	config := Config{
// 		DB_Username: "root",
// 		DB_Password: "nokia123",
// 		DB_Port:     "3306",
// 		DB_Host:     "localhost",
// 		DB_Name:     "tugas",
// 	}

// 	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
// 		config.DB_Username,
// 		config.DB_Password,
// 		config.DB_Host,
// 		config.DB_Port,
// 		config.DB_Name,
// 	)

// 	var err error
// 	DB, err = gorm.Open("mysql", connectionString)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// type User struct {
// 	gorm.Model
// 	Name     string `json:"name" form:"name"`
// 	Email    string `json:"email" form:"email"`
// 	Password string `json:"password" form:"password"`
// }

// func InitialMigration() {
// 	DB.AutoMigrate(&User{})
// }

// // get all users
// func GetUsersController(c echo.Context) error {
// 	var users []User

// 	if err := DB.Find(&users).Error; err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success get all users",
// 		"users":   users,
// 	})
// }

// // get user by id
// func GetUserController(c echo.Context) error {
// 	var users []User
// 	id := c.Param("id")
// 	if err := DB.Find(&users).Error; err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}
// 	return c.String(http.StatusOK, id)
// }

// // create new user
// func CreateUserController(c echo.Context) error {
// 	user := User{}
// 	c.Bind(&user)

// 	if err := DB.Save(&user).Error; err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success create new user",
// 		"user":    user,
// 	})
// }

// // delete user by id
// func DeleteUserController(c echo.Context) error {
// 	var users = map[int]*User{}
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	delete(users, id)
// 	return c.NoContent(http.StatusNoContent)
// }

// // update user by id
// func UpdateUserController(c echo.Context) error {
// 	var Users = map[int]*User{}
// 	u := new(User)
// 	if err := c.Bind(u); err != nil {
// 		return err
// 	}
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	Users[id].Name = u.Name
// 	return c.JSON(http.StatusOK, Users[id])
// }
