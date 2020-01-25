package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	var FileName string = ""
	fmt.Scanln(&FileName)
	if _, err := os.Stat(FileName); os.IsNotExist(err) {
		fmt.Printf("File does not exist\n")

	} else if err == nil {
		fmt.Println(FileName + " File Exists")
	}

	FileName, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Location your file in is: ")
	fmt.Println(FileName)
}
