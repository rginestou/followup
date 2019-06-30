package main

import (
	"followup/models"
	"followup/utils"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func login(c echo.Context) error {
	var body credentials
	err := c.Bind(&body)
	if err != nil {
		return echo.ErrBadRequest
	}

	// Fetch user
	user, err := models.GetUserByUsername(body.Username)
	if err != nil {
		return echo.ErrUnauthorized
	}

	// Check password
	err = utils.CompareWithHash(user.Password, body.Password)
	if err != nil {
		return echo.ErrUnauthorized
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"user":  user,
		"token": t,
	})
}
