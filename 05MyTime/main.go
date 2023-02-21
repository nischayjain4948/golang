package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the time package!")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) // To check the current valid time we use the Format function

}
