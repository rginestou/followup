package controllers

import (
	"followup/models"
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

// GetUsers returns all users
func GetUsers(c echo.Context) error {
	users, err := models.GetUsers()
	if err != nil {
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, users)
}

// GetUser returns one user by their ID
func GetUser(c echo.Context) error {
	if !bson.IsObjectIdHex(c.Param("id")) {
		return echo.ErrBadRequest
	}

	id := bson.ObjectIdHex(c.Param("id"))
	user, err := models.GetUserByID(id)
	if err != nil {
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, user)
}

// GetMe returns the current user
func GetMe(c echo.Context) error {
	id := c.Get("user_id").(bson.ObjectId)
	user, err := models.GetUserByID(id)
	if err != nil {
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, user)
}

// AddUser inserts a new user in the database
func AddUser(c echo.Context) error {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		return echo.ErrBadRequest
	}

	id, err := models.AddUser(user)
	if err != nil {
		return echo.ErrNotFound
	}

	iuser, _ := models.GetUserByID(id)

	return c.JSON(http.StatusOK, iuser)
}

func UpdateUser(c echo.Context) error {
	return nil
}

func DeleteUser(c echo.Context) error {
	return nil
}
