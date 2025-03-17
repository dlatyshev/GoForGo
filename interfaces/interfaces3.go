package interfaces

import "fmt"

type Worker interface {
	Work()
}

type Employee struct {
	name string
	age  int
}

func (e Employee) Work() {
	fmt.Printf("Employee %s is working\n", e.name)
}

func Describe(w Worker) {
	fmt.Printf("Type: %T, Value: %v\n", w, w)
}

func ShowInterfaces3Example() {
	emp := Employee{name: "John", age: 30}
	emp.Work()
	Describe(emp)
}
