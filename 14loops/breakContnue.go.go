package main

import "fmt"

func main() {
	fmt.Println("This is an breakContinue statement..")

	roughValue := 1

	for roughValue < 10 {

		if roughValue == 2 {
			goto lco
		}

		if roughValue == 5 {
			roughValue++
			continue

		}

		fmt.Println("Value is : ", roughValue)
		roughValue++

	}

lco:
	fmt.Println("go to statement is working with lable")
}
