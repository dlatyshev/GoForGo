package methods

import "fmt"

type Employee struct {
	name     string
	position string
	salary   int
	currency string
}

func (e Employee) DisplayInfo() {
	fmt.Println("Name:", e.name)
	fmt.Println("Position:", e.position)
	fmt.Println("Salary:", e.salary)
	fmt.Println("Currency:", e.currency)
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	length float64
	width  float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func ShowMethodsExamples() {
	employee := Employee{
		name:     "John Doe",
		position: "Software Engineer",
		salary:   60000,
		currency: "USD",
	}

	employee.DisplayInfo()

	circle := Circle{radius: 5}
	rectangle := Rectangle{length: 10, width: 5}
	fmt.Printf("Circle Area: %.2f\n", circle.Area())
	fmt.Printf("Rectangle Area: %.2f\n", rectangle.Area())
}
