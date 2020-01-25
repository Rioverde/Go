package database

import (
	"errors"

	models "github.com/Rioverde/Go/Projects/Web-Application/Models"
	"github.com/Rioverde/Go/Projects/Web-Application/auth"
)

var errNoUser = errors.New("No user exists")
var err error

//AddUser adds user in database
func AddUser(email, password string) {
	db := auth.ConnectToDB()
	defer db.Close()

	user := models.Users{
		Email:    email,
		Password: password,
	}

	db.Table("Users").Save(&user)

}

// CheckUser checks if a user exists in a db, if yes returns a struct of user, else error.
func CheckUser(email string, password string) (*models.Users, error) {
	db := auth.ConnectToDB()
	defer db.Close()

	user := models.Users{}

	if password == "" {
		//search a user where email = ?, and return a row from database
		r := db.Table("Users").Where("email = ?", email).Row()
		//scan that database row into the user struct
		r.Scan(&user.ID, &user.Email, &user.Password)
	} else {
		r := db.Table("Users").Where("email = ? AND password = ?", email, password).Row()
		r.Scan(&user.ID, &user.Email, &user.Password)
	}

	if user.Email != "" {
		return &user, nil
	}

	return nil, errNoUser
}
