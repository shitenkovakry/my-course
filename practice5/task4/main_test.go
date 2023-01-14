package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Coin_Case1(t *testing.T) {
	costOfProduct := 1000
	denominationCoin1 := 10
	denominationCoin2 := 0
	denominationCoin3 := 0
	expected := true

	actual := CanPayWithoutShortChange(costOfProduct, denominationCoin1, denominationCoin2, denominationCoin3)

	assert.Equal(t, expected, actual)
}

func Test_Coin_Case2(t *testing.T) {
	costOfProduct := 1001
	denominationCoin1 := 10
	denominationCoin2 := 0
	denominationCoin3 := 0
	expected := false

	actual := CanPayWithoutShortChange(costOfProduct, denominationCoin1, denominationCoin2, denominationCoin3)

	assert.Equal(t, expected, actual)
}

func Test_Coin_Case3(t *testing.T) {
	costOfProduct := 1000
	denominationCoin1 := 11
	denominationCoin2 := 20
	denominationCoin3 := 0
	expected := true

	actual := CanPayWithoutShortChange(costOfProduct, denominationCoin1, denominationCoin2, denominationCoin3)

	assert.Equal(t, expected, actual)
}

func Test_Coin_Case4(t *testing.T) {
	costOfProduct := 1000
	denominationCoin1 := 11
	denominationCoin2 := 21
	denominationCoin3 := 50
	expected := true

	actual := CanPayWithoutShortChange(costOfProduct, denominationCoin1, denominationCoin2, denominationCoin3)

	assert.Equal(t, expected, actual)
}

func Test_Coin_Case5(t *testing.T) {
	costOfProduct := 90
	denominationCoin1 := 50
	denominationCoin2 := 20
	denominationCoin3 := 0
	expected := true

	actual := CanPayWithoutShortChange(costOfProduct, denominationCoin1, denominationCoin2, denominationCoin3)

	assert.Equal(t, expected, actual)
}

func Test_Coin_Case6(t *testing.T) {
	costOfProduct := 90
	denominationCoin1 := 20
	denominationCoin2 := 50
	denominationCoin3 := 0
	expected := true

	actual := CanPayWithoutShortChange(costOfProduct, denominationCoin1, denominationCoin2, denominationCoin3)

	assert.Equal(t, expected, actual)
}

func Test_Coin_Case7(t *testing.T) {
	costOfProduct := 90
	denominationCoin1 := 0
	denominationCoin2 := 50
	denominationCoin3 := 20
	expected := true

	actual := CanPayWithoutShortChange(costOfProduct, denominationCoin1, denominationCoin2, denominationCoin3)

	assert.Equal(t, expected, actual)
}

func Test_Coin_Case8(t *testing.T) {
	costOfProduct := 90
	denominationCoin1 := 0
	denominationCoin2 := 20
	denominationCoin3 := 50
	expected := true

	actual := CanPayWithoutShortChange(costOfProduct, denominationCoin1, denominationCoin2, denominationCoin3)

	assert.Equal(t, expected, actual)
}

func Test_Coin_Case9(t *testing.T) {
	costOfProduct := 97
	denominationCoin1 := 50
	denominationCoin2 := 20
	denominationCoin3 := 7
	expected := true

	actual := CanPayWithoutShortChange(costOfProduct, denominationCoin1, denominationCoin2, denominationCoin3)

	assert.Equal(t, expected, actual)
}
