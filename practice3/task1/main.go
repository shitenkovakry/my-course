package main

import "fmt"

func main() {
	var planet string
	var star string
	var ranger string
	var money int

	fmt.Println("Как тебя зовут?")
	fmt.Scan(&ranger)

	fmt.Println("Название планеты?")
	fmt.Scan(&planet)

	fmt.Println("Звездная система?")
	fmt.Scan(&star)

	fmt.Println("Сумма вознагрождения?")
	fmt.Scan(&money)

	fmt.Print("Как вам,", ranger, ", известно, мы - раса мирная, поэтому на наших военных кораблях используются наемники с других планет. Система набора отработана давным-давно.")
	fmt.Println("Обычно мы приглашаем на наши корабли людей с планеты", planet, "системы", star, ".")
	fmt.Println()
	fmt.Print("Но случилась неприятность - в связи с большими потерями, в последнее время, престиж профессии сильно упал, и теперь не так-то просто завербовать призывников.")
	fmt.Println("Главный комиссар планеты", planet, ",впрочем, отлично справлялся, но недавно его наградили орденом Сутулого с закруткой на спине, и его на радостях парализовало! Призыв под угрозой срыва!")
	fmt.Println()
	fmt.Print(ranger, ", вы должны прибыть на планету ", planet, " системы ", star, " и помочь выполнить план призыва.")
	fmt.Print("Мы готовы выплатить вам премию в ", money, " кредитов за эту маленькую услугу.")

}
