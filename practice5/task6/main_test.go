package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Is_Lucky_Ticket(t *testing.T) {
	ticket := 3425
	expected := luckyTicket

	actual := CheckTicket(ticket)

	assert.Equal(t, expected, actual)
}

func Test_Is_Mirror_Ticket(t *testing.T) {
	ticket := 1221
	expected := mirrorTicket

	actual := CheckTicket(ticket)

	assert.Equal(t, expected, actual)
}

func Test_Is_Regular_Ticket(t *testing.T) {
	ticket := 1234
	expected := regularTicket

	actual := CheckTicket(ticket)

	assert.Equal(t, expected, actual)
}

func Test_ConvertDigits1(t *testing.T) {
	ticket := 1234
	expectedDigit1 := 1
	expectedDigit2 := 2
	expectedDigit3 := 3
	expectedDigit4 := 4

	actual1, actual2, actual3, actual4 := ConvertInto4Digits(ticket)

	assert.Equal(t, expectedDigit1, actual1)
	assert.Equal(t, expectedDigit2, actual2)
	assert.Equal(t, expectedDigit3, actual3)
	assert.Equal(t, expectedDigit4, actual4)
}

func Test_ConvertDigits2(t *testing.T) {
	ticket := 9
	expectedDigit1 := 0
	expectedDigit2 := 0
	expectedDigit3 := 0
	expectedDigit4 := 9

	actual1, actual2, actual3, actual4 := ConvertInto4Digits(ticket)

	assert.Equal(t, expectedDigit1, actual1)
	assert.Equal(t, expectedDigit2, actual2)
	assert.Equal(t, expectedDigit3, actual3)
	assert.Equal(t, expectedDigit4, actual4)
}
