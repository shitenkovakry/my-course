package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_case0(t *testing.T) {
	input := "Martin 20 1"

	expectedName := "Martin"
	expectedAge := 20
	expectedGrade := 1

	actualName, actualAge, actualGrade, err := ConvertStudentStr(input)

	assert.Nil(t, err)
	assert.Equal(t, expectedAge, actualAge)
	assert.Equal(t, expectedGrade, actualGrade)
	assert.Equal(t, expectedName, actualName)
}

func Test_case1(t *testing.T) {
	input := "Martin 20 1"

	expected := &Student{
		Name:  "Martin",
		Age:   20,
		Grade: 1,
	}

	actual, err := ConvertStudent(input)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}

func Test_Case1(t *testing.T) {
	inputLines := `Martin 20 1
Mikhail 30 2



Ondrys 38 1
`

	expected := map[string]*Student{
		"Martin": {
			Name:  "Martin",
			Age:   20,
			Grade: 1,
		},
		"Mikhail": {
			Name:  "Mikhail",
			Age:   30,
			Grade: 2,
		},
		"Ondrys": {
			Name:  "Ondrys",
			Age:   38,
			Grade: 1,
		},
	}

	actual, err := ConvertLinesToMapStudents(inputLines)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
