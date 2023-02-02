package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Case1(t *testing.T) {
	input := []int{10, 3, 5}

	expected := []int{3, 5, 10}
	actual := SortByMinimums(input)

	assert.Equal(t, expected, actual)
}
