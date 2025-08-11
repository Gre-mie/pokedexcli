package main

import (
	f "fmt"
	"encoding/json"
	"io"
	"net/http"
	"os"
	str "strings"
)

// lower cases input and splits it into an array of strings
func cleanInput(input string) []string {
	words := str.Split(str.ToLower(input), " ")
	return words
}

func setConfig(locations *pageLocations, conf *config) {
	if locations.Next == nil {
		f.Println("locations.Next is Nil - exiting programme")
		os.Exit(1)
	}
	if locations.Last == nil {
		f.Println("locations.Last is Nil - exiting programme")
		os.Exit(1)
	}
	if len(locations.Locations) < 1 {
		f.Println("locations array is empty - exiting programme")
		os.Exit(1)
	}


	f.Printf("\tconfig:\nnext: %v\nlast: %v\n\tlocatuions:\nnext: %v\nlast: %v\nlocations:\n%v\n", conf.next, conf.last, locations.Next, locations.Last, locations.Locations)
}

// turn JSON data into a struct
func getJSON(url string) (pageLocations, error) {

	result, err := http.Get(url)
	if err != nil {
		return pageLocations{}, f.Errorf("Couldn't GET from url - %v", err)
	}
	if result.StatusCode < 200 || result.StatusCode >= 300 {
		return pageLocations{}, f.Errorf("HTTP status code: %v", result.StatusCode)
	}

	body, err := io.ReadAll(result.Body)
	defer result.Body.Close()
	if err != nil {
		return pageLocations{}, f.Errorf("io couldn't read result.Body - %v", err)
	}

	f.Println(string(body)) // debug print

	data := pageLocations{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return pageLocations{}, f.Errorf("Couldn't convert json - %v", err)
	}

	// debug test vvv
	if data.Next == nil {
		f.Println("data.next is Nil - exiting programme")
		os.Exit(1)
	}
	if data.Last == nil {
		f.Println("data.last is Nil - exiting programme")
		os.Exit(1)
	}
	if len(data.Locations) < 1 {
		f.Println("locations array is empty - exiting programme")
		os.Exit(1)
	}
	// debug test ^^^

	return data, nil
}