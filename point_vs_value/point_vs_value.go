package pointvsvalue

import "fmt"

type Rectangle struct {
	length int
	width  int
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

func Perimeter(r Rectangle) int {
	return 2 * (r.length + r.width)
}

func ShowPointVsValueExample() {
	rectangle := Rectangle{length: 10, width: 5}
	fmt.Println("Perimeter using method:", rectangle.Perimeter())
	fmt.Println("Perimeter using function:", Perimeter(rectangle))

	rectangleAsPointer := &rectangle
	fmt.Println("Perimeter using method with pointer:", rectangleAsPointer.Perimeter())

	fmt.Println("Perimeter using function with value from pointer", Perimeter(*rectangleAsPointer))
}
