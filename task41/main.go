package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func CompareTwoKeys(left, right string) bool {
	splitedLeft := strings.Split(left, ".")
	splitedRight := strings.Split(right, ".")

	if len(splitedLeft) > len(splitedRight) {
		differenceBetweenLens := len(splitedLeft) - len(splitedRight)

		for i := 0; i < differenceBetweenLens; i++ {
			zero := "0"
			splitedRight = append(splitedRight, zero)
		}
	}

	if len(splitedRight) > len(splitedLeft) {
		differenceBetweenLens := len(splitedRight) - len(splitedLeft)

		for i := 0; i < differenceBetweenLens; i++ {
			zero := "0"
			splitedLeft = append(splitedLeft, zero)
		}
	}

	for i := 0; i < len(splitedLeft); i++ {
		leftStr := splitedLeft[i]
		left, _ := strconv.Atoi(leftStr)

		rightStr := splitedRight[i]
		right, _ := strconv.Atoi(rightStr)

		if left == right {
			continue
		}

		if left > right {
			return false
		} else {
			return true
		}
	}

	return true
}

func Organize(input []string) []string {
	sorted := append([]string{}, input...)

	sort.Slice(sorted, func(i, j int) bool {
		left := sorted[i]
		right := sorted[j]

		fmt.Println(left, "<", right, left < right)

		return CompareTwoKeys(left, right)
	})

	return sorted
}

func main() {

}
