package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func FormatToFile(text string) string {
	formated := fmt.Sprintf("%s %s\n", time.Now().Format("2006-01-02 15:04:05"), text)

	return formated
}

func main() {
	fileName := "practice12/task4/fileTask4.txt"
	fmt.Println(fileName)

	toWrite := make([]byte, 0)

	for {
		var text string
		fmt.Scan(&text)

		if text == "exit" {
			break
		}

		x := FormatToFile(text)

		toWrite = append(toWrite, []byte(x)...)
	}

	err := ioutil.WriteFile(fileName, toWrite, 0755)
	if err != nil {
		panic(err)
	}

}
