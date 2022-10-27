package main

import (
	"fmt"

	"github.com/devfullcycle/fcutils-secret/pkg/events"
)

func main() {
   ed := events.NewEventDispatcher()
   fmt.Println(ed)
}