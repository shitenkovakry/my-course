package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Case1(t *testing.T) {
	inputNumber := 3

	expected := 9
	actual := NumberSquared(inputNumber)

	assert.Equal(t, expected, actual)
}

func Test_Case2(t *testing.T) {
	inputNumber := -3

	expected := 9
	actual := NumberSquared(inputNumber)

	assert.Equal(t, expected, actual)
}

func Test_Case3(t *testing.T) {
	inputNumber := 0

	expected := 0
	actual := NumberSquared(inputNumber)

	assert.Equal(t, expected, actual)
}
