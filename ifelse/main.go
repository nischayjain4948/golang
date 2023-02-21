package main

import "fmt"

func main() {
	fmt.Println("If else in GO Lang")

	pocketMoney := 1700

	if pocketMoney < 500 {
		fmt.Println("You can only shop with your parents")

	} else if pocketMoney > 1000 && pocketMoney < 10000 {
		fmt.Println("You can go with your cousins")
	} else {
		fmt.Println("You can shop with your girlfriend")
	}

	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is grethar than 10")
	}

}
