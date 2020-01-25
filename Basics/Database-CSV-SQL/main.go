package main

import (
	"fmt"
	"log"

	"github.com/Rioverde/Go/Database-CSV-SQL/db"
)

func main() {
	err := db.Insert("/home/resorter/Go/src/github.com/Rioverde/Go/Database-CSV-SQL/data.csv")
	if err != nil {
		log.Fatal(err)
	}

	var s db.Student

	name, err := s.Get("name", "adak")
	gpa, err := s.Get("gpa", "adak")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Name is: %v\n", name)
	fmt.Printf("CGPA is: %v\n", gpa)

}

// F12 to see definition of function
