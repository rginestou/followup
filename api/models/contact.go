package models

import (
	"gopkg.in/mgo.v2/bson"
)

// Contact model
type Contact struct {
	ID bson.ObjectId `json:"id" bson:"_id,omitempty"`

	Name    string `json:"name" bson:"name"`
	Email   string `json:"email" bson:"email"`
	Phone   string `json:"phone" bson:"phone"`
	Address string `json:"address" bson:"address"`

	Groups []Group `json:"groups" bson:"groups"`
}

// GetContacts returns all contacts
func GetContacts() (contacts []Contact, err error) {
	err = contactCol.Find(bson.M{}).All(&contacts)
	return
}

// GetContactsLike returns all contact with info mathcing a pattern
func GetContactsLike(pattern string) (contacts []Contact, err error) {
	err = contactCol.Find(bson.M{"name": "/" + pattern + "/i"}).One(&contacts)
	return
}

// GetContactByID returns one contact by their ID
func GetContactByID(id bson.ObjectId) (contact *Contact, err error) {
	err = contactCol.FindId(id).One(&contact)
	return
}

// AddContact inserts a new contact in the database
func AddContact(contact Contact) (id bson.ObjectId, err error) {
	id = bson.NewObjectId()
	contact.ID = id

	err = contactCol.Insert(contact)
	return
}

// UpdateContact updates an existing contact
func UpdateContact(id bson.ObjectId, contact Contact) (err error) {
	contact.ID = ""
	err = contactCol.UpdateId(id, contact)
	return
}

// DeleteContact removes a contact from the database
func DeleteContact(id bson.ObjectId) (err error) {
	err = contactCol.RemoveId(id)
	return
}
