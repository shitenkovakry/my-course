package main

import "fmt"

func CalculateCash(rub100, rub50, rub20, rub10, rub5, rub2, rub1, coint50, coint20, coint10, coint5, coint2, coint1 uint) float32 {
	totalCash := float32(0)
	totalCash += float32(rub100) * 100.0
	totalCash += float32(rub50) * 50.0
	totalCash += float32(rub20) * 20.0
	totalCash += float32(rub10) * 10.0
	totalCash += float32(rub5) * 5.0
	totalCash += float32(rub2) * 2.0
	totalCash += float32(rub1) * 1.0
	totalCash += float32(coint50) * 0.5
	totalCash += float32(coint20) * 0.2
	totalCash += float32(coint10) * 0.1
	totalCash += float32(coint5) * 0.05
	totalCash += float32(coint2) * 0.02
	totalCash += float32(coint1) * 0.0

	return totalCash
}

func main() {
	fmt.Println("Введите сколько 100 рублей:")
	var rub100 uint
	fmt.Scan(&rub100)

	fmt.Println("Введите сколько 50 рублей:")
	var rub50 uint
	fmt.Scan(&rub50)

	fmt.Println("Введите сколько 20 рублей:")
	var rub20 uint
	fmt.Scan(&rub20)

	fmt.Println("Введите сколько 10 рублей:")
	var rub10 uint
	fmt.Scan(&rub10)

	fmt.Println("Введите сколько 5 рублей:")
	var rub5 uint
	fmt.Scan(&rub5)

	fmt.Println("Введите сколько 2-х рублей:")
	var rub2 uint
	fmt.Scan(&rub2)

	fmt.Println("Введите сколько 1-го рубля:")
	var rub1 uint
	fmt.Scan(&rub1)

	fmt.Println("Введите сколько 50 копеек:")
	var coint50 uint
	fmt.Scan(&coint50)

	fmt.Println("Введите сколько 20 копеек:")
	var coint20 uint
	fmt.Scan(&coint20)

	fmt.Println("Введите сколько 10 копеек")
	var coint10 uint
	fmt.Scan(&coint10)

	fmt.Println("Введите сколько 5 копеек:")
	var coint5 uint
	fmt.Scan(&coint5)

	fmt.Println("Введите сколько 2 копеек:")
	var coint2 uint
	fmt.Scan(&coint2)

	fmt.Println("Введите сколько 1 копейки")
	var coint1 uint
	fmt.Scan(&coint1)

	result := CalculateCash(rub100, rub50, rub20, rub10, rub5, rub2, rub1, coint50, coint20, coint10, coint5, coint2, coint1)

	fmt.Println("Всего в кассе:", result, "белорусских рублей")
}
