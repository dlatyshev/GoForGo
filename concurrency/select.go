package concurrency

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "Server 1"
}

func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "Server 2"
}

func server3(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Server 3"
}

func ShowSelectExample() {
	output1 := make(chan string)
	output2 := make(chan string)
	output3 := make(chan string)

	go server1(output1)
	go server2(output2)
	go server3(output3)

	notReceivedCounter := 0
	for notReceivedCounter < 10 {
		time.Sleep(1000 * time.Millisecond)
		select {
		case msg1 := <-output1:
			fmt.Println("Received from server 1:", msg1)
			notReceivedCounter = 0
		case msg2 := <-output2:
			fmt.Println("Received from server 2:", msg2)
			notReceivedCounter = 0
		case msg3 := <-output3:
			fmt.Println("Received from server 3:", msg3)
			notReceivedCounter = 0
		default:
			fmt.Println("No messages received yet.")
			notReceivedCounter++
		}
	}
}
