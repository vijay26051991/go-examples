package main

import (
	"fmt"
	"reflect"
)

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println("Index", i)
	}

	count := 0
	//Commenting the infinite loop as need of executing following piece of code.
	/*for {
		count += 1
		fmt.Println("Index" , count)
	}*/

	count = 1
	for i := 0; i <= count; i++ {
		count *= i
		println("Count", count)
	}

	// find odd numbers
	num := 10
	oddnums := make([]int, 5)
	for i := 0; i <= num; i++ {
		if i%2 != 0 {
			oddnums = append(oddnums, i)
		}
	}
	fmt.Println(reflect.ValueOf(oddnums).Kind())
	fmt.Println(oddnums)

	strarray := []string{"Vijay", "Vimal", "Vishva", "Vignesh"}
	for index, value := range strarray {
		fmt.Println(index, " ", value)
	}
}
