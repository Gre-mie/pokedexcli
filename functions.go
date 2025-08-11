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

func setConfig(locations *pageLocations, conf *config) {
	f.Printf("\tconfig:\nnext: %v\nlast: %v\n\tlocatuions:\nnext: %v\nlast: %v\nlocations:\n", conf.next, conf.last, locations.next, locations.last, locations.locations)
}

// retrieve json data 
func getJSON(url string) (data *pageLocations, err error) {
	data = &pageLocations{}

	result, err := http.Get(url)
	if err != nil {
		return data, f.Errorf("Couldn't GET from url - %v", err)
	}
	if result.StatusCode < 200 || result.StatusCode >= 300 {
		return data, f.Errorf("HTTP status code: %v", result.StatusCode)
	}
	body, err := io.ReadAll(result.Body)
	defer result.Body.Close()
	if err != nil {
		return data, f.Errorf("io couldn't read result.Body - %v", err)
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, f.Errorf("Couldn't convert json - %v", err)
	}
	return data, nil
}