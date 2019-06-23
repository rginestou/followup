package main

import (
	"followup/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/mgo.v2/bson"
)

var jwtAuth = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
})

func extractUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := c.Get("user").(*jwt.Token)
		claims := u.Claims.(jwt.MapClaims)
		if err := claims.Valid(); err != nil {
			return echo.ErrUnauthorized
		}

		var id bson.ObjectId
		if rid, ok := claims["user_id"]; ok {
			id = bson.ObjectIdHex(rid.(string))
		}

		user, err := models.GetUserByID(id)
		if err != nil {
			return echo.ErrUnauthorized
		}

		c.Set("user_id", user.ID)
		return next(c)
	}
}
