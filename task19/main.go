package main

import "fmt"

func main() {
	fmt.Println("первое число:")
	var number1 int
	fmt.Scan(&number1)

	fmt.Println("второе число:")
	var number2 int
	fmt.Scan(&number2)

	fmt.Println("третье число:")
	var number3 int
	fmt.Scan(&number3)

	more := true
	less := false

	if number1 > number2 {
		fmt.Println(more, "- значит первое число максимальное из трех введенных")
		return
	} else if number1 > number3 {
		fmt.Println(more, "- значит первое число максимальное из трех введенных")
		return
	}

	fmt.Println(less, ", а значит первое число не является максимальным из трех введенных")

}
