package main

import "fmt"

func main() {
	type Game struct {
		ID   int
		Life int
		Name string
	}

	var ross Game
	ross.ID = 125
	ross.Life = 3
	ross.Name = "Resorter"

	fmt.Println(ross.ID)
}
