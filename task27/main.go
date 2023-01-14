package main

import "fmt"

func main() {
	fmt.Println("Вы за курицу, поэтому Вам нужно писать это слово. Начинаем. Введите слово:")
	var word string
	fmt.Scan(&word)

	stopWord := "надоело"

	for {
		if word == stopWord {
			break
		}

		fmt.Println("яйцо")

		fmt.Println("Введите сново слово:")
		fmt.Scan(&word)
	}

	fmt.Println("игра завершена")
}
