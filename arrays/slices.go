package arrays

import "fmt"

func ShowSliceExamples() {

	// Length and capacity of a slice
	wordsSlice := []string{"Hello", "World"}
	fmt.Println("slice:", wordsSlice) // [Hello World]
	fmt.Println("length:", len(wordsSlice)) // 2
	fmt.Println("capacity:", cap(wordsSlice)) // 2

	// Appending a new value to a slice. Capacity will be doubled if the current capacity is not enough.
	// In this case, the capacity will be 4.
	wordsSlice = append(wordsSlice, "!!!")
	fmt.Println("slice:", wordsSlice) // [Hello World !!!]
	fmt.Println("length:", len(wordsSlice)) // 3
	fmt.Println("capacity:", cap(wordsSlice)) // 4


	numArr := []int{1, 2, 3, 4, 5}
	numSlice := numArr[:]
	numSlice[0] = 10 // Changing the first element of the slice will also change the first element of the array
	numSlice = append(numSlice, 6) // New array will be created with a new capacity
	fmt.Println(numArr) // [1 2 3 4 5] // Original array remains unchanged
	fmt.Println(numSlice) // [1 2 3 4 5 6] // Slice with a new capacity


	// Slices with make function
	// The make function is used to create slices with a specific length and capacity.
	// The first argument is the type of the slice, the second argument is the length, and the third argument is the capacity.
	// If the capacity is not specified, it will be equal to the length.

	evenNumbers := make([]int, 5, 10) // Slice of length 5 and capacity 10
	fmt.Println("slice:", evenNumbers) // [0 0 0 0 0]


	// Merging slices
	// The append function can be used to merge two slices.

	firstSlice := []int{1, 2, 3}
	secondSlice := []int{4, 5, 6}
	mergedSlice := append(firstSlice, secondSlice...) // The ... operator is used to unpack the second slice
	fmt.Println("merged slice:", mergedSlice) // [1 2 3 4 5 6]

	// Multidimensional slices
	// Slices can be multidimensional, just like arrays.

	multiSlice := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 9},
	}

	fmt.Println("multidimensional slice:", multiSlice) // [[1 2 3] [4 5 6] [7 9]]
}