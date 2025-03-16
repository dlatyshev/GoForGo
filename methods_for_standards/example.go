package methodsforstandards

type MySuperInt int // MySuperInt is a custom type that wraps the int type

func (i *MySuperInt) isEven() bool { // Method to check if the number is even
	return *i%2 == 0
}

func ShowCustomTypeExample() {
	var i MySuperInt = 10 // Declare a variable of type MySuperInt
	if i.isEven() {       // Call the method to check if it's even
		println("The number is even")
	} else {
		println("The number is odd")
	}
	i++                              // Increment the value of i
	println("The number is now:", i) // Print the new value of i
}
