package main

import (
	// "github.com/dlatyshev/GoForGo/guessinggame"
	// "github.com/dlatyshev/GoForGo/variables"
	"github.com/dlatyshev/GoForGo/apiclient"
	"fmt"
	"encoding/json"
)

func main() {
	// variables.ShowVariables()
	// guessinggame.Game()
	resp := apiclient.Get("https://api.restful-api.dev/objects")

	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)

	// var objects []apiclient.Phone
	// err := json.NewDecoder(resp.Body).Decode(&objects)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for _, object := range objects {
	// 	fmt.Println(object.Name)
	// }
	
	var objects []apiclient.Phone
	err := json.NewDecoder(resp.Body).Decode(&objects)
	if err != nil {
		fmt.Println(err)
	}
	
	for _, object := range objects {
		fmt.Println(object.ID)
		fmt.Println(object.Name)
		fmt.Println(object.Capacity)
		fmt.Println(object.Color)
	}
}
