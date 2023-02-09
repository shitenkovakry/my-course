package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var path1 string
	var path2 string
	var resultPath string

	flag.StringVar(&path1, "p1", "", "the path for first file")
	flag.StringVar(&path2, "p2", "", "the path for second file")
	flag.StringVar(&resultPath, "r", "", "the path to write the result")

	flag.Parse()

	file1Name := path1
	file2Name := path2

	data1, err := os.ReadFile(file1Name)
	if err != nil {
		panic(err)
	}

	data2, err := os.ReadFile(file2Name)
	if err != nil {
		panic(err)
	}

	unionFiles := fmt.Sprintf("%s%s", string(data1), string(data2))

	if err := os.WriteFile(resultPath, []byte(unionFiles), os.ModePerm); err != nil {
		panic(err)
	}
}
