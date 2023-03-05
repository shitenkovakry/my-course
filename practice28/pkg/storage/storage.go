package storage

import (
	"strings"

	"curse/practice28/pkg/student"
)

type Student struct {
	Name  string
	Age   int
	Grade int
}

func ConvertLinesToMapStudents(line string) (map[string]*Student, error) {
	splitedLines := strings.Split(line, "\n")

	result := map[string]*Student{}

	for _, line := range splitedLines {
		if line == "" {
			continue
		}

		student, err := convertStudent(line)
		if err != nil {
			return nil, err
		}

		result[student.Name] = student
	}

	return result, nil
}

func convertStudent(line string) (*Student, error) {
	name, age, grade, err := student.ConvertStudentStr(line)
	if err != nil {
		return nil, err
	}

	student := &Student{
		Name:  name,
		Age:   age,
		Grade: grade,
	}

	return student, nil
}
