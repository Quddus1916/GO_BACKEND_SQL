package controller

import (
	"github.com/labstack/echo"
	//"github.com/jinzhu/gorm"
	"github.com/Quddus1916/GO_BACKEND_SQL/database"
	"github.com/Quddus1916/GO_BACKEND_SQL/models"
	"net/http"
)

var Users []models.User

func GetUsers(c echo.Context) error {

	Users := models.GetAllUsers()
	if Users == nil {
		msg := "failed to fetch data"
		return c.JSON(http.StatusInternalServerError, msg)
	}

	return c.JSON(http.StatusOK, Users)
}

func CreateUser(c echo.Context) error {
	user := new(models.User)
	err := c.Bind(user)
	if err != nil {
		msg := " failed to bind data"
		return c.JSON(http.StatusBadRequest, msg)
	}
	res := models.CreateUser(user)
	var deferrer = database.GetDB()
	defer deferrer.Close()
	if res == nil {
		msg := " failed to load data from db"
		return c.JSON(http.StatusInternalServerError, msg)

	}

	return c.JSON(http.StatusOK, res)

}

func GetUserById(c echo.Context) error {
	user := new(models.Id)
	err := c.Bind(user)
	if err != nil {
		msg := "bind failed"
		return c.JSON(http.StatusBadRequest, msg)

	}
	res, _ := models.GetUserById(user)
	if res == nil {
		msg := "no response"
		return c.JSON(http.StatusInternalServerError, msg)
	}
	return c.JSON(http.StatusOK, res)

}
