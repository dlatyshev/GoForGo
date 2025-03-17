package interfaces

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type BigWord interface {
	IsBig() bool
	ToUpper() string
}

type MySuperString string

func (s MySuperString) IsBig() bool {
	return utf8.RuneCountInString(string(s)) > 10
}

func (s MySuperString) ToUpper() string {
	return strings.ToUpper(string(s))
}

func ShowInterfaces2Example() {
	var s BigWord = MySuperString("Hello, World!")

	if s.IsBig() {
		fmt.Println("The string is big")
	} else {
		fmt.Println("The string is small")
	}
	fmt.Println(s.ToUpper())
}
