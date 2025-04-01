package concurr

import (
	"fmt"
	"math/rand/v2"
)

func producer2(ch chan<- int) {
	for {
		// sleep for random time
		sleep()

		// generate a random number
		n := rand.IntN(100)

		// send number
		fmt.Printf("-> %d\n", n)
		ch <- n
	}
}

func consumer2(ch <-chan int, name string) {
	for n := range ch {
		fmt.Printf("Consumer %s <- %d\n", name, n)
	}
}

func fanOut(chA <-chan int, chB, chC chan<- int) {
	for n := range chA {
		if n < 50 {
			chB <- n
		} else {
			chC <- n
		}
	}
}

func FanOut() {
	chA := make(chan int)
	chB := make(chan int)
	chC := make(chan int)

	go producer2(chA)
	go consumer2(chB, "B")
	go consumer2(chC, "C")

	fanOut(chA, chB, chC)
}
