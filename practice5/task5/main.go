package main

import (
	"errors"
	"fmt"
	"math"
)

var (
	ErrNoResult = errors.New("нет действительных корней")
)

func Discriminant(a, b, c float64) (float64, float64, error) {
	discriminant := (b * b) - (4 * a * c)
	d := math.Sqrt(discriminant)

	if discriminant > 0 {
		x1 := (-b + d) / (2 * a)
		x2 := (-b - d) / (2 * a)

		return x1, x2, nil
	} else if discriminant == 0 {
		x := -b / (2 * a)

		return x, x, nil
	}

	return 0, 0, ErrNoResult
}

func main() {
	fmt.Println("Введите коэффициент а:")
	var a float64
	fmt.Scan(&a)

	fmt.Println("Введите коэффициент b:")
	var b float64
	fmt.Scan(&b)

	fmt.Println("Введите коэффициент с:")
	var c float64
	fmt.Scan(&c)

	x1, x2, err := Discriminant(a, b, c)
	if err != nil {
		fmt.Println("нет дейсвтительных корней")

		return
	}

	fmt.Println(x1, x2)
}
