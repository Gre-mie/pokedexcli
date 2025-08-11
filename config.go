package main

// next and last are URLs
type config struct {
	next string
	last string
}

type location struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

type pageLocations struct {
	Count int `json:"count"`
	Next *string `json:"next"`				// can be null, dereference value
	Last *string `json:"previous"`			// can be null, dereference value
	Locations []location `json:"results"`
}
