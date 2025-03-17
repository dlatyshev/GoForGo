package embedinterfaces

import "fmt"

type AreaCalculator interface {
	Area() float64
}

type PerimeterCalculator interface {
	Perimeter() float64
}

type FigureParamsCalculator interface {
	AreaCalculator
	PerimeterCalculator
}

type Rect struct {
	width, height float64
	color         string
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func ShowExamples() {
	var r FigureParamsCalculator
	r = Rect{width: 5, height: 10, color: "red"}
	fmt.Println("Area:", r.Area())
	fmt.Println("Perimeter:", r.Perimeter())
	fmt.Println("Color:", r.(Rect).color)
}
