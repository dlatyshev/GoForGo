package methods

import "fmt"

type Address struct {
	city   string
	street string
}

type Person struct {
	firstName string
	lastName  string
	Address
}

func (p *Person) GetFullName() string {
	return p.firstName + " " + p.lastName
}

func (a *Address) GetFullAddressInfo() string {
	return a.city + ", " + a.street
}

func ShowAnonymousMethodsExample() {
	p := Person{
		firstName: "John",
		lastName:  "Doe",
		Address: Address{
			city:   "New York",
			street: "5th Avenue",
		},
	}
	fmt.Println(p.GetFullName())        // This will call the method from the Person struct
	fmt.Println(p.GetFullAddressInfo()) // This will call the method from the Address struct
}
