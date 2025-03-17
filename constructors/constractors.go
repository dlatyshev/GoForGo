package constructors

import "fmt"

type Rectangle struct {
	length int
	width  int
}

type Circle struct {
	radius int
}

// Constructor function for Circle
func NewCircle(radius int) *Circle {
	if radius <= 0 {
		fmt.Println("radius must be positive")
		return nil
	}
	return &Circle{
		radius: radius,
	}
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

	circle := NewCircle(5)
	fmt.Println("Radius:", circle.radius)
}
