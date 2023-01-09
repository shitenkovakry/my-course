package main

import "fmt"

func main() {
	var login string
	var password string

	fmt.Print("Введите логин:")

	fmt.Scan(&login)

	fmt.Print("Введите пароль:")

	fmt.Scan(&password)

	fmt.Println("-----")
	fmt.Print("[", login, "],", " Вы успешно зашли!\n")
}
