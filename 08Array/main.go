package main

import "fmt"

func main() {
	fmt.Println("welconme to array in golang!")

	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "strawbeerys"
	fruitList[2] = "banana"
	fruitList[3] = "gavava"

	for i := 0; i < 4; i++ {
		fmt.Println(fruitList[i])
	}

	fmt.Println("Some vegitables are listed down!")

	vegitable := [4]string{"tomato", "potato", "mashrooms", "paneer"}

	fmt.Println(vegitable)

}
