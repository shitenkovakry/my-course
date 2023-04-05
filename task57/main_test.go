package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Case1(t *testing.T) {
	inputArray := []int{73, 67, 38, 33}

	expected := []int{75, 67, 40, 33}
	actual := GradeStudents(inputArray)

	assert.Equal(t, expected, actual)
}
