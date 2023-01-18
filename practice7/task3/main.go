package main

import "fmt"

func main() {
	fmt.Println("Введите высоту елочки:")
	var height int
	fmt.Scan(&height)

	for row := 0; row < height; row++ {
		toPrintRow := ""
		spaces := height - row - 1

		totalSpaces := ""

		for space := 0; space < spaces; space++ {
			totalSpaces += "-"
		}

		totalStars := "*"
		stars := row * 2

		for star := 0; star < stars; star++ {
			totalStars += "*"
		}

		toPrintRow = totalSpaces + totalStars + totalSpaces
		fmt.Println(toPrintRow)
	}
}
