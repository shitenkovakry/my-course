package main

import (
	"fmt"
	"strings"
)

func isCapital(word string) bool {
	firstLetter := word[0:1]
	capitalisedFirstLetter := strings.ToUpper(firstLetter)

	if capitalisedFirstLetter != firstLetter {
		return false
	}

	otherLetters := word[1:]
	loweredOtherLetters := strings.ToLower(otherLetters)

	if otherLetters != loweredOtherLetters {
		return false
	}

	return true
}

func main() {
	line := " Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software."
	words := strings.Split(line, " ")
	numberOfCapitalisedWords := 0

	for i := 0; i < len(words); i++ {
		word := words[i]
		isWordCapital := isCapital(word)

		if isWordCapital == true {
			numberOfCapitalisedWords += 1
		}
	}

	fmt.Println(numberOfCapitalisedWords)
}
