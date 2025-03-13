package functions


import (
	"fmt"
)

func ShowFunctionsExamples() {
    a := 5
	b := 10
	c := Add(a, b)
    fmt.Printf("Sum of %d and %d is %d\n", a, b, c) // Sum of 5 and 10 is 15

	area, perimeter := RectParameters(5, 10)
	fmt.Printf("Area: %d, Perimeter: %d\n", area, perimeter) // Area: 50, Perimeter: 30
    
	fmt.Println((GetFirstAndLastName())) // Dmytro Latyshev

	fmt.Println(Sum(1, 2, 3, 4, 5)) // 15

	// Anonymous function
	anonFunc := func(x, y int) int {
		return x * y
	}

	result := anonFunc(5, 10)
	fmt.Printf("Anonymous function result: %d\n", result) // Anonymous function result: 50
}

func Add(x, y int) int {
	return x + y
}

func RectParameters(length, width int) (int, int) {
	area := length * width
	perimeter := 2 * (length + width)
	return area, perimeter
}

// Named return values
func GetFirstAndLastName() (firstName, lastName string) {
	firstName = "Dmytro"
	lastName = "Latyshev"
	return
}

// Empty return
func GetEmptyReturn() {
   _ = 5 + 5
}

// Variadic function
func Sum(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

