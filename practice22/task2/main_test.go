package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Case1(t *testing.T) {
	input := []int{1, 2, 2, 2, 3}
	value := 2

	expected := 1
	expectedErr := error(nil)
	actual, actualErr := DefineIndex(input, value)

	assert.Equal(t, expected, actual)
	assert.Equal(t, expectedErr, actualErr)

}

func Test_Case2(t *testing.T) {
	input := []int{1, 2, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	value := 2

	expected := 1
	expectedErr := error(nil)
	actual, actualErr := DefineIndex(input, value)

	assert.Equal(t, expected, actual)
	assert.Equal(t, expectedErr, actualErr)

}

func Test_Case3(t *testing.T) {
	input := []int{1, 2, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	value := 11

	expected := 0
	expectedErr := error(ErrNoResult)
	actual, actualErr := DefineIndex(input, value)

	assert.Equal(t, expected, actual)
	assert.Equal(t, expectedErr, actualErr)

}
