package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Case1(t *testing.T) {
	input := []int{5, 9, 1, 2}
	value := 1

	expected := 1
	expectedErr := error(nil)
	actual, actualErr := FindNumberFromArray(input, value)

	assert.Equal(t, expected, actual)
	assert.Equal(t, expectedErr, actualErr)

}

func Test_Case2(t *testing.T) {
	input := []int{5, 9, 1, 2}
	value := 6

	expected := 0
	expectedErr := error(ErrNoResult)
	actual, actualErr := FindNumberFromArray(input, value)

	assert.Equal(t, expected, actual)
	assert.Equal(t, expectedErr, actualErr)

}
