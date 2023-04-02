package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Case1(t *testing.T) {
	inputNumbre := 47

	expected := []int{1, 0, 1, 1, 1, 1}
	actual := ConvertToBinaryRepresentation(inputNumbre)

	assert.Equal(t, expected, actual)
}
