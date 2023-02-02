package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_vectorXvector_case1(t *testing.T) {
	vector1 := []int{1, 2, 3}
	vector2 := []int{7, 9, 11}

	expected := 58
	actual := VectorXVector(vector1, vector2)

	assert.Equal(t, expected, actual)
}

func Test_Case1(t *testing.T) {
	input := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
	}

	input2 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
		{17, 18, 19, 20},
	}

	expected := [][]int{
		{175, 190, 205, 220},
		{400, 440, 480, 520},
		{625, 690, 755, 820},
	}

	actual := CalculateMatrix(input, input2)

	assert.Equal(t, expected, actual)
}

func Test_obtainColumn_case1(t *testing.T) {
	inputColumn := 2
	inputMatrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
		{17, 18, 19, 20},
	}

	expected := []int{3, 7, 11, 15, 19}

	actual := ObtainColumnFromMatrix(inputColumn, inputMatrix)
	assert.Equal(t, expected, actual)
}

func Test_obtainRow_case1(t *testing.T) {
	inputRow := 2
	inputMatrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
		{17, 18, 19, 20},
	}

	expected := []int{9, 10, 11, 12}

	actual := ObtainRowFromMatrix(inputRow, inputMatrix)
	assert.Equal(t, expected, actual)
}
