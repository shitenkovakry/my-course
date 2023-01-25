package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Case1(t *testing.T) {
	number := 2

	expected := true
	actual := IsEven(number)

	assert.Equal(t, expected, actual)
}

func Test_Case2(t *testing.T) {
	number := 7

	expected := false
	actual := IsEven(number)

	assert.Equal(t, expected, actual)
}

func Test_Case3(t *testing.T) {
	number := -6

	expected := true
	actual := IsEven(number)

	assert.Equal(t, expected, actual)
}

func Test_Case4(t *testing.T) {
	number := -1

	expected := false
	actual := IsEven(number)

	assert.Equal(t, expected, actual)
}

func Test_Case5(t *testing.T) {
	number := 0

	expected := true
	actual := IsEven(number)

	assert.Equal(t, expected, actual)
}
