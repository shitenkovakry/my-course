package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	read, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Println("here we go")
	fmt.Println(string(read))
}
