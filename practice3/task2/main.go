package main

import "fmt"

func main() {
	var arrivalStop string
	var passengersOnBus int
	var passengersEntered1 int
	var passengersEntered2 int
	var passengersEntered3 int
	var passengersEntered int
	var passengersExited int

	fmt.Print("Прибываем на остановку ")
	fmt.Scan(&arrivalStop)

	fmt.Print("В салоне пассажиров: ")
	fmt.Scan(&passengersOnBus)

	fmt.Print("Сколько пассажиров зашло на остановке: ")
	fmt.Scan(&passengersEntered1)

	fmt.Print("Отпрввляемся с остановки ")
	fmt.Scan(&arrivalStop)
	fmt.Println()

	fmt.Print("Прибываем на остановку ")
	fmt.Scan(&arrivalStop)

	fmt.Print("В салоне пассажиров: ")
	fmt.Scan(&passengersOnBus)

	fmt.Print("Сколько пассажиров вышло на остановке: ")
	fmt.Scan(&passengersExited)

	fmt.Print("Сколько пассажиров зашло на остановке: ")
	fmt.Scan(&passengersEntered2)

	fmt.Print("Отправляемся с остановки ")
	fmt.Scan(&arrivalStop)
	fmt.Println()

	fmt.Print("Прибываем на остановку ")
	fmt.Scan(&arrivalStop)

	fmt.Print("В салоне пассажиров: ")
	fmt.Scan(&passengersOnBus)

	fmt.Print("Сколько пассажиров вышло на остановке: ")
	fmt.Scan(&passengersExited)

	fmt.Print("Сколько пассажиров зашло на остановке: ")
	fmt.Scan(&passengersEntered3)

	fmt.Print("Отправляемся с остановки ")
	fmt.Scan(&arrivalStop)
	fmt.Println()

	fmt.Print("Прибываем на остановку ")
	fmt.Scan(&arrivalStop)

	fmt.Print("В салоне пассажиров: ")
	fmt.Scan(&passengersOnBus)

	fmt.Println("Конечная, все пассажиры вышли")
	fmt.Println()

	ticketPrice := 20
	passengersEntered = passengersEntered1 + passengersEntered2 + passengersEntered3
	totalSum := ticketPrice * passengersEntered
	driverSalary := totalSum / 4
	fuel := totalSum / 5
	taxes := totalSum / 5
	fixes := totalSum / 5

	income := totalSum

	fmt.Println("Всего заработано", totalSum, "руб.")
	fmt.Println("Зарплата водителя:", driverSalary, "руб.")
	fmt.Println("Расходы на топливо:", fuel, "руб.")
	fmt.Println("Налоги", taxes, "руб.")
	fmt.Println("Расходы на ремонт машины:", fixes, "руб.")
	fmt.Println("Итого доход:", income, "руб.")

}
