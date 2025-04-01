package workerpool

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func echoWorker(in, out chan int) {
	for {
		n := <-in
		time.Sleep(time.Duration(rand.IntN(3000)) * time.Millisecond)
		out <- n
	}
}

func produce(ch chan int) {
	i := 0
	for {
		fmt.Printf("-> send job: %d\n", i)
		ch <- i
		i++
	}
}

func Main() {
	in := make(chan int)
	out := make(chan int)

	for i := 0; i < 1000; i++ {
		go echoWorker(in, out)
	}
	go produce(in)

	for n := range out {
		fmt.Printf("<- Recev job %d\n", n)
	}
}
