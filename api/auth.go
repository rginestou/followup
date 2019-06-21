package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var jwtAuth = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
})

func extractUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("user", "ME!")
		// user := c.Get("user").(*jwt.Token)
		// claims := user.Claims.(jwt.MapClaims)
		// isAdmin := claims["admin"].(bool)
		// if isAdmin == false {
		// 	return echo.ErrUnauthorized
		// }
		return next(c)
	}
}
