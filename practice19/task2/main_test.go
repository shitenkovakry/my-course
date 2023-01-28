package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Case1(t *testing.T) {
	array := []int{7, 2, 15, 3, 36, 1}

	expected := []int{1, 2, 3, 7, 15, 36}
	actual := BubbleSort(array)

	assert.Equal(t, expected, actual)
}

func Test_Case2(t *testing.T) {
	array := []int{7, 1}

	expected := []int{1, 7}
	actual := BubbleSort(array)

	assert.Equal(t, expected, actual)
}

func Test_Case3(t *testing.T) {
	array := []int{7, 5, 1}

	expected := []int{1, 5, 7}
	actual := BubbleSort(array)

	assert.Equal(t, expected, actual)
}
