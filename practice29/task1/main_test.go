package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Case1(t *testing.T) {
	number := 3

	expected := 9
	actual := ElevateSquareNumber(number)

	assert.Equal(t, expected, actual)
}

func Test_Case2(t *testing.T) {
	squareOfNumber := 9

	expected := 18
	actual := MultiplyingTheSquareOfNumber(squareOfNumber)

	assert.Equal(t, expected, actual)
}
