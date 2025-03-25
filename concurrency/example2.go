package concurrency

import (
	"fmt"
	"time"
)

func printEvenNumbers() {
	for i := 0; i <= 20; i++ {
		time.Sleep(250 * time.Millisecond) // Simulate some work
		if i%2 == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println("Even numbers goroutine finished")
}

func printOddNumbers() {
	for i := 0; i <= 20; i++ {
		time.Sleep(450 * time.Millisecond) // Simulate some work
		if i%2 != 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println("Odd numbers goroutine finished")
}

func ExampleGoroutine2() {
	go printEvenNumbers()
	go printOddNumbers()

	// Wait for the goroutines to finish
	time.Sleep(10 * time.Second)
	fmt.Println("Finished printing even and odd numbers")
}
