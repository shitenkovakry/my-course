package main

import "log"

func CalculateByFunc(x, y int, fn func(int, int) int) {
	defer func() {
		z := fn(x, y)
		log.Println("defered", x, y, z)
	}()

	z := fn(x, y)
	log.Println("normal", x, y, z)
	x++
	y++
}

func main() {
	add := func(a, b int) int {
		return a + b
	}

	sub := func(a, b int) int {
		return a - b
	}

	mult := func(a, b int) int {
		return a * b
	}

	CalculateByFunc(1, 2, add)
	CalculateByFunc(1, 2, sub)
	CalculateByFunc(1, 2, mult)
}
