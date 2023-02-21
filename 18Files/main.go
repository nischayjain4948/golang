package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("welcome to the tutorial on files in go language")
	file, error := os.Create("./mylocalFile.txt") // os module or package for creating the file

	if error != nil {
		panic(error)
	}

	length, err := io.WriteString(file, "Hello this is a txt file written with io.WriteString function") // io package for input output..

	if err != nil {
		panic(err)
	}
	fmt.Println("The length of the file is ", length)

	defer file.Close()

}
