package main

import (
	"fmt"
	"os"
	"time"
)

func FormatToFile(text string) string {
	formated := fmt.Sprintf("%s %s\n", time.Now().Format("2006-01-02 15:04:05"), text)

	return formated
}

func main() {
	file, err := os.Create("practice12/task1/fileTask1.txt")
	if err != nil {
		fmt.Println("файл не создан")
	}
	defer file.Close()

	for {
		var text string
		fmt.Scan(&text)

		if text == "exit" {
			break
		}

		x := FormatToFile(text)

		file.WriteString(x)
	}
}
