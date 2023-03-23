package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Case1(t *testing.T) {
	array := []int{1, 3, 5, 7, 9}

	expected1 := 16
	expected2 := 24
	actual1, actual2 := CalculateMaxAndMin(array)

	assert.Equal(t, expected1, actual1)
	assert.Equal(t, expected2, actual2)
}

func Test_Case4(t *testing.T) {
	array := []int{1, 3, 5, 7, 9}

	expected := 1
	actual := SearchMin(array)

	assert.Equal(t, expected, actual)
}

func Test_Case5(t *testing.T) {
	array := []int{1, 3, 5, 7, 9}

	expected := 9
	actual := SearchMax(array)

	assert.Equal(t, expected, actual)
}

func Test_Case6(t *testing.T) {
	array := []int{156873294, 719583602, 581240736, 605827319, 895647130}

	expected1 := 2063524951
	expected2 := 2802298787
	actual1, actual2 := CalculateMaxAndMin(array)

	assert.Equal(t, expected1, actual1)
	assert.Equal(t, expected2, actual2)
}

func Test_Case7(t *testing.T) {
	array := []int{5, 5, 5, 5, 5}

	expected1 := 20
	expected2 := 20
	actual1, actual2 := CalculateMaxAndMin(array)

	assert.Equal(t, expected1, actual1)
	assert.Equal(t, expected2, actual2)
}
