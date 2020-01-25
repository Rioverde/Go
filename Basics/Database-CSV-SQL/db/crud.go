//Package db file crud, is used for create, read, update and delete functions
package db

import (
	"encoding/csv" //for working with csv files
	"errors"
	"io"
	"os"      //for working with file
	"strconv" //for converting string to float

	//sqllite
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var errWrongParameter = errors.New("Wrong Parameter")

// Insert puts a data into a database table, from csv file. Path parameter is a path to .csv file. Return value is error or nil.
// This function opens a file, then it reads from this file data and copies it to db
func Insert(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	//NewReader functions reads from a file and you have to include an open file in it's parameters
	r := csv.NewReader(f)

	db := ConnectToDB()
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		var temp = Student{
			Name:       record[0],
			Surname:    record[1],
			Department: record[2],
		}
		//because ParseFloat function returns 2 values and you can't assign 2 values to temp.Gpa
		temp.Gpa, err = strconv.ParseFloat(record[3], 64)
		if err != nil {
			return err
		}

		//INSERT INTO student (name, surname, department, gpa) VALUES adakw
		db.Table("student").Create(&temp)
	}

	return nil
}

// Get is a receiver function, uppercase GetName, because to use it in other packages.
// If we don't know return type, we return interface{}.
func (s *Student) Get(par, val string) (interface{}, error) {
	//this is connection to database
	db := ConnectToDB()

	switch par {
	case "name":
		db.Table("student").Where("name = ?", val).First(&s)
		return s.Name, nil
	case "surname":
		db.Table("student").Where("surname = ?", val).First(&s)
		return s.Surname, nil
	case "department":
		db.Table("student").Where("department = ?", val).First(&s)
		return s.Department, nil
	case "gpa":
		db.Table("student").Where("gpa = ?", val).First(&s)
		return s.Gpa, nil
	}

	return "", errWrongParameter
}

/*	same as previous function
func GetName(s *Student) {
	s.Name
}
*/

/*
class student {
	string name;

	string getname() {
		return (*this).name;
	} //returns name
}

student s;

s.getname()
*/
