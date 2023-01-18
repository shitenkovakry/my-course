package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Case1(t *testing.T) {
	number1 := int64(-2000)
	number2 := int64(100)

	expected := maxInt32str
	actual := WhichDataType(number1, number2)

	assert.Equal(t, expected, actual)
}

func Test_Case2(t *testing.T) {
	number1 := int64(2000)
	number2 := int64(100)

	expected := maxUint32str
	actual := WhichDataType(number1, number2)

	assert.Equal(t, expected, actual)
}

func Test_Case3(t *testing.T) {
	number1 := int64(1)
	number2 := int64(230)

	expected := maxUint8str
	actual := WhichDataType(number1, number2)

	assert.Equal(t, expected, actual)
}

func Test_Case4(t *testing.T) {
	number1 := int64(1)
	number2 := int64(-127)

	expected := maxInt8str
	actual := WhichDataType(number1, number2)

	assert.Equal(t, expected, actual)
}

func Test_Case5(t *testing.T) {
	number1 := int64(-1)
	number2 := int64(32768)

	expected := maxInt16str
	actual := WhichDataType(number1, number2)

	assert.Equal(t, expected, actual)
}

func Test_Case6(t *testing.T) {
	number1 := int64(-1)
	number2 := int64(128)

	expected := maxInt8str
	actual := WhichDataType(number1, number2)

	assert.Equal(t, expected, actual)
}

func Test_Case7(t *testing.T) {
	number1 := int64(1)
	number2 := int64(65535)

	expected := maxUint16str
	actual := WhichDataType(number1, number2)

	assert.Equal(t, expected, actual)
}

func Test_Case8(t *testing.T) {
	number1 := int64(1)
	number2 := int64(4294967295)

	expected := maxUint32str
	actual := WhichDataType(number1, number2)

	assert.Equal(t, expected, actual)
}

func Test_Case9(t *testing.T) {
	number1 := int64(1)
	number2 := int64(14294967295)

	expected := maxUint64str
	actual := WhichDataType(number1, number2)

	assert.Equal(t, expected, actual)
}

func Test_Case10(t *testing.T) {
	number1 := int64(-1)
	number2 := int64(14294967295)

	expected := maxInt64str
	actual := WhichDataType(number1, number2)

	assert.Equal(t, expected, actual)
}
