package main

import (
	c "followup/controllers"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func main() {
	e := echo.New()
	e.HideBanner = true

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", accessible)
	e.POST("/login", login)

	// Authorized group
	r := e.Group("/users")
	r.Use(jwtAuth, extractUser)
	r.GET("", c.Restricted)

	e.Logger.Fatal(e.Start(":3000"))
}
