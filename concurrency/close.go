package concurrency

import "fmt"

func generator(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}

func ShowChannelCloseExample() {
	ch := make(chan int)
	go generator(ch)

	for {
		val, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(val)
	}

	fmt.Println("Channel closed")
	fmt.Println("End of ShowChannelCloseExample")
}
