package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on : ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		defer fmt.Println("------------")
		fmt.Println("Linux")
	default:
		fmt.Print("%s.\n", os)
	}

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

}
