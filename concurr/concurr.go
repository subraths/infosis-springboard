package concurr

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func process2(ch chan int) {
	n := rand.IntN(3000)
	time.Sleep(time.Duration(n) * time.Millisecond)
	ch <- n
}

func Concurr() {
	ch := make(chan int)
	go process2(ch)

	fmt.Println("waiting for process")
	n := <-ch
	fmt.Printf("Process took %dms\n", n)
}
