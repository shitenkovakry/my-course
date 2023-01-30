package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Case1(t *testing.T) {
	input := []string{
		"15",
		"15.7",
		"15.7.1",
		"15.7.2",
		"15.7.3",
		"15.2",
		"15.2.1",
		"15.2.2",
		"11",
		"11.1",
		"11.1.1",
		"11.1.3",
		"11.1.2",
		"11.2",
		"11.2.1",
		"11.2.2",
	}

	expected := []string{
		"11",
		"11.1",
		"11.1.1",
		"11.1.2",
		"11.1.3",
		"11.2",
		"11.2.1",
		"11.2.2",
		"15",
		"15.2",
		"15.2.1",
		"15.2.2",
		"15.7",
		"15.7.1",
		"15.7.2",
		"15.7.3",
	}

	actual := Organize(input)

	assert.Equal(t, expected, actual)

}

func Test_Case2(t *testing.T) {
	input := []string{
		"15",
		"15.7",
		"15.7.1",
		"15.7.2",
		"15.7.3",
		"15.2",
		"15.2.1",
		"15.2.2",
		"9",
		"9.1",
		"9.1.1",
		"9.1.3",
		"9.1.2",
		"9.2",
		"9.2.1",
		"9.2.2",
	}

	expected := []string{
		"9",
		"9.1",
		"9.1.1",
		"9.1.2",
		"9.1.3",
		"9.2",
		"9.2.1",
		"9.2.2",
		"15",
		"15.2",
		"15.2.1",
		"15.2.2",
		"15.7",
		"15.7.1",
		"15.7.2",
		"15.7.3",
	}

	actual := Organize(input)

	assert.Equal(t, expected, actual)

}

func Test_Case4(t *testing.T) {
	left := "15.2"
	right := "9.1"

	expected := false

	actual := CompareTwoKeys(left, right)

	assert.Equal(t, expected, actual)
}

func Test_Case5(t *testing.T) {
	left := "9.1"
	right := "15.2"

	expected := true

	actual := CompareTwoKeys(left, right)

	assert.Equal(t, expected, actual)
}

func Test_Case8(t *testing.T) {
	left := "9"
	right := "9.1.1"

	expected := true

	actual := CompareTwoKeys(left, right)

	assert.Equal(t, expected, actual)
}

func Test_Case9(t *testing.T) {
	left := "9"
	right := "15.5.3"

	expected := true

	actual := CompareTwoKeys(left, right)

	assert.Equal(t, expected, actual)
}
