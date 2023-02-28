package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to the part of slices...")
	// slices are the most used Data structure

	fruitList := []string{"apple", "gavava", "peaches", "graps", "strawbeery"}

	// adding the values into slices.
	fruitList = append(fruitList, "banana", "mango")
	fmt.Println(fruitList[1:]) // [1:] - print slice from index 1 to limit.

	sort.Strings(fruitList)
	fmt.Println(fruitList)

	highScore := make([]int, 4)
	highScore[0] = 123
	highScore[1] = 12345
	highScore[2] = 1235
	highScore[3] = 1236
	sort.Ints(highScore)
	fmt.Println(highScore)

	fmt.Println(sort.IntsAreSorted(highScore))

	// remove the values from slices :)

	course := []string{"react", "angular", "node", "javascript", "swift", "ruby", "golang"}
	fmt.Println("hello there are some of the best course..", course)

	// Q1. the problem statement is that to remove the react from course slice;
	course = append(course[1:])
	fmt.Println("The react is removed from the course...", course)

	// Q2. the problem statement is that to remove the react from course slice;

	var index int = 2
	course = append(course[:index], course[index+1:]...)

	fmt.Println(course)

}
