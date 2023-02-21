package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Case1(t *testing.T) {
	inputString := "привет"
	partForCheck := "вет"

	expected := true
	actual := CheckPartOfLine(inputString, partForCheck)

	assert.Equal(t, expected, actual)
}
