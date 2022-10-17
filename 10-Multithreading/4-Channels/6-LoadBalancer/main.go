package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data <-chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int)
	QtdWorkers := 100000

	// inicializa os workers baseado na qtd
	for i := 0; i < QtdWorkers; i++ {
		go worker(i, data)
	}

	// envia os dados para os workers
	for i := 0; i < 100000000; i++ {
		data <- i
	}
}
