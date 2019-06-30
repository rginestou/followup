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
	// Test launch setup
	if err := testLaunchSetup(); err != nil {
		panic(err)
	}

	e := echo.New()
	e.HideBanner = true

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/", accessible)
	e.POST("/login", login)

	// Authorized groups
	// Users
	ru := e.Group("/users")
	ru.Use(jwtAuth, extractUser)
	ru.GET("/me", c.GetMe)
	ru.GET("/:id", c.GetUser)
	ru.GET("", c.GetUsers)
	ru.POST("", c.AddUser)
	ru.PUT("/:id", c.UpdateUser)
	ru.DELETE("/:id", c.DeleteUser)

	// Contacts
	rc := e.Group("/contacts")
	rc.Use(jwtAuth, extractUser)
	rc.GET("/like/:pattern", c.GetContactsLike)
	rc.GET("/:id", c.GetContact)
	rc.GET("", c.GetContacts)
	rc.POST("", c.AddContact)
	rc.PUT("/:id", c.UpdateContact)
	rc.DELETE("/:id", c.DeleteContact)

	e.Logger.Fatal(e.Start(":3000"))
}
