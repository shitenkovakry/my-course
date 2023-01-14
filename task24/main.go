package main

import "fmt"

const (
	EnoughMessage = "Сообщений достаточно"
)

func main() {
	fmt.Println("Введите сообщение:")
	var message string
	fmt.Scan(&message)

	i := 0

	fmt.Println("Начинаем")

	for i < 3 {
		fmt.Println(message)
		i++
	}

	fmt.Println(EnoughMessage)

}
