package concurrency

import (
	"fmt"
)

func squareGoroutine(num int, result chan int) {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit
		num /= 10
	}

	result <- sum
}

func cubesGproutine(num int, result chan int) {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit * digit
		num /= 10
	}

	result <- sum
}

func ShowChannels3Example() {
	// Create a channel
	squaresChanel := make(chan int)
	cubesChanel := make(chan int)

	// Start goroutines
	go squareGoroutine(1234, squaresChanel)
	go cubesGproutine(1234, cubesChanel)

	// Wait for the goroutines to finish
	squareResult := <-squaresChanel
	cubesResult := <-cubesChanel

	fmt.Println("Square result: ", squareResult)
	fmt.Println("Cubes result: ", cubesResult)
}
