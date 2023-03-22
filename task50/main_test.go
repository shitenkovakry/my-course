package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Case1(t *testing.T) {
	array := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}

	expected := 2
	actual := CalculateSumOfNumbersDiagonally(array)

	assert.Equal(t, expected, actual)
}
