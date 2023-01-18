package main

import "fmt"

func main() {
	fmt.Println("Введите ширину:")
	var width int
	fmt.Scan(&width)

	fmt.Println("Введите высоту")
	var height int
	fmt.Scan(&height)

	maxRows := height
	maxColumns := width

	for row := 0; row < maxRows; row++ {
		toPrintRow := ""
		color := true

		if row%2 == 0 {
			color = false
		}

		for column := 0; column < maxColumns; column++ {
			if color == true {
				toPrintRow = toPrintRow + "*"
			} else if color == false {
				toPrintRow = toPrintRow + "-"
			}
			color = !color
		}

		fmt.Println(toPrintRow)
	}
}
