package deferred

import "fmt"

func ShowDeferredExample() {
	fmt.Println("Connection established")
	defer fmt.Println("Connection closed")
	fmt.Println("Performing operations")
	fmt.Println("Operation 1 completed")
	fmt.Println("Operation 2 completed")
	fmt.Println("Operation 3 completed")
	fmt.Println("All operations completed")
	fmt.Println("Exiting")
}
