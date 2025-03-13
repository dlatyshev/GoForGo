package strings

import (
	"fmt"
)

func ShowMapsExamples() {

	// string is byte slice
	word := "OhMyGosh"
	fmt.Println("word:", word) // OhMyGosh
	
	// What bytes are in the string?
	for i := 0; i < len(word); i++ {
		fmt.Printf("word[%d]: %d\n", i, word[i])
	}

	// What characters are in the string?
	for i := 0; i < len(word); i++ {
		fmt.Printf("%c ", word[i])
	}
	fmt.Println()

	// Rune is a character in Go
	// Rune is an alias for int32
    runeSlice := []rune(word)
	for i := 0; i < len(runeSlice); i++ {
        fmt.Printf("%c ", runeSlice[i])
	}

	// Create a string from a byte slice
	fmt.Println()
	myByteSlice := []byte{72, 101, 108, 108, 111}
	fmt.Println("myByteSlice:", myByteSlice) // [72 101 108 108 111]
	myStr := string(myByteSlice)
	fmt.Println("myStr:", myStr) // Hello
   

	// Create a string from a rune slice
	fmt.Println()
	myRuneSlice := []rune{72, 101, 108, 108, 111}
	myStringFromRune := string(myRuneSlice)
	fmt.Println("myStringFromRune:", myStringFromRune) // Hello
    
	// Create a string from a rune slice with literals
	fmt.Println()
	myRuneSliceWithLiterals := []rune{'W', 'o', 'r', 'l', 'd'}
	myStringFromRuneWithLiterals := string(myRuneSliceWithLiterals)
	fmt.Println("myStringFromRuneWithLiterals:", myStringFromRuneWithLiterals) // World


	// Strings comparison
	fmt.Println()
	fistName, lastName := "Dmytro", "Latyshev"
	anotherFirstName, anotherLastName := "Dmytro", "Petrov"

	fmt.Println("fistName == anotherFirstName:", fistName == anotherFirstName) // true
	fmt.Println("lastName == anotherLastName:", lastName == anotherLastName) // false


	// String concatenation
	fmt.Println()
	fullName := fistName + " " + lastName
	fmt.Println("fullName:", fullName) // Dmytro Latyshev

	// String builder
	fmt.Println()
	adminName := "Benjamin"
	adminSurname := "Franklin"

	adminFullName := fmt.Sprintf("%s %s", adminName, adminSurname)
	fmt.Println("adminFullName:", adminFullName) // Benjamin Franklin
}