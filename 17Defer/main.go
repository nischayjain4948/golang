package main

import "fmt"

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func main() {

	defer fmt.Println("This is an defer") // defer keyword will works at the end part
	fmt.Println("Welcome to the class on defer")

	defer fmt.Println("One")
	defer fmt.Println("two")
	defer fmt.Println("three")

	myDefer()

}
