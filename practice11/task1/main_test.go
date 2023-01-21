package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isCapital_Case1(t *testing.T) {
	word := "Go"

	expected := true
	actual := isCapital(word)

	assert.Equal(t, expected, actual)
}

func Test_isCapital_Case2(t *testing.T) {
	word := "is"

	expected := false
	actual := isCapital(word)

	assert.Equal(t, expected, actual)
}

func Test_isCapital_Case3(t *testing.T) {
	word := "cApital"

	expected := false
	actual := isCapital(word)

	assert.Equal(t, expected, actual)
}

func Test_isCapital_Case4(t *testing.T) {
	word := "CApital"

	expected := false
	actual := isCapital(word)

	assert.Equal(t, expected, actual)
}
