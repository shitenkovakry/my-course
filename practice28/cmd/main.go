package main

import (
	"fmt"
	"io"
	"os"

	"curse/practice28/pkg/storage"
)

func main() {
	lines, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	students, err := storage.ConvertLinesToMapStudents(string(lines))
	if err != nil {
		panic(err)
	}

	for _, value := range students {
		fmt.Println(value.Name, value.Age, value.Grade)
	}
}
