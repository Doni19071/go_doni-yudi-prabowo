package controllers

import (
	"net/http"
	"praktikum/models"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
)

var (
	DB *gorm.DB
)

// get all users
func GetBooksController(c echo.Context) error {
	var books []models.Book

	if err := DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   books,
	})
}

// get user by id
func GetBookController(c echo.Context) error {
	var books []models.Book
	id := c.Param("id")
	if err := DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, id)
}

// create new user
func CreateBookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)

	if err := DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"book":    book,
	})
}

// delete user by id
func DeleteBookController(c echo.Context) error {
	var books = map[int]*models.Book{}
	id, _ := strconv.Atoi(c.Param("id"))
	delete(books, id)
	return c.NoContent(http.StatusNoContent)
}

// update user by id
func UpdateBookController(c echo.Context) error {
	var Books = map[int]*models.Book{}
	u := new(models.Book)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	Books[id].Name = u.Name
	return c.JSON(http.StatusOK, Books[id])
}
