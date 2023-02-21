package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
func Test_Case1(t *testing.T) {
	input := []string{"Hello", "miu"}
	input2 := []rune{'H', 'L'}

	expected := 0
	expected2 := 3

	actual, actual2 := ParseTest(input, input2)

	assert.Equal(t, expected, actual)
	assert.Equal(t, expected2, actual2)
}
*/

func Test_case_mini_1(t *testing.T) {
	inputStr := "Hello world"
	inputRunes := []rune{'H'}
	expected := []int{0}

	actual := TestStringAgainstRune(inputStr, inputRunes)
	assert.Equal(t, expected, actual)
}

func Test_case_mini_2(t *testing.T) {
	inputStr := "Hello world"
	inputRunes := []rune{'l'}
	expected := []int{9}

	actual := TestStringAgainstRune(inputStr, inputRunes)
	assert.Equal(t, expected, actual)
}

func Test_case_mini_3(t *testing.T) {
	inputStr := "Hello world"
	inputRunes := []rune{'k'}
	expected := []int{-1}

	actual := TestStringAgainstRune(inputStr, inputRunes)
	assert.Equal(t, expected, actual)
}

func Test_case_mini_4(t *testing.T) {
	inputStr := "Hello world"
	inputRunes := []rune{'H', 'l', 'k'}
	expected := []int{0, 9, -1}

	actual := TestStringAgainstRune(inputStr, inputRunes)
	assert.Equal(t, expected, actual)
}

func Test_case_mini_5(t *testing.T) {
	inputStr := "Hello world"
	inputRunes := []rune{'h', 'k', 'l', 'j'}
	expected := []int{0, -1, 9, -1}

	actual := TestStringAgainstRune(inputStr, inputRunes)
	assert.Equal(t, expected, actual)
}

func Test_Case1(t *testing.T) {
	inputArrayOfStrings := []string{"Hello world", "miu"}
	inputRunes := []rune{'H', 'w'}

	expected := [][]int{
		{0, 6},
		{-1, -1},
	}

	actual := ParseTest(inputArrayOfStrings, inputRunes)

	assert.Equal(t, expected, actual)
}
