package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fileName := "practice12/task4/secondPart/fileTask4.txt"

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
