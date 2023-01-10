package main

import "fmt"

func main() {
	fmt.Println("Баллы ЕГЭ.")

	fmt.Println("Введите результат первого экзамена:")
	var resultFirst int
	fmt.Scan(&resultFirst)

	fmt.Println("Введите результат второго экзамена:")
	var resultSecond int
	fmt.Scan(&resultSecond)

	fmt.Println("Введите результат третьего экзамена:")
	var resultThird int
	fmt.Scan(&resultThird)

	amountOfPasses := 275

	totalResult := resultFirst + resultSecond + resultThird

	if totalResult >= amountOfPasses {
		fmt.Println("Сумма проходных баллов:", amountOfPasses)
		fmt.Println("Количество набранных баллов:", totalResult)
		fmt.Println("Вы поступили!")
		return
	}

	fmt.Println("Сумма проходных баллов:", amountOfPasses)
	fmt.Println("Количество набранных баллов", totalResult)
	fmt.Println("Вы не поступили:(")

}
