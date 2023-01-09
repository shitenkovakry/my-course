package main

import "fmt"

func main() {
	var height int
	var growthRate int
	var decreaseRate int
	var days int

	fmt.Println("Высота саженца?")
	fmt.Scan(&height)

	fmt.Println("Скорость роста?")
	fmt.Scan(&growthRate)

	fmt.Println("Скорость поедания?")
	fmt.Scan(&decreaseRate)

	fmt.Println("Количество дней?")
	fmt.Scan(&days)

	finalGrowthRate := growthRate*2 + growthRate/2
	finalDecreaseRate := decreaseRate * 2
	finalHeight := height + finalGrowthRate - finalDecreaseRate

	fmt.Println("К середине", days, "дня бамбук вырастет до", finalHeight, "сантиметров")

	fmt.Printf("К середине %v дня бамбук вырастит до %v сантиметоров\n", days, finalHeight)
}
