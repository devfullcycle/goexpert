package main

import "fmt"

func main() {
	ch := make(chan string, 20000)

	ch <- "Hello"
	ch <- "World"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
