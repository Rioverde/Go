//Package auth this file is for connection to db
package auth

import (
	"log"
	"os"

	//this is needed for working with sqlite database, underscore _ is needed when another library is using this library
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/jinzhu/gorm"
)

//export WEBAPP_DB=path/to/your/db.file
var pathToDB = os.Getenv("SAVEDATA")

//ConnectToDB returns a connection to gorm database.It returns a pointer(or address) to DB, because in other case it creates it's clone. This is for performance issues.
func DBconnection() *gorm.DB {
	db, err := gorm.Open("sqlite3", pathToDB)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
