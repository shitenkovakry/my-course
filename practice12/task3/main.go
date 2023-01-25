package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("fileTask3.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	toReadInto := make([]byte, 100)

	m, err := file.Read(toReadInto)
	if err != nil {
		panic(err)
	}

	fmt.Printf("read m=%d bytes [%s]", m, string(toReadInto[0:m]))

	if err := os.Chmod("fileTask3.txt", 0400); err != nil {
		panic(err)
	}

	n, err := file.WriteString("cannot write this string in file")
	if err != nil {
		panic(err)
	}

	fmt.Printf("wrote n=%d bytes", n)
}
