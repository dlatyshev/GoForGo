package apiclient


import (
	"net/http"
	"log"
)


func Get(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

func Post(url string) *http.Response {
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

