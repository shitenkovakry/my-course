package main

import (
	"fmt"
	"strings"
)

func CreateTree(n int) /*[][]string */ {
	arrayOfString := make([][]string, n)
	space := " "
	symbol := "#"

	for i := 1; i <= n; i++ {

		currentLine := make([]string, n)

		for j := 0; j < len(currentLine); j++ {
			currentLine[j] = symbol
		}

		for j := 0; j < len(currentLine)-i; j++ {
			currentLine[j] = space
		}

		arrayOfString[i-1] = currentLine
	}

	for _, line := range arrayOfString {
		fmt.Println(strings.Join(line, ""))
	}
}

func main() {
	CreateTree(6)
}
