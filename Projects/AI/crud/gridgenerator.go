package gridgenerator

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//GridLenght is size of gris
const GridLenght = 8

// Blackcell needed to create cells that unavailbale
var Blackcell string

// WhiteCell needed to create available cells
var WhiteCell string

//Grid is creating map of values
var Grid = make(map[int]string)

//takingWhite string to track white Cells
var takingWhite string

//White Taking whole white cells
var White string

//GenerateGrid us a grid with Black and White Cells
func GenerateGrid() {

	//Creating a grid.txt file
	f, err := os.Create("Full-Grid.txt")
	check(err)
	defer f.Close()

	_, err = f.WriteString("White Cells: \n")
	check(err)

	var z int = 1

	for i := 0; i < GridLenght; i++ {
		for y := 0; y < GridLenght; y++ {
			Row := strconv.Itoa(i)
			Col := strconv.Itoa(y)
			key := GenerateRandomNumber(1, 100)
			if key >= 15 {
				Whitecell := ("(" + Row + "," + Col + ")")
				Grid[z] = "1"
				z++
				White += Whitecell
				_, err = f.WriteString(Whitecell)
				check(err)
				//Giving chance for Black Cells to born
			} else if key <= 15 {
				Whitecell := ("(" + Row + "," + Col + ")")
				Grid[z] = "-1"
				z++
				Blackcell += Whitecell
			}
		}
	}

	//Here we generate starting node and ending node
	var UpdateStart = GenerateRandomNumber(1, 64)
	Grid[UpdateStart] = "Start"
	var UpdateEnd = GenerateRandomNumber(1, 64)
	Grid[UpdateEnd] = "End"

	var Massive string = "\n\nBlack Cells: \n" + Blackcell +
		"\n\nStart Point " + strconv.Itoa(UpdateStart) + " Node" +
		"\nEnd Point " + strconv.Itoa(UpdateEnd) + " Node"

	_, err = f.WriteString(Massive)
	check(err)

	fmt.Println(Grid)

}

//GenerateRandomNumber generates random number
func GenerateRandomNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

//check function needed to check errors
func check(err error) {
	if err != nil {
		fmt.Printf("\nError : %s", err.Error())
		os.Exit(1)
	}
}
