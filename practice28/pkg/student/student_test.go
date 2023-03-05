package student

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
