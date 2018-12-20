package main

import (
    "fmt"
    "github.com/fernanlukban/cronus/event"
)

func main() {
    fmt.Println("Hello, world!")

    event1 := event.Event{}

    fmt.Println(event1.Title())
}