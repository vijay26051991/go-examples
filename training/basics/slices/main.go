package main

import (
	"fmt"
)

func main() {

	var firstnames = make([]string, 10)
	firstnames = append(firstnames, "vijay", "vimal")

	var lastnames = []string {"kumar", "andrew"}

	var strings []string = append(firstnames, lastnames...)
	result := make([]string, len(strings))

	copy(result, strings) // copy a slice

	sli := firstnames[:] // slicing the array

	fmt.Println(sli)
	fmt.Println("Original slice")

	//Original slice
	for index, val := range strings {
		if val != "" {
			fmt.Printf("\nIndex %d and My string %s \n", index, val)
		}
	}

	fmt.Println("Copied slice")
	//Copied slice
	for index, val := range result {
		if val != "" {
			fmt.Printf("\nIndex %d and My string %s \n", index, val)
		}
	}
}
