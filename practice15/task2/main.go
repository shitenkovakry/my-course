package main

import "fmt"

func ReverseArray(a []int) []int {
	newArray := []int{}

	for i := len(a) - 1; i >= 0; i-- {
		newArray = append(newArray, a[i])
	}

	return newArray
}

func main() {
	a := []int{6, 7, 9, 12}
	b := ReverseArray(a)

	fmt.Println(b)
}
