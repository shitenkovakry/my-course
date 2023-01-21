package main

import "fmt"

func main() {
	line := "Golang is programming language"

	findLater := "a"
	count := 0

	for index := 0; index < len(line); index++ {
		char := string(line[index])

		if char == findLater {
			count++
		}
	}

	fmt.Println(count)
}
