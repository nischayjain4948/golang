package main

import "fmt"

func main() {
	hits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight"}
	// fmt.Println(hits[:3])
	// fmt.Println(hits[3:])

	for v, i := range hits {
		fmt.Println(i, "-", v)
	}

}
