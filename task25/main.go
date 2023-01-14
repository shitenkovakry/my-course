package main

import "fmt"

func main() {
	fmt.Println("Введите первое число:")
	var number int
	fmt.Scan(&number)

	maxRun := 4

	run := 1
	summa := 0

	for {
		run++
		summa += number

		if run > maxRun {
			break
		}

		fmt.Println("Введите", run, "й число:")
		fmt.Scan(&number)

	}

	fmt.Println("summa:", summa)
}
