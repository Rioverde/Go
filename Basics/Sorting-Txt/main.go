package main

import (
	"bufio"
	"fmt"
	"os"
)

//cat data.txt | go run main.go

func main() {
	in := bufio.NewScanner(os.Stdin)
	seen := make(map[string]bool)
	for in.Scan() {

		txt := in.Text()

		if _, found := seen[txt]; found {
			continue
		}

		seen[txt] = true
		fmt.Println(txt)
	}
}
