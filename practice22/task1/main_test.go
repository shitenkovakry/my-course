package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Case1(t *testing.T) {
	input := []int{5, 9, 1, 2}
	value := 1

	expected := 1
	actual := FindNumberFromArray(input, value)

	assert.Equal(t, expected, actual)

}
