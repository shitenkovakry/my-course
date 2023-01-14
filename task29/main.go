package main

import "fmt"

func main() {
	number := -1

	for {
		number++
		if number > 10 {
			break
		}

		if number%2 != 0 {
			continue
		}

		fmt.Println(number, "кратен 2")
	}
}

/*func main1() {
	for number := 0; number <= 10; number++ {
		if number%2 != 0 {
			continue
		}

		fmt.Println(number, "кратен 2")
	}
}
*/
