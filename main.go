package main

import (
	// "github.com/dlatyshev/GoForGo/guessinggame"
	// "github.com/dlatyshev/GoForGo/variables"
	// "github.com/dlatyshev/GoForGo/apiclient"
	"fmt"
	"github.com/dlatyshev/GoForGo/interfaces"
)

func main() {
	rectangle := interfaces.Rectangle{Width: 10, Height: 5}
	circle := interfaces.Circle{Radius: 3}

	figures := []interfaces.Shape{rectangle, circle}
	
	for _, figure := range figures {
		fmt.Printf("Area of %T: %.2f\n", figure, figure.Area())
	}
}
