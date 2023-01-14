package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Case1(t *testing.T) {
	price := 3000.0
	discount := 20.0

	expected := 600.0

	actual := ComputeDiscount(price, discount)

	assert.Equal(t, expected, actual)
}

func Test_Case2(t *testing.T) {
	price := 30000.0
	discount := 20.0

	expected := 2000.0

	actual := ComputeDiscount(price, discount)

	assert.Equal(t, expected, actual)
}

func Test_Case3(t *testing.T) {
	price := 1000000.0
	discount := 20.0

	expected := 2000.0

	actual := ComputeDiscount(price, discount)

	assert.Equal(t, expected, actual)
}

func Test_Case4(t *testing.T) {
	price := 1000.0
	discount := 70.0

	expected := 300.0

	actual := ComputeDiscount(price, discount)

	assert.Equal(t, expected, actual)
}
