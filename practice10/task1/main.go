package main

import (
	"fmt"
)

func main() {
	fmt.Println("введите степень x:")
	var x float64
	fmt.Scan(&x)

	fmt.Println("введите n:")
	var n int
	fmt.Scan(&n)

	ex := 1.0
	xi := 1.0

	for i := 1; i < n; i++ {
		xi *= x / float64(i)
		ex = ex + xi
	}

	fmt.Println(ex)
}

// 2.718281828459045
// 2.7182818284590455
