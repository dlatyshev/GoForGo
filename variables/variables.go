package variables

import (
	"fmt"
	"unicode/utf8"
	"unsafe"
)


func ShowVariables() {
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


	fmt.Println(unsafe.Sizeof(a)) // 8 (bytes)

	fmt.Println(unsafe.Sizeof(width)) // 8 (bytes)

	var oneByteValue byte = 1
	var oneByteValue2 int8 = 1
	fmt.Println(unsafe.Sizeof(oneByteValue)) // 1 (bytes)
	fmt.Println(unsafe.Sizeof(oneByteValue2)) // 1 (bytes)


	var sampleRune rune = 'a'
	var anotherRule rune = 1234
	fmt.Printf("%c %c\n", sampleRune, anotherRule)

	fmt.Println(utf8.RuneCountInString("Hello, World!")) // 13
	fmt.Println("abc" > "ab") // true

	fmt.Printf("Is 14 even? %t\n", isEvenNumber(14)) // Is 14 even? true

	color := "blue"
	colorCode := getColorCode(color)
	fmt.Printf("Color code for %s is %v\n", color, colorCode) // Color code for blue is [0 0 255]
}

func isEvenNumber(number int) bool {
	return number % 2 == 0
}


func getColorCode(color string) []int {
	if color == "red" {
		return []int{255, 0, 0}
	} else if color == "green" {
		return []int{0, 255, 0}
	} else if color == "blue" {
		return []int{0, 0, 255}
	} else {
		return []int{0, 0, 0}
	}
}