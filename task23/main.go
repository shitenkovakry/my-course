package main

import "fmt"

func main() {
	fmt.Println("введите возраст первого ребенка:")
	var ageFirstChild int
	fmt.Scan(&ageFirstChild)

	fmt.Println("введите возраст второго ребенка:")
	var ageSecondChild int
	fmt.Scan(&ageSecondChild)

	needSubsidy := false

	if ageFirstChild > 3 && ageFirstChild < 16 {
		needSubsidy = true
	}

	if ageSecondChild > 3 && ageSecondChild < 16 {
		needSubsidy = true
	}

	if needSubsidy {
		fmt.Println("субсидия положена")
	} else {
		fmt.Println("субсидия не положена")
	}
}
