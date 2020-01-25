package main

import "fmt"

func main() {
	var sum int
	for i := 0; i < 10; i++ {
		sum += i
		if sum > 10 {
			fmt.Println(sum)
		}
	}

	if _, err := fmt.Println(123); err != nil { // check if there is error
		fmt.Println("bad")
	} else { // if no then print good
		fmt.Println("good")
	}

	fmt.Println(sum)
}
