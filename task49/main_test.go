package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Case1(t *testing.T) {
	array := []int{10001, 10002, 10003, 10004, 10005}

	expected := 50015
	actual := CalculateSummaOfArray(array)

	assert.Equal(t, expected, actual)
}

func Test_Case2(t *testing.T) {
	array := []int{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}

	expected := 5000000015
	actual := CalculateSummaOfArray(array)

	assert.Equal(t, expected, actual)
}
