package maps

import "fmt"

func ShowMapsExamples() {
	squares := make(map[int]int) // Create a map with int keys and int values
	squares[1] = 1
	squares[2] = 4
	squares[3] = 9
	squares[4] = 16

	fmt.Println(squares) // Print the map

	// map literal
	phoneBook := map[string]string{
		"John":   "123-456-7890",
		"Jane":   "987-654-3210",
		"Michael": "555-555-5555",
	}

	fmt.Println(phoneBook) // map[Jane:987-654-3210 John:123-456-7890 Michael:555-555-5555]

	// What can be used as a key in a map?
	// In Gomap, keys can be of any type that is comparable.
	// This includes built-in types like int, string, and float64, as well as structs and arrays.
	// However, slices, maps, and functions cannot be used as keys because they are not comparable.
	// For example, the following code will not compile:
	// myMap := make(map[[]int]string) // This will not compile because []int is not comparable

	// Zero value of a map is nil
	// A nil map behaves like an empty map when reading but causes a panic when writing.

	var myMap map[string]int // Declare a map variable without initializing it
	fmt.Println(myMap) // Print the nil map (output: map[])
	// myMap["key"] = 1 // This will cause a panic: assignment to entry in nil map


	// Get elements from a map
	fmt.Println(phoneBook["John"]) // Print the value associated with the key "John"

	// Get not-existing element from a map
	fmt.Println(phoneBook["NotExistingKey"]) // Print the value associated with the key "NotExistingKey" (output: empty string)


	// Check if a key exists in a map
	if value, exists := phoneBook["Billy"]; exists {
		fmt.Println("Billy's phone number:", value) // Print the value associated with the key "Billy"
	} else {
		fmt.Println("Billy not found in phone book")
	}


	// Iterate over a map
	for name, number := range phoneBook {
		fmt.Printf("%s: %s\n", name, number) // Print each key-value pair in the map
	}

	// Delete an element from a map
	delete(phoneBook, "John") // Delete the key "John" from the map
	fmt.Println(phoneBook) // map[Jane:987-654-3210 Michael:555-555-5555]
}
