package main

import (
	"fmt"
	"strings"
)

const (
	notFound = -1
)

func TestStringAgainstRune(line string, input []rune) []int {
	arrayOfResults := make([]int, len(input))

	for index := range input {
		arrayOfResults[index] = notFound
	}

	for positionOfLine := 0; positionOfLine < len(line); positionOfLine++ {
		for positionOfRune := 0; positionOfRune < len(input); positionOfRune++ {
			runeToCompare := strings.ToUpper(string(input[positionOfRune]))
			lineToCompare := strings.ToUpper(string(line[positionOfLine]))

			if runeToCompare == lineToCompare {
				arrayOfResults[positionOfRune] = positionOfLine
			}
		}
	}

	return arrayOfResults
}

func ParseTest(arrayOfLines []string, runes []rune) [][]int {
	result := [][]int{}
	for _, line := range arrayOfLines {
		res := TestStringAgainstRune(line, runes)

		result = append(result, res)
	}

	return result
}

func main() {
	arrayOfStrings := []string{"Hello world", "miu"}
	runes := []rune{'h', 'l', 'k'}

	testStrings := ParseTest(arrayOfStrings, runes)

	for _, result := range testStrings {
		for positionRune, position := range result {
			if position == notFound {
				continue
			}

			fmt.Println(string(runes[positionRune]), "position", position)
		}
	}

}
