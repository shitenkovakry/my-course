package main

import "fmt"

func main() {
	fmt.Println("Бамбук против вредителей")

	startHeight := 100
	growth := 50
	losses := 20

	finalGrowth := growth*2 + growth/2
	finalLosses := losses * 2

	finalHeight := startHeight + finalGrowth - finalLosses

	fmt.Println("Высота бамбука к середине третьего дня: ", finalHeight)

}
