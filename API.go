package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// Make a GET request
	resp, err := http.Get("https://binaryjazz.us/wp-json/genrenator/v1/story/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
