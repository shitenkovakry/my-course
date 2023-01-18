package main

import "fmt"

func main() {
	fmt.Println("Дни недели.")

	fmt.Println("Введите будний день недели: пн,вт,ср, чт, пт")
	var weekDay string
	fmt.Scan(&weekDay)

	switch weekDay {
	case "пн":
		fmt.Println("понедельник")
	case "вт":
		fmt.Println("вторник")
	case "ср":
		fmt.Println("среда")
	case "чт":
		fmt.Println("четверг")
	case "пт":
		fmt.Println("пятница")
	case "сб", "вс":
		fmt.Println("это выходной день, просили ввести будний. введите будний день")
	default:
		fmt.Println("нет такого дня недели")
	}
}
