package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type animal struct {
	id          int
	anymal_type string
	nickname    string
	location    int
	age         int
}

func main() {
	//connect to Database
	db, err := sql.Open("mysql", "resorter:resorter1403@/Dino")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	//Query

	rows, err := db.Query("select from * Dino.animals where id=?", 1)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
}
