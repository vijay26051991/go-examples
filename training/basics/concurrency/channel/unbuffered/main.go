package main

import "fmt"

var mathresult chan int

func main() {
	mathresult = make(chan int)
	go SendToChannel(Add(3, 5))
	result := ReceiveFromChannel()
	fmt.Println(result)
}

func Add(a, b int) int {
	return a + b
}

func SendToChannel(result int) {
	mathresult <- result
}

func ReceiveFromChannel() int {
	result := <-mathresult
	return result
}
