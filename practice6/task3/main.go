package main

import "fmt"

func ComputeDiscount(price, discount float64) float64 {
	if discount > 30.0 {
		discount = 30.0
	}

	result := price * discount / 100.0

	if result > 2000 {
		return 2000
	}

	return result
}

func main() {
	fmt.Println("Введите стоимость товара:")
	var price float64
	fmt.Scan(&price)

	fmt.Println("Пример, 20 означает скидку на 20%. Введите скидку:")
	var discount float64
	fmt.Scan(&discount)

	discountResult := ComputeDiscount(price, discount)
	fmt.Println("Скидка равна", discountResult)

}
