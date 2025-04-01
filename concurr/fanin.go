package concurr

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func sleep() {
	time.Sleep(time.Duration(rand.IntN(3000)) * time.Millisecond)
}

func producer(ch chan<- int, name string) {
	for {
		// sleep for random time
		sleep()

		// generate a random number
		n := rand.IntN(100)

		// send number
		fmt.Printf("Channel %s -> %d\n", name, n)
		ch <- n
	}
}

func consumer(ch <-chan int) {
	for n := range ch {
		fmt.Printf("<- %d\n", n)
	}
}

func fanIn(chA, chB <-chan int, chC chan<- int) {
	var n int

	for {
		select {
		case n = <-chA:
			chC <- n
		case n = <-chB:
			chC <- n
		}
	}
}

func FanIn() {
	chA := make(chan int)
	chB := make(chan int)
	chC := make(chan int)

	go producer(chA, "A")
	go producer(chB, "B")
	go consumer(chC)

	fanIn(chA, chB, chC)
}
