package main

import (
	"flag"
	"fmt"
)

func CheckPartOfLine(line string, part string) bool {
	checkedPart := []byte{}

	for indexOfLine := 0; indexOfLine < len(line); indexOfLine++ {
		for indexOfPart := 0; indexOfPart < len(part); indexOfPart++ {
			p := byte(part[indexOfPart])
			l := byte(line[indexOfLine+indexOfPart])
			if p == l {
				checkedPart = append(checkedPart, p)

				if string(checkedPart) == part {
					return true
				}
			} else {
				checkedPart = []byte{}

				break
			}
		}
	}

	return false
}

func main() {
	var str string
	var substr string

	flag.StringVar(&str, "str", "", "this is the line")
	flag.StringVar(&substr, "substr", "", "this is the substring")

	flag.Parse()

	line := str
	substring := substr

	result := CheckPartOfLine(line, substring)

	fmt.Println(result)
}
