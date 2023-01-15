package main

import (
	"fmt"
	"time"
)

func main() {
	kryDR, err := time.Parse(time.RFC3339, "2002-10-03T05:00:00Z")

	fmt.Println(kryDR, kryDR.Unix(), err)

	ondDR, err := time.Parse(time.RFC3339, "1984-04-28T13:40:00Z")

	fmt.Println(ondDR, ondDR.Unix(), err)

}
