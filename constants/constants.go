package constants

import (
	"fmt"
)

const (
	// Constants are immutable and cannot be changed
	BASE_URL = "https://api.example.com"
	PORT     = "8080"
)

func ShowConstantsExamples() {
    const PI = 3.14
	const E = 2.71828
	fmt.Println("PI:", PI) // PI: 3.14
	fmt.Println("E:", E) // E: 2.71828
    fmt.Println("URL:", BASE_URL + ":" + PORT) // URL: https://api.example.com:8080

	// Const value shoukd be known at compile time
	//const a = math.Sqrt(16) // Cannot be used. Compile time error.
    

	// Const with types
    const ADMIN_EMAIL string = "admin@salesforce.com"
	fmt.Println("Admin email:", ADMIN_EMAIL) // Admin email: admin@salesforce.com

	// Const without types
	const NUMERIC = 100
	var numInt8 int8 = NUMERIC // Implicit conversion to int8
	var numInt16 int16 = NUMERIC // Implicit conversion to int16
	var numInt32 int32 = NUMERIC // Implicit conversion to int32
	var numInt64 int64 = NUMERIC // Implicit conversion to int64
	var numFloat32 float32 = NUMERIC // Implicit conversion to float32
	var numFloat64 float64 = NUMERIC // Implicit conversion to float64

	fmt.Println(numInt8, numInt16, numInt32, numInt64, numFloat32, numFloat64) // 100 100 100 100 100 100
}