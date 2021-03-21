package main

import (
	"errors"
	"fmt"
)

var employeemap map[string]Employee

func main() {
	employee := NewEmployee("vijay", "kumar", "vijay6.vit@gmail.com", 8122114207)
	employeemap = make(map[string]Employee)

	employeemap[employee.Firstname] = *employee
	AddEmployee(employee)

	employee.Firstname = "Vimal"
	employee.Lastname = "Durairaj"
	employee.Email = "abc@gmail.com"
	employee.Phone = 91212232323

	AddEmployee(employee)

	employee.Firstname = "Tom"
	employee.Lastname = "Cherry"
	employee.Email = "123@gmail.com"
	employee.Phone = 123456789

	AddEmployee(employee)

	employee.Firstname = "Andrew"
	employee.Lastname = "Cherry"
	employee.Email = "123@yahoo.com"
	employee.Phone = 987654321

	AddEmployee(employee)

	fmt.Println("\n ====[AddEmployee] function=== \n")
	fmt.Println(employeemap)

	fmt.Println("\n\n ====[GetByID] function=== \n\n")
	ConsumeReturnValues(GetByKey("Tom"))

	fmt.Println("\n\n ====[GetByID] function : Key doesn't exist=== \n\n")

	ConsumeReturnValues(GetByKey("Tom1"))

	ClearAll()

	PrintMap()

	fmt.Println("\n\n ====[AddEmployee] function \n\n")

	AddEmployee(employee)

	PrintMap()

	employee.Firstname = "Ed"
	employee.Lastname = "grahams"

	fmt.Println("\n\n ====[UpdateEmployee] function \n\n")

	UpdateEmployee("Andrew", employee)

	PrintMap()

	fmt.Println("\n\n ====[DeleteEmployee] function \n\n")

	DeleteEmployee("Andrew")

	PrintMap()
}

func AddEmployee(employee *Employee) {
	employeemap[employee.Firstname] = *employee
}

func UpdateEmployee(key string, employee *Employee) (Employee, error) {
	if !checkKeyExists(key) {
		return Employee{}, errors.New("Key doesn't exist")
	}
	employeemap[key] = *employee
	return employeemap[key], nil
}

func ClearAll() {
	for key := range employeemap {
		delete(employeemap, key)
	}
}

func GetByKey(key string) (Employee, error) {
	if !checkKeyExists(key) {
		return Employee{}, errors.New("Key doesn't exist")
	}
	return employeemap[key], nil
}

func DeleteEmployee(key string) error {
	if !checkKeyExists(key) {
		return nil
	}
	delete(employeemap, key)
	return nil
}

type Employee struct {
	Firstname string
	Lastname  string
	Email     string
	Phone     int
}

func NewEmployee(firstname string, lastname string, email string, phone int) *Employee {
	return &Employee{Firstname: firstname, Lastname: lastname, Email: email, Phone: phone}
}

func checkKeyExists(thekey string) bool {
	if _, exists := employeemap[thekey]; !exists {
		return false
	}
	return true
}

func ConsumeReturnValues(x interface{}, err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(x)
}

func PrintMap() {
	fmt.Println("\nCurrent data in the employee map:")
	fmt.Println(employeemap)
}
