package main

import "fmt"

func main() {
	a := []int{10, 2, 13, 15, 6, 7, 3}

	for i, v := range a {
		fmt.Println(i, v)
	}

}
