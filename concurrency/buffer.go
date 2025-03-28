package concurrency

import "fmt"

func ShowBufferExample() {
	ch := make(chan string, 2)
	ch <- "Hello"
	ch <- "World"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
