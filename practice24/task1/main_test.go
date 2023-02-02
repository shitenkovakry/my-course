package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Case1(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	expected := []int{2, 4, 6, 8, 10}
	expected2 := []int{1, 3, 5, 7, 9}

	actual, actual2 := ReturnArraysOfOddAndEven(input)

	assert.Equal(t, expected, actual)
	assert.Equal(t, expected2, actual2)
}
