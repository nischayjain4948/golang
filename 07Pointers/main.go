package main

import "fmt"

func main() {
	fmt.Println("Welcome to  the  class , we are using the learning concepts..")
	one := 50
	var ptrOne = &one // ptrOne is used to hold the address of &one variable.

	fmt.Println(ptrOne)  // printing the address which is store in the variable ptrOne.
	fmt.Println(*ptrOne) //  printing the value which is present in the *ptrOne address.

	*ptrOne = *ptrOne * 2
	fmt.Println("After multiplying the pointer from 2 the result is : ", *ptrOne)

}
