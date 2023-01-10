package main

import "fmt"

func main() {
	fmt.Println("IQ первого космонавта:")
	var iq1 int
	fmt.Scan(&iq1)

	fmt.Println("iq второго космонавта:")
	var iq2 int
	fmt.Scan(&iq2)

	fmt.Println("iq третьего космонавта:")
	var iq3 int
	fmt.Scan(&iq3)

	max := iq1

	if max < iq2 {
		max = iq2
	} 
	
	if max < iq3 {
		max = iq3
	}

	fmt.Println(max)

}
