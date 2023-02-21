package main

import (
	"fmt"
)

func add(a, b int) int {
	sum := a + b
	return sum
}

func greet(name string) {
	fmt.Printf("Namaste %v From GoLang!\n", name)
}

func proAdder(values ...int) (int, string) { // Three tripple dots are veradic function same as spread operator
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "output is this"
}

func main() {
	fmt.Println("Welcome to the topic, functions..")
	greet("nischay")
	result := add(5, 4)
	fmt.Println("The result is : ", result)
	proResult, response := proAdder(1, 2, 3, 4, 5)
	fmt.Println(response, "=", proResult)

	func() {
		fmt.Println("iife function, This is annonomous function")
	}()
}
