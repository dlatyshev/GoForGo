package concurrency

import (
	"fmt"
	"sync"
)

var y int = 0

func incrementMutexed(wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	y++
	mu.Unlock()
	wg.Done()
}

func ShowMutexExample() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incrementMutexed(&wg, &mu)
	}

	wg.Wait()
	fmt.Println("Final value of y:", y)
}
