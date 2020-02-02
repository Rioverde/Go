package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Rioverde/Go/Projects/DataSave/fun"
)

func main() {
	//reading an integer
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	//reading a string
	reader2 := bufio.NewReader(os.Stdin)
	fmt.Print("Enter password: ")
	password, _ := reader2.ReadString('\n')

	fmt.Println(fun.CheckUser(username, password))
}
