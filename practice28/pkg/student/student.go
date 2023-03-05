package student

import (
	"strconv"
	"strings"
)

func ConvertStudentStr(line string) (string, int, int, error) {
	splitedResult := strings.Split(line, " ")

	name := splitedResult[0]
	ageStr := splitedResult[1]
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return "", 0, 0, err
	}

	gradeStr := splitedResult[2]
	grade, err := strconv.Atoi(gradeStr)
	if err != nil {
		return "", 0, 0, err
	}

	return name, age, grade, nil
}
