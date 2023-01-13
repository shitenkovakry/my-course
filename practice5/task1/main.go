package main

import "fmt"

func main() {
	fmt.Println("введите координату Х:")
	var coordinateX int
	fmt.Scan(&coordinateX)

	fmt.Println("введите координату Y:")
	var coordinateY int
	fmt.Scan(&coordinateY)

	if coordinateX > 0 && coordinateY > 0 {
		fmt.Println("Ваша координата находится в 1-ой четверти")
		return
	} else if coordinateX < 0 && coordinateY > 0 {
		fmt.Println("Ваша координата находится во 2-ой четверти")
		return
	} else if coordinateX < 0 && coordinateY < 0 {
		fmt.Println("Ваша координата находится в 3-ей четверти")
		return
	} else if coordinateX > 0 && coordinateY < 0 {
		fmt.Println("Ваша координата находится в 4-ой четверти")
		return
	} else if coordinateX > 0 && coordinateY == 0 {
		fmt.Println("Ваша координата не принадлежит к какой-либо четверти")
		return
	} else if coordinateX < 0 && coordinateY == 0 {
		fmt.Println("Ваша координата не принадлежит к какой-либо четверти")
		return
	} else if coordinateX == 0 && coordinateY > 0 {
		fmt.Println("Ваша координата не принадлежит к какой-либо четверти")
		return
	} else if coordinateX == 0 && coordinateY < 0 {
		fmt.Println("Ваша координата не принадлежит к какой-либо четверти")
		return
	}

	fmt.Println("Ваша координата находится в центре")
}
