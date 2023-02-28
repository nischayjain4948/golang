package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in go lang.")

	days := []string{"Sunday", "Monday", "Tuesday", "Thrusday", "Friday", "Saturday"}
	fmt.Println(days)

	// normal for
	fmt.Println("Normal for loop")
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])

	}

	// comma ok loop
	fmt.Println("--------")
	fmt.Println("comma ok loop")
	for _, i := range days {
		fmt.Println(i)
	}

	//  type of forEach loop
	fmt.Println("--------")
	fmt.Println("for with foreach")
	for v, i := range days {
		fmt.Println(v, "-", i)
	}

}
