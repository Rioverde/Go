package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	/*
		var a, b int	//declaration
		var a int = 2	//initialization
		var c, d int = 4, 5		// multiple initialization

		var c, d = 4, 3		// without type

		var r, f, j = "ale", 3.0, true
	*/
}
