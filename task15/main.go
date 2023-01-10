package main

import "fmt"

func main() {
	fmt.Println("введите число, чтобы узнать уникальное меню:")
	var number int
	fmt.Scan(&number)

	if number == 1 {
		fmt.Println("день недели: понедельник.")
		fmt.Println("меню: сосиски и картофельное пюре")
	} else if number == 2 {
		fmt.Println("день недели: вторник.")
		fmt.Println("меню: жареная катошка с лисичками")
	} else if number == 3 {
		fmt.Println("день недели: среда.")
		fmt.Println("меню: макароны с мясом по-франчузски")
	} else if number == 4 {
		fmt.Println("день недели: четверг.")
		fmt.Println("меню: фуагра")
	} else if number == 5 {
		fmt.Println("день недели: пятница.")
		fmt.Println("меню: ризотто")
	} else if number == 6 {
		fmt.Println("день недели: суббота.")
		fmt.Println("меню: гребешки")
	} else if number == 7 {
		fmt.Println("день недели: воскресенье.")
		fmt.Println("меню: лимонный тарт")
	}
}
