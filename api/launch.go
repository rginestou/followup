package main

import (
	"fmt"
	"followup/models"

	c "github.com/labstack/gommon/color"
)

func testLaunchSetup() error {
	// Create first user
	if models.GetUsersCount() == 0 {
		var username, password string

		c.Println(c.Green("Creating first user:"))
		fmt.Print("Username: ")
		fmt.Scan(&username)
		fmt.Print("Password: ")
		fmt.Scan(&password)

		_, err := models.AddUser(models.User{
			Username: username,
			Password: password,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
