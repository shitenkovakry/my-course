package main

import "fmt"

func main() {
	var userName string
	var password string
	var age int

	fmt.Scan(&userName)
	fmt.Scan(&password)
	fmt.Scan(&age)

	fmt.Println("Поздравляю,", userName, ", теперь Вы зарегестрированы!")
	fmt.Println("Ваш пароль:", password)
	fmt.Println("Ваш возраст:", age)
}
