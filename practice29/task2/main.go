package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func NumberSquared(number int) int {
	squared := int(math.Pow(float64(number), 2.0))

	d := rand.Int63n(3000)
	fmt.Println("wait milliseconds = ", d)
	time.Sleep(time.Millisecond * time.Duration(d))

	return squared
}

func Process(number int) int {
	fmt.Println("this takes a lot:", number)

	result := NumberSquared(number)

	return result
}

func main() {
	wg := &sync.WaitGroup{}
	number := 0
	waitForSignalShutdown := make(chan os.Signal, 1)
	shutDown := make(chan struct{})

	signal.Notify(waitForSignalShutdown, syscall.SIGTERM)

	go func() {
		fmt.Println("waiting for signal")
		<-waitForSignalShutdown

		fmt.Println("shuting down the server")
		shutDown <- struct{}{}
	}()

	for {
		for j := 0; j < 10; j++ { // pool of workers
			number++

			wg.Add(1)
			go func(counter, number int) {
				defer wg.Done()
				result := Process(number) //very heavy computing

				fmt.Println(counter, "--", result)
			}(j, number)
		}

		wg.Wait()

		select {
		case <-shutDown:
			fmt.Println("server is down")

			return
		default:
			continue
		}
	}
}
