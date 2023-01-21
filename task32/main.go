package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	a = 0
	epsion := 0.0001

	for i := 0; i < 15; i++ {
		a += 0.1
		fmt.Println(a)

		if math.Abs(a-1) < epsion {
			fmt.Println("а приблизительно равно 1")
		}

		fmt.Println("не равно 1")
	}
}
