package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	ReadContentFromFile()

}

func OpenFile(path string) *os.File {
	src, err := os.Open(path)
	if err != nil {
		logErrorMessage(err)
		panic(fmt.Sprintf("Panicking!!!"))
	}
	return src

}

func ReadContentFromFile() {

	defer func() {
		fmt.Println("I am in the defer function..")
		if r := recover(); r != nil {
			fmt.Println(r)
			ProcessFile("sample.txt")
		}
	}()

	ProcessFile("sample1.txt")
}

func ProcessFile(path string) {
	file := OpenFile(path)
	all, err := io.ReadAll(file)
	defer file.Close()

	logErrorMessage(err)
	fmt.Println("===File content===", string(all))
}

func logErrorMessage(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
