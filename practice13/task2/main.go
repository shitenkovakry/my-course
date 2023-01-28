package main

import (
	log "github.com/sirupsen/logrus"
)

func Swap(a *int, b *int) {
	*a, *b = *b, *a
}

func main() {
	x := 5
	y := 8

	Swap(&x, &y)

	log.Info(x, y)
}
