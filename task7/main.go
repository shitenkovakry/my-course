package main

import "fmt"

func main() {
	fmt.Println("Первый товар:")
	var product1 string
	fmt.Scan(&product1)

	fmt.Println("Стоимость товара:")
	var costOfProduct1 int
	fmt.Scan(&costOfProduct1)

	fmt.Println("Второй товар:")
	var product2 string
	fmt.Scan(&product2)

	fmt.Println("Стоимость товара:")
	var costOfProduct2 int
	fmt.Scan(&costOfProduct2)

	fmt.Println("Третий товар:")
	var product3 string
	fmt.Scan(&product3)

	fmt.Println("Стоимость товара:")
	var costOfProduct3 int
	fmt.Scan(&costOfProduct3)

	totalProductsCost := float64(costOfProduct1 + costOfProduct2 + costOfProduct3)

	if totalProductsCost > 10000 {
		discount := 0.1 * totalProductsCost
		fmt.Println("Вы купили 3 товара, привышающие 10000 руб, поэтому у вас скидка 10%", discount)
		totalProductsCost = totalProductsCost - discount
	}

	fmt.Println("Итого:", totalProductsCost)

}
