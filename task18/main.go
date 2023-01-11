package main

import "fmt"

func main() {
	fmt.Println("Введите] сколько 100 рублей:")
	var rub100 float64
	fmt.Scan(&rub100)

	fmt.Println("Введите сколько 50 рублей:")
	var rub50 float64
	fmt.Scan(&rub50)

	fmt.Println("Введите сколько 20 рублей:")
	var rub20 float64
	fmt.Scan(&rub20)

	fmt.Println("Введите сколько 10 рублей:")
	var rub10 float64
	fmt.Scan(&rub10)

	fmt.Println("Введите сколько 5 рублей:")
	var rub5 float64
	fmt.Scan(&rub5)

	fmt.Println("Введите сколько 2-х рублей:")
	var rub2 float64
	fmt.Scan(&rub2)

	fmt.Println("Введите сколько 1-го рубля:")
	var rub1 float64
	fmt.Scan(&rub1)

	fmt.Println("Введите сколько 50 копеек:")
	var coint50 float64
	fmt.Scan(&coint50)

	fmt.Println("Введите сколько 20 копеек:")
	var coint20 float64
	fmt.Scan(&coint20)

	fmt.Println("Введите сколько 10 копеек")
	var coint10 float64
	fmt.Scan(&coint10)

	fmt.Println("Введите сколько 5 копеек:")
	var coint5 float64
	fmt.Scan(&coint5)

	fmt.Println("Введите сколько 2 копеек:")
	var coint2 float64
	fmt.Scan(&coint2)

	fmt.Println("Введите сколько 1 копейки")
	var coint1 float64
	fmt.Scan(&coint1)

	totalRun := (rub100 * 100) + (rub50 * 50) + (rub20 * 20) + (rub10 * 10) + (rub5 * 5) + (rub2 * 2) + (rub1 * 1) + (coint50 * 0.5) + (coint20 * 0.2) + (coint10 * 0.1) + (coint5 * 0.05) + (coint2 * 0.02) + (coint1 * 0.01)

	fmt.Println("Всего в кассе:", totalRun, "белорусских рублей")
}
