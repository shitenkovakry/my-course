package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func CountTo(N int) {
	milliseconds := time.Duration(rand.Intn(5000)) * time.Millisecond

	for i := 0; i < N; i++ {
		fmt.Println(i, "CountTo    :", N, milliseconds)

		time.Sleep(milliseconds)
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	waitGroup := &sync.WaitGroup{}

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()

		CountTo(10)
	}()

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()

		CountTo(20)
	}()

	fmt.Println("ia nachinaiu zhdat")
	waitGroup.Wait()
	fmt.Println("ia poshel")
}
