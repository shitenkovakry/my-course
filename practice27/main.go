package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func ConvertLinesToMapStudents(line string) (map[string]*Student, error) {
	splitedLines := strings.Split(line, "\n")

	result := map[string]*Student{}

	for _, line := range splitedLines {
		if line == "" {
			continue
		}

		student, err := ConvertStudent(line)
		if err != nil {
			return nil, err
		}

		result[student.Name] = student
	}

	return result, nil
}

func ConvertStudent(line string) (*Student, error) {
	name, age, grade, err := ConvertStudentStr(line)
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

type Student struct {
	Name  string
	Age   int
	Grade int
}

func main() {
	lines, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	students, err := ConvertLinesToMapStudents(string(lines))
	if err != nil {
		panic(err)
	}

	for _, value := range students {
		fmt.Println(value.Name, value.Age, value.Grade)
	}
}
