//Package db this file is for connection to db
package db

import (
	"log"

	//this is needed for working with sqlite database, underscore _ is needed when another library is using this library
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/jinzhu/gorm"
)

//ConnectToDB returns a connection to gorm database.It returns a pointer(or address) to DB, because in other case it creates it's clone. This is for performance issues.
func ConnectToDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "/home/resorter/Go/src/github.com/Rioverde/Go/Database-CSV-SQL/student.db")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
