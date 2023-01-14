package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Solve_Discriminant(t *testing.T) {
	a := 1.0
	b := -4.0
	c := -5.0

	expected1 := 5.0
	expected2 := -1.0
	expectedErr := error(nil)

	actual1, actual2, actualErr := Discriminant(a, b, c)

	assert.Equal(t, expected1, actual1)
	assert.Equal(t, expected2, actual2)
	assert.Equal(t, expectedErr, actualErr)
}
