package methods

import "fmt"

type Book struct {
	title  string
	author string
	price  float64
	pages  int
}

func (b Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func ShowMethods2Example() {
	book := Book{
		title:  "The Go Programming Language",
		author: "Alan A. A. Donovan",
		price:  45.0,
		pages:  400,
	}

	fmt.Println("Before setting title and author:")
	fmt.Println("Title:", book.title)
	book.SetTitle("The Go Programming Language (2nd Edition)")
	fmt.Println("After setting title:")
	fmt.Println("Title:", book.title)

	fmt.Println("Before setting author:")
	fmt.Println("Author:", book.author)
	book.SetAuthor("Brian W. Kernighan")
	fmt.Println("After setting author:")
	fmt.Println("Author:", book.author)
}
