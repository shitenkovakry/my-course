package main

import "fmt"

func main() {
	var productCost int
	var deliveryCost int
	var discount int

	fmt.Scan(&productCost)
	fmt.Scan(&deliveryCost)
	fmt.Scan(&discount)

	price := productCost + deliveryCost - discount

	fmt.Println("Стоимость товара:", productCost)
	fmt.Println("Стоимость доставки:", deliveryCost)
	fmt.Println("Размер скидки:", discount)
	fmt.Println("---------")
	fmt.Println("Итого:", price)
}
