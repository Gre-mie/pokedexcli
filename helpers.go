package main

func formatArr [T any] (arr []T, format func(T) T) []T {
	if len(arr) < 1 {return arr}
	newarr := make([]T, len(arr))
	for i, item := range arr {
		newarr[i] = format(item)
	}
	return newarr
} 
