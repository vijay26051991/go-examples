package main

import (
	"time"
	"fmt"
)

type Employee struct {
	FirstName, LastName string
	Dob time.Time
	Email string
}


func (person Employee) PrintDetails() {
	fmt.Printf("First name %s", person.FirstName)
}

func (developer Developer) PrintDetails() {
	fmt.Printf("First name %s", developer.Skills)
}


type Developer struct {
	Employee //type embedding for composition
	Skills []string

}

func main()  {
	var d Developer
	d = Developer{
		Employee: Employee{
			FirstName: "Vijay",
		},
		Skills: []string{"java"},
	}
	d.PrintDetails()

}
