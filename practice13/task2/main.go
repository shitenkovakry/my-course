package main

import "fmt"

func Swap(a *int, b *int) {
	*a, *b = *b, *a
}

func main() {
	x := 5
	y := 8

	Swap(&x, &y)

	fmt.Println(x, y)
}
