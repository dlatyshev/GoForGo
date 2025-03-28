package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("Started Goroutine", i)
	time.Sleep(2 * time.Second)
	fmt.Println("Finished Goroutine", i)
	wg.Done()
}

func ShowWaitGroupExample() {
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}

	wg.Wait()
	fmt.Println("All Goroutines finished")
}
