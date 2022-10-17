package main

import (
	"fmt"
	"sync"
	"time"
)

// Thread 1
func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go count(ch, &wg) // background thread
	go reader(ch, &wg)
	wg.Wait()
}

func count(ch chan int, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		// calculo importante
		ch <- i
		time.Sleep(time.Second)
	}
	close(ch)
	wg.Done()
}

func reader(ch chan int, wg *sync.WaitGroup) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
		wg.Done()
	}
}
