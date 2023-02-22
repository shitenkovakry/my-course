package main

import (
	"fmt"
	"math"
	"strconv"
	"sync"
)

func ElevateSquareNumber(number int) int {
	squareOfNumber := int(math.Pow(float64(number), 2.0))

	return squareOfNumber
}

func MultiplyingTheSquareOfNumber(squareOfNumber int) int {
	multupliedSquareOfNumber := squareOfNumber * 2

	return multupliedSquareOfNumber
}

func process1(number, c chan int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	result1 := ElevateSquareNumber(<-number)
	c <- result1
	fmt.Println("квадрат:", result1)
}

func process2(number, c chan int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	result2 := MultiplyingTheSquareOfNumber(<-c)
	fmt.Println("произведение:", result2)
}

func main() {
	waitGroup := sync.WaitGroup{}
	number := make(chan int, 1)
	c := make(chan int, 1)

	for {

		fmt.Println("введите число:")

		var input string

		_, err := fmt.Scan(&input)
		if err != nil {
			panic(err)
		}

		if input == "stop" {
			break
		}

		waitGroup.Add(1)
		go process1(number, c, &waitGroup)

		waitGroup.Add(1)
		go process2(number, c, &waitGroup)

		n, err := strconv.Atoi(input)
		if err != nil {
			panic(err)
		}

		number <- n
	}

	fmt.Println("done")
	waitGroup.Wait()
}
