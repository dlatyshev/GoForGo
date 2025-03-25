package concurrency

import "time"

func EvenGoroutine(number chan int) {
	for i := 0; i <= 20; i += 2 {
		number <- i
		time.Sleep(50 * time.Millisecond) // Simulate some work
	}
}

func OddGoroutine(number chan int) {
	for i := 1; i <= 20; i += 2 {
		number <- i
		time.Sleep(50 * time.Millisecond) // Simulate some work
	}
}

func ShowChannels2Example() {
	// Create a channel
	number := make(chan int)
	go EvenGoroutine(number)
	go OddGoroutine(number)

	// Wait for the goroutines to finish
	for i := 0; i < 20; i++ {
		num := <-number
		print(num, " ")
	}

	println("Finished printing even and odd numbers")
}
