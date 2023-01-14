package main

import "fmt"

func main() {
	fmt.Println("Введите число:")
	var x int
	fmt.Scan(&x)

	if x >= 0 {
		fmt.Print(x)
		return
	}

	x = -x
	fmt.Print(x)

}
