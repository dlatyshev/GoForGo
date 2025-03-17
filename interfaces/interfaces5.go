package interfaces

import "fmt"

type Flier interface {
	Fly() string
}

type Bird struct {
	name string
}

func (b Bird) Fly() string {
	return b.name + " is flying"
}

type Parrot struct {
	name string
}

func (p *Parrot) Fly() string {
	return p.name + " is flying"
}

func TypeFinder(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	case Flier:
		fmt.Println("Flyable")
	default:
		fmt.Println("unknown")
	}
}

func ShowInterfaces5Example() {
	i := 5
	s := "hello"
	TypeFinder(i)
	TypeFinder(s)
	TypeFinder(3.14)
	TypeFinder(Bird{"Sparrow"})
	TypeFinder(Parrot{"Parrot"})  // This will not work because Parrot is a pointer receiver
	TypeFinder(&Parrot{"Parrot"}) // This will work because we are passing a pointer to Parrot
}
