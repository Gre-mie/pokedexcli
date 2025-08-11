package main

import (
	f "fmt"
	"encoding/json"
	"io"
	"net/http"
	str "strings"
)

// lower cases input and splits it into an array of strings
func cleanInput(input string) []string {
	words := str.Split(str.ToLower(input), " ")
	return words
}

// retrieve json data 
func getJSON(url string, conf *config) (data string, err error) {
	result, err := http.Get(url)
	if err != nil {
		return "", f.Errorf("Couldn't GET from url - %v", err)
	}
	if result.StatusCode < 200 && result.StatusCode >= 300 {
		return "", f.Errorf("HTTP status code: %v - %v", result.StatusCode, err)
	}
	body, err := io.ReadAll(result.Body)
	result.Body.Close()
	if err != nil {
		return "", f.Errorf("io couldn't read result.Body - %v", err)
	}
	err = json.Unmarshal(body, conf)
	if err != nil {
		return "", f.Errorf("Couldn't convert json - %v", err)
	}



	f.Println("--- yay its working - http.Get(url) all good :D")
	f.Println(conf)

	return "", nil // temp
}