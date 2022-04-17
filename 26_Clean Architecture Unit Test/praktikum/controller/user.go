package controller

import (
	"belajar-go-echo/model"
	"belajar-go-echo/service"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	ps service.UserService
}

func (pc UserController) Get(c echo.Context) error {
	users, err := pc.ps.Get()
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "cannot get users"})
	}
	return c.JSON(http.StatusOK, users)
}

func (pc UserController) Add(c echo.Context) error {
	var newUser model.User
	c.Bind(&newUser)
	user, err := pc.ps.Add(newUser)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "cannot get users"})
	}
	return c.JSON(http.StatusOK, user)
}

func NewUserController(ps service.UserService) UserController {
	return UserController{
		ps: ps,
	}
}
