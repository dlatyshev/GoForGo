package pointers

import (
	"fmt"
)

func ShowArrayPointersExamples() {
    arr := [3]int{1, 2, 3}
	fmt.Println("Original array:", arr) // Original array: [1 2 3]
    
	mutation(&arr) // Pass pointer to array
	fmt.Println("Array after mutation:", arr) // Array after mutation: [100 200 300]

	mutationSlice(arr[:]) // Pass slice of array
	fmt.Println("Array after mutationSlice:", arr) // Array after mutationSlice: [200 400 600]

}

func mutation(arr *[3]int) {
	(*arr)[0] = 100
	(*arr)[1] = 200
	(*arr)[2] = 300
}

// User slices instead of arrays
func mutationSlice(arr []int) {
	arr[0] += 100
	arr[1] += 200
	arr[2] += 300
}