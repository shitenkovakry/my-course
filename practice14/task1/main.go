package main

import "fmt"

func IsEven(number int) bool {
	if number%2 == 0 {
		return true
	}

	return false
}

func main() {
	number := 35

	result := IsEven(number)

	fmt.Println(result)
}
