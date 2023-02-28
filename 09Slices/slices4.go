package main

import "fmt"

func makeSlices(x int) ([]int, string) {

	data := []int{}
	for i := 0; i <= x; i++ {
		data = append(data, i)
	}

	return data, "This is an making of slices..."

}

func main() {
	fmt.Println("This file is used to learn the slices....")

	fruits := []string{"peaches", "banana", "watermelon", "muskmelon", "mangos"}

	fmt.Println(fruits[:4])

	for _, v := range fruits {
		fmt.Println(v)
	}

	d, s := makeSlices(10)
	fmt.Println(s, "-", d)

	array := [...]int{4: 1, 2, 3, 4} //[...] unfinite slices...
	fmt.Println(array)

}
