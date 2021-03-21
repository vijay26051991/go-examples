package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wait sync.WaitGroup

func main()  {
	wait.Add(2)

	go add()
	go mul()

	wait.Done()
}

func add()  {
	defer wait.Done()
	for i :=0; i < 20; i++ {
		rand := rand.Int63n(1000)
		time.Sleep(time.Duration(rand) * time.Millisecond)

		for j :=1 ; j <20 ; j++ {
			fmt.Printf("\n\nAdd function called %d, %d, %d\n\n" , i , j , (i+j))
		}
	}
}

func mul()  {
	defer wait.Done()
	for i :=0; i < 20; i++ {
		rand := rand.Int63n(1000)
		time.Sleep(time.Duration(rand) * time.Millisecond)

		for j :=1 ; j <20 ; j++ {
			fmt.Printf("\n\nAdd function called %d, %d, %d\n\n" , i , j , (i*j))
		}
	}
}
