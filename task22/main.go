package main

import "fmt"

func main() {
	fmt.Println("напишите являетесь ли вы пользователем сайта")
	var isUser bool
	fmt.Scan(&isUser)

	if !isUser {
		fmt.Println("нужна регистрация")
	}
}
