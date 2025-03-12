package maps

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

    
}