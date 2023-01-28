package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Case1(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	expexted := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	actual := ReverseArray(input)

	assert.Equal(t, expexted, actual)
}
