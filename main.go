package main

import (
	// "github.com/dlatyshev/GoForGo/guessinggame"
	// "github.com/dlatyshev/GoForGo/variables"
	"github.com/dlatyshev/GoForGo/apiclient"
	"fmt"
)

func main() {
	
	var objects []apiclient.Phone
	response := apiclient.Get("https://api.restful-api.dev/objects")
    err := apiclient.Decode(response, &objects)
	if err != nil {
		fmt.Println(err)
	}
	
	for _, object := range objects {
		fmt.Println(object.ToString())
	}
}
