package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	fmt.Println("This Go file is to read the file.....")

	fmt.Println("iotil is used for read the file....")
	dataByte, error := ioutil.ReadFile("mylocalFile.txt")

	if error != nil {
		fmt.Println(error)
		panic(error)
	}

	fmt.Println("The text Data  of the file is \n", string(dataByte))
}
