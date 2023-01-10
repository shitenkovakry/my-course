package main

import "fmt"

func main() {
	distanceBetweenInKm := 200
	timeInHours := 2

	fmt.Println("Введите среднюю скорость автомобиля:")
	var averageSpeed int
	fmt.Scan(&averageSpeed)

	distance := averageSpeed * timeInHours

	fmt.Println("Пройденное расстояние:", distance)

	if distance == distanceBetweenInKm {
		fmt.Println("Вы приехали!")
		return
	}

	if distance >= distanceBetweenInKm {
		fmt.Println("Вы переехали:(")
		return
	}

	fmt.Println("Вы не приехали:(")
}
