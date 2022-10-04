package main

import "github.com/google/uuid"

func main() {
    println(uuid.New().String())
}