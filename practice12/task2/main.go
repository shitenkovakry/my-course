package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("fileTask1.txt")
	if err != nil {
		fmt.Println("ошибка открытия файла", err)
		return
	}
	defer file.Close()

	b := make([]byte, 11)

	for {
		n, err := file.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		fmt.Print(string(b[0:n]))
	}
}
