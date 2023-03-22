package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Case1(t *testing.T) {
	arrayOfAlise := []int{1, 2, 3}
	arrayOfBob := []int{4, 1, 5}

	expected := []int{1, 2}
	actual := GiveComprasionResult(arrayOfAlise, arrayOfBob)

	assert.Equal(t, expected, actual)
}

func Test_Case2(t *testing.T) {
	arrayOfAlise := []int{1, 1, 1}
	arrayOfBob := []int{1, 1, 1}

	expected := []int{0, 0}
	actual := GiveComprasionResult(arrayOfAlise, arrayOfBob)

	assert.Equal(t, expected, actual)
}

func Test_Case3(t *testing.T) {
	arrayOfAlise := []int{3, 6, 7, 0, 10}
	arrayOfBob := []int{3, 5, 2, 8, 9}

	expected := []int{3, 1}
	actual := GiveComprasionResult(arrayOfAlise, arrayOfBob)

	assert.Equal(t, expected, actual)
}
