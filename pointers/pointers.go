package pointers

import (
	"math/rand"
	"fmt"
)

// Pointers are variables that store the memory address of another variable
func ShowPointersExamples() {
	var myName string = "Dima"
	var myNamePointer *string = &myName // Pointer to myName variable

	fmt.Println("Type of pointer: ", fmt.Sprintf("%T", myNamePointer)) // Type of pointer: *string
	fmt.Println("Pointer value:", myNamePointer) // Pointer value: 0xc00000a0b8
	fmt.Println("Pointer value dereferenced:", *myNamePointer) // Pointer value dereferenced: Dima

	// Pointer to zero value
	var myZeroVar int
	var myZeroVarPointer *int = &myZeroVar // Pointer to myZeroVar variable
	fmt.Println("Pointer to zero *value:", *myZeroVarPointer) // Pointer to zero value: 0

	// Shortcut for zero value pointer
	zeroPoint := new(int) // Pointer to zero value
	fmt.Println("Pointer to zero value:", *zeroPoint) // Pointer to zero value: 0

    // Change value of pointer
	*zeroPoint = 100 // Change value of pointer
	fmt.Println("Pointer to zero value:", *zeroPoint) // Pointer to zero value: 100

	// Modify variable using pointers
	balance := 1000
	visaPointer := &balance // Pointer to balance variable
	masterCardPointer := &balance // Pointer to balance variable

	*visaPointer += 500 // Change value via visaPointer
	*masterCardPointer -= 200 // Change value via masterCardPointer
	fmt.Println("Balance:", balance) // Balance: 1300

	myAge := 30
	modifyValue(&myAge) // Pass pointer to modifyValue function
	fmt.Println("My age:", myAge) // My age: 40

	lenght := returnPointer()
	width := returnPointer()
	fmt.Println("Pointer to width value:", *width) // Pointer to width value: 42
	fmt.Println("Pointer to length value:", *lenght) // Pointer to lenght value: 42

}

func modifyValue(value *int) {
	*value += 10 // Change value via pointer
}

func returnPointer() *int {
	var myNumber int = rand.Intn(100) // Random number between 0 and 100
	return &myNumber // Return pointer to myNumber variable
}