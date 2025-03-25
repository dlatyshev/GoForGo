package concurrency

import (
	"fmt"
	"time"
)

func newGoroutine() {
	fmt.Println("Hello from a goroutine")
}

func ExampleGoroutine() {
	go newGoroutine()
	time.Sleep(1 * time.Second) // Wait for the goroutine to finish
	fmt.Println("Hello from the main function")
}
