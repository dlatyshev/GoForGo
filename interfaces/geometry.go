package interfaces

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}
