package main

import (
	"fmt"
	"strings"
)

var stringchannel chan []string

func init() {
	stringchannel = make(chan []string, 3)
}
func main() {
	strings := make([]string, 3)

	strings = append(strings, "vijay", "vishva", "tom")

	go func() {
		fmt.Println(strings)
		stringchannel <- strings
		close(stringchannel)
	}()
	//time.Sleep(10 * time.Second)
	receive()

}
func receive() {
	//To upper case...
	for {
		result, ok := <-stringchannel

		if !ok {
			fmt.Println("No data left and channel is closed")
			break
		}

		fmt.Println(len(result))
		for _, str := range result {
			fmt.Println(strings.ToUpper(str))
		}
	}
}
