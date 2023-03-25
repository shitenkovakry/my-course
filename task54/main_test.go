package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Case1(t *testing.T) {
	array := []int{1, 2, 3, 4, 4}

	expected := 4
	actual := FindMaxNumberFromArray(array)

	assert.Equal(t, expected, actual)
}

func Test_Case2(t *testing.T) {
	array := []int{1, 2, 3, 4, 4}

	expected := 2
	actual := CalculateCountOfMaxNumber(array)

	assert.Equal(t, expected, actual)
}
