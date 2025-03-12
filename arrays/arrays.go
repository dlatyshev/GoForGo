package arrays

import "fmt"

func ShowArrayExamples() {
    var arr [5]int
	fmt.Println(arr) // [0 0 0 0 0]

	for i := range len(arr) {
		arr[i] = i + 1
	}
	fmt.Println(arr) // [1 2 3 4 5]

	arr2 := [5]int{10, 20, 30, 40, 50}
	fmt.Println(arr2) // [10 20 30 40 50]

	arr3 := [5]int{100, 200, 300}
	fmt.Println(arr3) // [100 200 300 0 0]
    
	// The [...] syntax tells the compiler to infer the size of the array based on the number of elements in the initializer.
	arr4 := [...]int{10, 20, 30, 40}
	fmt.Println(arr4) // [10 20 30 40]

	// The [] syntax without the ... indicates a slice.
	// Slices are dynamic, meaning their size can change during runtime.
    arr5 := []int{1, 2, 3, 4, 5}
	fmt.Println(arr5) // [1 2 3 4 5]
    
	// Appending a new value to a slice
    arr5 = append(arr5, 6)
	fmt.Println(arr5) // [1 2 3 4 5 6]

	// Arrays copying
	names := [...]string{"Dima", "Sasha", "Vasya"}
    
	namesCopy := names
	names[0] = "Oleg"

	fmt.Println(names) // [Oleg Sasha Vasya]
	fmt.Println(namesCopy) // [Dima Sasha Vasya]

	// Arrays and its size
	var aArr [5]int
	var bArr [6]int

	aArr[0] = 100
	//bArr = aArr Cannot use aArr (type [5]int) as type [6]int in assignment
    fmt.Print(bArr)

	// Array iteration
	floatArr := [5]float64{1.205, 4.444, 5.232, 6.2321, 0}
	
	// range
	for index, value := range floatArr {
		fmt.Printf("%d element of array is %.2f\n", index, value)
	}

	// indexing
    for i := 0; i < len(floatArr); i++ {
		fmt.Printf("%d element of array is %.2f\n", i, floatArr[i])
	}

	// Multidimensional arrays
	multiArr := [2][2]int{
		{1, 2},
		{3, 4},
	}

	fmt.Println(multiArr) // [[1 2] [3 4]]


	// Slices
	// Slices are more flexible than arrays and can grow or shrink in size.
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice) // [1 2 3 4 5]

	slice = append(slice, 6) // Appending a new value to a slice
	fmt.Println(slice) // [1 2 3 4 5 6]

	slice = append(slice, 7, 8, 9) // Appending multiple values to a slice
	fmt.Println(slice) // [1 2 3 4 5 6 7 8 9]

	slice = append(slice[:2], slice[3:]...) // Removing an element '3' from a slice
	fmt.Println(slice) // [1 2 4 5 6 7 8 9]


	// Empty slices
	emptySlice := []int{} // An empty slice
	fmt.Println(emptySlice) // []

    // Slices can be created from arrays
	startArray := [5]int{1, 2, 3, 4, 5}
	startSlice := startArray[1:3] // Slicing the array from index 1 to 2 (exclusive)
	fmt.Println(startSlice) // [2 3]


	// Arrays modification using slices
	// When you modify a slice, the underlying array is modified as well.
	// This is because slices are references to the underlying array.
	initialArray := [5]int{1, 2, 3, 4, 5}
	mySlice := initialArray[1:3] // Slicing the array from index 1 to 2 (exclusive)
	mySlice[0] = 10 // Modifying the slice
	fmt.Println(initialArray) // [1 10 3 4 5]
}