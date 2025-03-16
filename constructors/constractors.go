package constructors

import "fmt"

type Rectangle struct {
	length int
	width  int
}

// Constructor function for Rectangle
// By convention, constructor functions in Go are named with a capital letter
// and return a pointer to the struct type they create.
func NewRectangle(length, width int) *Rectangle {
	return &Rectangle{
		length: length,
		width:  width,
	}
}

func ShowConstructorsExample() {
	rect := NewRectangle(10, 5)
	fmt.Println("Length:", rect.length)
}
