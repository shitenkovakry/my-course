package main

import "fmt"

func CalculateDeterminant3x3(matrix [][]int) int {
	left := ((matrix[0][0] * matrix[1][1] * matrix[2][2]) +
		(matrix[1][0] * matrix[2][1] * matrix[0][2]) +
		(matrix[2][0] * matrix[0][1] * matrix[1][2]))

	right := ((matrix[0][2] * matrix[1][1] * matrix[2][0]) +
		(matrix[1][2] * matrix[2][1] * matrix[0][0]) +
		(matrix[2][2] * matrix[0][1] * matrix[1][0]))

	determinant := left - right

	return determinant
}

func main() {
	matrix := [][]int{
		{4, 3, 1},
		{8, 6, -2},
		{5, -1, 2},
	}

	result := CalculateDeterminant3x3(matrix)

	fmt.Println(result)
}
