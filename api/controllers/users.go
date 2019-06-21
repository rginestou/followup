package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func Restricted(c echo.Context) error {
	user := c.Get("user").(string)
	return c.String(http.StatusOK, user)
}
