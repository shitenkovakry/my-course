package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Case1(t *testing.T) {
	number1 := 2
	number2 := 9

	expected := 20
	actual := CountevenNumbers(number1, number2)

	assert.Equal(t, expected, actual)
}

func Test_Case2(t *testing.T) {
	number1 := 0
	number2 := 5

	expected := 6
	actual := CountevenNumbers(number1, number2)

	assert.Equal(t, expected, actual)
}

func Test_Case3(t *testing.T) {
	number1 := -6
	number2 := -1

	expected := -12
	actual := CountevenNumbers(number1, number2)

	assert.Equal(t, expected, actual)
}

func Test_Case4(t *testing.T) {
	number1 := 6
	number2 := 2

	expected := 12
	actual := CountevenNumbers(number1, number2)

	assert.Equal(t, expected, actual)
}

func Test_Case5(t *testing.T) {
	number1 := 100
	number2 := 100

	expected := 100
	actual := CountevenNumbers(number1, number2)

	assert.Equal(t, expected, actual)
}

func Test_Case6(t *testing.T) {
	number1 := 99
	number2 := 99

	expected := 0
	actual := CountevenNumbers(number1, number2)

	assert.Equal(t, expected, actual)
}
