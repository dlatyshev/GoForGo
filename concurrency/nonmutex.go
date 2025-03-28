package concurrency

import (
	"fmt"
	"sync"
)

var x int = 0

func increment(wg *sync.WaitGroup) {
	x++
	wg.Done()
}

func ShowNonMutexExample() {
	var wg sync.WaitGroup
	for range 1000 {
		wg.Add(1)
		go increment(&wg)
	}

	fmt.Println("Final value of x:", x)
}
