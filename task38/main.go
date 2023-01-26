package main

import "fmt"

func main() {
	fmt.Println("введите число начальное:")
	var i int
	fmt.Scan(&i)

	fmt.Println("введите число конечное:")
	var y int
	fmt.Scan(&y)

	for a := i; a <= y; a++ {
		if a%2 == 0 {
			fmt.Println(a)
		}
	}
}
