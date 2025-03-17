package interfaces

import "fmt"

// Interface is a type that defines a contract for other types to implement.
// It specifies a set of methods that must be implemented by any type that
// satisfies the interface. Interfaces are a powerful feature of Go that
// allow for polymorphism and code reuse. They enable you to write
// functions that can accept different types as arguments, as long as those
// types implement the required methods.

// In Go, an interface is defined using the `interface` keyword, followed by
// a list of method signatures. A type satisfies an interface if it implements
// all the methods defined in the interface. There is no explicit declaration
// of intent to implement an interface; if a type has the required methods,
// it satisfies the interface automatically. This is known as "structural
// typing".
//

type FigureBuilder interface {
	Area() float64
	Perimeter() float64
}

type Square struct {
	length float64
}

func (s Square) Area() float64 {
	return s.length * s.length
}

func (s Square) Perimeter() float64 {
	return 4 * s.length
}

type Triangle struct {
	base   float64
	height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

func (t Triangle) Perimeter() float64 {
	return t.base + 2*t.height // Assuming an isosceles triangle for simplicity
}

func ShowInterfacesExample() {
	square1 := Square{length: 5}
	square2 := Square{length: 6}
	square3 := Square{length: 7}
	triangle1 := Triangle{base: 3, height: 4}
	triangle2 := Triangle{base: 5, height: 6}
	triangle3 := Triangle{base: 7, height: 8}

	figures := []FigureBuilder{square1, square2, square3, triangle1, triangle2, triangle3}
	totalArea := 0.0
	totalPerimeter := 0.0
	for _, figure := range figures {
		totalArea += figure.Area()
		totalPerimeter += figure.Perimeter()
	}
	fmt.Printf("Total Area: %.2f\n", totalArea)
	fmt.Printf("Total Perimeter: %.2f\n", totalPerimeter)
}
