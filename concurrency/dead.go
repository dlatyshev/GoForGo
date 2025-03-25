package concurrency

func deadlock() {
	ch := make(chan int)
	ch <- 1
}

func ShowDeadlockExample() {
	deadlock()
}
