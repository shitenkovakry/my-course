package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Case1(t *testing.T) {
	array := []int{1, 1, 0, -1, -1}

	expected1 := 0.400000
	expected2 := 0.200000
	expected3 := 0.400000

	actual1, actual2, actual3 := CalculateDivisionOfNumberOfArrayFromLen(array)

	assert.Equal(t, expected1, actual1)
	assert.Equal(t, expected2, actual2)
	assert.Equal(t, expected3, actual3)
}

func Test_Case2(t *testing.T) {
	array := []int{-4, 3, -9, 0, 4, 1}

	expected1 := 0.500000
	expected2 := 0.16666666666666666
	expected3 := 0.3333333333333333

	actual1, actual2, actual3 := CalculateDivisionOfNumberOfArrayFromLen(array)

	assert.Equal(t, expected1, actual1)
	assert.Equal(t, expected2, actual2)
	assert.Equal(t, expected3, actual3)
}
