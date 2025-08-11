package main

// next and last are URLs
type config struct {
	next string
	last string
}

type location struct {
	name string `json:"name"`
	url string `json:"url"`
}

type pageLocations struct {
	next *string `json:"next"`				// can be null, dereference value
	last *string `json:"previous"`			// can be null, dereference value
	locations []location `json:"results"`
}
