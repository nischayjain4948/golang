package main

import "fmt"

func main() {
	fmt.Println("This is used to cover the slices topic..")
	course := []string{"react", "angular", "node", "javascript", "swift", "ruby", "golang"}
	fmt.Println("hello there are some of the best course..", course)

	for _, i := range course { // range function for looping the slices...

		fmt.Println("The course are", i)
	}

	index := 3
	course = append(course[:index], course[index+1:]...)

	fmt.Println("We have successfully removed the javascript from course variable...", course)

}
