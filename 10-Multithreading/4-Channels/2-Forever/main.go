package main

import "fmt"

// Thread 1
func main() {
	forever := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Hello")
		}
		forever <- true
	}()

	<-forever
}
