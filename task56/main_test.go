package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Case1(t *testing.T) {
	inputNumber := 75

	expected := 75
	actual := GradeStudent(inputNumber)

	assert.Equal(t, expected, actual)
}

func Test_Case2(t *testing.T) {
	inputNumber := 73

	expected := 75
	actual := GradeStudent(inputNumber)

	assert.Equal(t, expected, actual)
}

func Test_Case3(t *testing.T) {
	inputNumber := 57

	expected := 57
	actual := GradeStudent(inputNumber)

	assert.Equal(t, expected, actual)
}

func Test_Case4(t *testing.T) {
	inputNumber := 29

	expected := 29
	actual := GradeStudent(inputNumber)

	assert.Equal(t, expected, actual)
}
