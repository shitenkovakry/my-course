package main

import "fmt"

func main() {
	var i int
	a := i * i
	fmt.Scan(&i)
	fmt.Println("квадрат числа", i, "будет равен", a)
}
