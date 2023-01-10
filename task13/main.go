package main

import "fmt"

func main() {
	fmt.Println("Напишите Ваши очки опыта:")
	var experiencePoints int
	fmt.Scan(&experiencePoints)

	if experiencePoints >= 5000 {
		fmt.Println("поздравляем! Вы достигли 4-го уровня")

		return
	}

	if experiencePoints >= 2500 {
		fmt.Println("поздравляем! Вы достигли 3-го уровня")

		return
	}

	if experiencePoints >= 1000 {
		fmt.Println("поздравляем! Вы достигли 2-го уровня")

		return
	}

	fmt.Println("вы на первом уровне")
}
