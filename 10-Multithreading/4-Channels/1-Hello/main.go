package main

import "fmt"

// Thread 1
func main() {
	canal := make(chan string) // Vazio

	// Thread 2
	go func() {
		canal <- "Olá Mundo!" // Está cheio
	}()

	// Thread 1
	msg := <-canal // Canal esvazia
	fmt.Println(msg)
}
