package buf

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


func ShowBufExamples() {
	var name string
	fmt.Print("Enter your name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // Read the input from the user and store it in buffer
	name = scanner.Text() // Get the text from the buffer
	fmt.Println("Hello, " + name + "!")

	allEnteredLines := make([]string, 0, 10)
    
    fmt.Println("Enter lines of text (empty line to quit):")
	for scanner.Scan() {
		line := scanner.Text()
		allEnteredLines = append(allEnteredLines, line) // Append the entered line to the slice
		if line == "" {
			break // Exit the loop if an empty line is entered
		}
		fmt.Println("You entered:", line)
	}

	fmt.Println("All entered lines:", allEnteredLines) // Print all entered lines


	// Convert string to integer
	numStr := "12345"
	num, err := strconv.Atoi(numStr) // Convert string to integer (Atoi - Any to integer 32)

	if err != nil {
		fmt.Println("Error converting string to integer:", err)
	} else {
		fmt.Println("Converted number:", num) // Print the converted number
	}

	num64int, err := strconv.ParseInt(numStr, 10, 64) // Convert string to int64 (ParseInt - Parse Integer)
	if err != nil {
		fmt.Println("Error converting string to int64:", err)
	} else {
		fmt.Println("Converted number to int64:", num64int) // Print the converted number
	}


}