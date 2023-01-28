package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Case1(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	input2 := []int{6, 7, 8, 9, 10}

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	actual := MergerArrays(input, input2)

	assert.Equal(t, expected, actual)
}
