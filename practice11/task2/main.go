package main

import (
	"fmt"
	"strconv"
	"strings"
)

func SplitIntoWords(line string) []string {
	return strings.Split(line, " ")
}

func StringToInteger(separated []string) []int {
	toInts := []int{}
	for i := 0; i < len(separated); i++ {
		word := separated[i]
		toInt, err := strconv.Atoi(word)
		if err != nil {
			continue
		}

		toInts = append(toInts, toInt)
	}

	return toInts
}

func main() {
	fmt.Println("исходная строка:")
	var line string
	fmt.Scan(&line)

	separated := SplitIntoWords(line)
	fmt.Println(separated)

	wordToInt := StringToInteger(separated)
	fmt.Println("В строке содержатся числа в десятичном формате:", wordToInt)
}
