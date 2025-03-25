package concurrency

// We can create only one-way channel
// One-way channels are useful when we want to restrict the direction of data flow
// to ensure that a channel is only used for sending or receiving data.

func sender(ch chan<- int) {
	ch <- 42
}

func ShowDirectChannelExample() {
	ch := make(chan<- int)
	go sender(ch)
}
