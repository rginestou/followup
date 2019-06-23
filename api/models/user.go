package models

import (
	"followup/utils"

	"gopkg.in/mgo.v2/bson"
)

// User model
type User struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Username string        `json:"username" bson:"username"`
	Email    string        `json:"email" bson:"email"`
	Password string        `json:"-" bson:"password"`
}

// GetUsers returns all users
func GetUsers() (users []User, err error) {
	err = userCol.Find(bson.M{}).All(&users)
	return
}

// GetUserByID returns one user by their ID
func GetUserByID(id bson.ObjectId) (user *User, err error) {
	err = userCol.Find(bson.M{}).One(&user)
	return
}

// GetUserByUsername returns one user by their username
func GetUserByUsername(username string) (user *User, err error) {
	err = userCol.Find(bson.M{"username": username}).One(&user)
	return
}

// AddUser inserts a new user in the database
func AddUser(user User) (id bson.ObjectId, err error) {
	id = bson.NewObjectId()
	user.ID = id
	user.Password, err = utils.GenerateHash(user.Password)
	if err != nil {
		return id, nil
	}

	err = userCol.Insert(user)
	return
}

// UpdateUser updates an existing user
func UpdateUser(id bson.ObjectId, user User) (err error) {
	user.ID = ""
	err = userCol.UpdateId(id, user)
	return
}

// DeleteUser removes a user from the database
func DeleteUser(id bson.ObjectId) (err error) {
	err = userCol.RemoveId(id)
	return
}
