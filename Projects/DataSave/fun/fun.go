package fun

import (
	"errors"

	auth "github.com/Rioverde/Go/Projects/DataSave/conn"
	"github.com/Rioverde/Go/Projects/DataSave/models"
)

var errNoUser = errors.New("No user exists")
var err error

// CheckUser checks if a user exists in a db, if yes returns a struct of user, else error.
func CheckUser(name string, password string) (*models.User, error) {
	db := auth.DBconnection()
	defer db.Close()

	user := models.User{}

	if password == "" {
		//search a user where name = ?, and return a row from database
		r := db.Table("User").Where("name = ?", name).Row()
		//scan that database row into the user struct
		r.Scan(&user.ID, &user.Name, &user.Password)
	} else {
		r := db.Table("User").Where("name = ? AND password = ?", name, password).Row()
		r.Scan(&user.ID, &user.Name, &user.Password)
	}

	if user.Name != "" {
		return &user, nil
	}

	return nil, errNoUser
}

//AddDataToDB is function to add data to db
func AddDataToDB(name string, email string, password string) {
	db := auth.DBconnection()
	defer db.Close()

	data := models.Data{
		Name:     name,
		Email:    email,
		Password: password,
	}

	db.Table("Users").Save(&data)
}
