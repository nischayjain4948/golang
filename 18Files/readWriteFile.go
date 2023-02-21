package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFIle(fileName, data string) (int, string) {

	file, error := os.Create(fileName)

	if error != nil {
		panic(error)

	}
	length, err := io.WriteString(file, data)
	if err != nil {
		panic(err)
	}
	return length, "The file is created"
}

func readFile(fileName string) (string, string) {
	fileContent, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	return "The file content", string(fileContent)

}

func main() {
	fmt.Println("This file is used to write the file and read the file respectively")
	writeFileResult, msg := readFIle("localfile.txt", "This is a local file of txt which")
	fmt.Printf(msg, "character length are file is %v\n", writeFileResult)

	msg, readFileReslt := readFile("localfile.txt")

	fmt.Println(msg, "_", readFileReslt)

}
