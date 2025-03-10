package apiclient

import (
	"net/http"
	"log"
	"encoding/json"
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

func Decode(resp *http.Response, v any) error {
	err := json.NewDecoder(resp.Body).Decode(&v)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
