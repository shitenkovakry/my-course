package main

import "fmt"

func CalculateS(x int16, y uint8, z float32) float32 {
	s := 2*float32(x) + float32(y)*2 - 3/z

	return s
}

func main() {
	x := int16(-2)
	y := uint8(6)
	z := float32(1.5)

	result := CalculateS(x, y, z)

	fmt.Println(result)
}
