package main

import (
	"fmt"
)

func VectorXVector(vector1, vector2 []int) int {
	summaOfVectors := 0

	for i := 0; i < len(vector1) || i < len(vector2); i++ {
		summaOfVectors += (vector1[i] * vector2[i])
	}

	return summaOfVectors
}

func ObtainColumnFromMatrix(inputColumn int, inputMatrix [][]int) []int {
	newArray := []int{}

	for i := 0; i < len(inputMatrix); i++ {
		row := inputMatrix[i]
		value := row[inputColumn]
		newArray = append(newArray, value)
	}

	return newArray
}

func ObtainRowFromMatrix(inputRow int, inputMatrix [][]int) []int {
	return inputMatrix[inputRow]
}

func CalculateMatrix(matrix1, matrix2 [][]int) [][]int {
	matrixMultiplication := [][]int{}

	numberOfRows := len(matrix1)
	numberOfColumns := len(matrix2[0])

	for i := 0; i < numberOfRows; i++ {
		row := []int{}
		for j := 0; j < numberOfColumns; j++ {
			row = append(row, 0)
		}

		matrixMultiplication = append(matrixMultiplication, row)
	}

	for i := 0; i < numberOfRows; i++ {
		row := ObtainRowFromMatrix(i, matrix1)

		for j := 0; j < numberOfColumns; j++ {
			column := ObtainColumnFromMatrix(j, matrix2)

			valueIJ := VectorXVector(row, column)
			matrixMultiplication[i][j] = valueIJ
		}
	}

	return matrixMultiplication
}

func main() {
	matrix1 := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
	}

	matrix2 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
		{17, 18, 19, 20},
	}

	result := CalculateMatrix(matrix1, matrix2)

	fmt.Println(result)
}
