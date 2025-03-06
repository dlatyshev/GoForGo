package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	var name string
	var age int

	name = "Dmytro"
	age = 35

    fmt.Printf("Hello, %s! You are %d years old\n", name, age)

	var a, b int
	fmt.Println(a, b) // 0 0

	var c, d int = 1, 2
	fmt.Println(c, d) // 1 2


	width, height := 100, 200
	fmt.Println(width, height) // 100 200

}
