package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName, LastName string
	Dob                 time.Time
	Email               string
}

func (person Person) PrintDetails() {
	fmt.Printf("First name %s", person.FirstName)
}

func main() {
	var person Person
	person.FirstName = "VijayKumar"
	person.LastName = "Subramani"
	person.Dob = time.Date(1991, 05, 26, 0, 0, 0, 0, time.UTC)
	person.Email = "vijay6.vit@gmail.com"

	person.PrintDetails()
}
