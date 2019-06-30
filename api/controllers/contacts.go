package controllers

import (
	"followup/models"
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

// GetContacts returns all contacts
func GetContacts(c echo.Context) error {
	contacts, err := models.GetContacts()
	if err != nil {
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, contacts)
}

// GetContactsLike returns all contacts matching a pattern
func GetContactsLike(c echo.Context) error {
	contacts, _ := models.GetContactsLike(c.Param("pattern"))

	return c.JSON(http.StatusOK, contacts)
}

// GetContact returns one contact by their ID
func GetContact(c echo.Context) error {
	if !bson.IsObjectIdHex(c.Param("id")) {
		return echo.ErrBadRequest
	}
	id := bson.ObjectIdHex(c.Param("id"))

	contact, err := models.GetContactByID(id)
	if err != nil {
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, contact)
}

// AddContact inserts a new contact in the database
func AddContact(c echo.Context) error {
	var body models.Contact
	err := c.Bind(&body)
	if err != nil {
		return echo.ErrBadRequest
	}

	id, err := models.AddContact(body)
	if err != nil {
		return echo.ErrNotFound
	}

	contact, _ := models.GetContactByID(id)

	return c.JSON(http.StatusOK, contact)
}

// UpdateContact updates a contact by their ID
func UpdateContact(c echo.Context) error {
	if !bson.IsObjectIdHex(c.Param("id")) {
		return echo.ErrBadRequest
	}
	id := bson.ObjectIdHex(c.Param("id"))

	var body models.Contact
	err := c.Bind(&body)
	if err != nil {
		return echo.ErrBadRequest
	}

	err = models.UpdateContact(id, body)
	if err != nil {
		return echo.ErrBadRequest
	}

	contact, _ := models.GetContactByID(id)

	return c.JSON(http.StatusOK, contact)
}

// DeleteContact deletes a contact by their ID
func DeleteContact(c echo.Context) error {
	if !bson.IsObjectIdHex(c.Param("id")) {
		return echo.ErrBadRequest
	}
	id := bson.ObjectIdHex(c.Param("id"))

	err := models.DeleteContact(id)
	if err != nil {
		return echo.ErrNotFound
	}

	return c.NoContent(http.StatusOK)
}
