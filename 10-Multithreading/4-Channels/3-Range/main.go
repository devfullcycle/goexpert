package main

import (
	"fmt"
	"time"
)

// Thread 1
func main() {
	ch := make(chan int)
	go count(ch) // background thread

	for x := range ch {
		fmt.Printf("Received %d\n", x)
	}
}

func count(ch chan int) {
	for i := 0; i < 10; i++ {
		// calculo importante
		ch <- i
		time.Sleep(time.Second)
	}
	close(ch)
}

func reader(ch chan int) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
	}
}
