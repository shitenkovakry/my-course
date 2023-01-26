package main

import "fmt"

func main() {
	a := int(10)

	var u1a *int
	var u2a *int

	u1a = &a
	u2a = &a
	u3a := u1a

	*u3a = 30

	fmt.Println(*u1a, *u2a, *u3a)
}
