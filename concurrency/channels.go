package concurrency

import (
	"fmt"
)

func BoolGoroutine(done chan bool) {
	fmt.Println("Hello from a Bool goroutine")
	done <- true
}

func IntGoroutine(done chan int) {
	fmt.Println("Hello from an Int goroutine")
	done <- 100
}

func ShowChannelsExample() {
	// Create a channel
	ch := make(chan int)
	fmt.Printf("Type of ch: %T\n", ch)

	// Start a goroutine
	boolCh := make(chan bool)
	go BoolGoroutine(boolCh)
	go IntGoroutine(ch)

	val := <-boolCh
	valInt := <-ch
	fmt.Println("Hello from the main function")
	fmt.Println("Value from bool goroutine: ", val)
	fmt.Println("Value from int goroutine: ", valInt)
}
