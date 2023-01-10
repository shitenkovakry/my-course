package main

import "fmt"

func main() {
	fmt.Println("Введите число:")
	var number int
	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Println("число", number, "четное")
		return
	}

	fmt.Println("число", number, "нечетное")
}
