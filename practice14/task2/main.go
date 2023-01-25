package main

import (
	"fmt"
	"math/rand"
)

func RandomNumbers() (int, int) {
	number1 := rand.Intn(100)
	number2 := rand.Intn(100)

	return number1, number2
}

func ChangeBasis(x1, y1 int) (int, int) {
	x := 2*x1 + 10
	y := -3*y1 - 5.

	return x, y
}

func ChangeBasis2(x2, y2 int) (int, int) {
	x := 2*x2 + 10
	y := -3*y2 - 5.

	return x, y
}

func ChangeBasis3(x3, y3 int) (int, int) {
	x := 2*x3 + 10
	y := -3*y3 - 5.

	return x, y
}

func main() {
	x1, y1 := RandomNumbers()
	x2, y2 := RandomNumbers()
	x3, y3 := RandomNumbers()

	xc1, yc1 := ChangeBasis(x1, y1)
	fmt.Println(xc1, yc1)

	xc2, yc2 := ChangeBasis2(x2, y2)
	fmt.Println(xc2, yc2)

	xc3, yc3 := ChangeBasis3(x3, y3)
	fmt.Println(xc3, yc3)
}
