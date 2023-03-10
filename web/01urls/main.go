package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentId=123nischayjain"

func main() {

	fmt.Println("Welcome to the class on the URLS in GOLang!")
	fmt.Println(myUrl)

	result, err := url.Parse(myUrl)

	if err != nil {
		panic(err)
	}
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.RawQuery)
	// fmt.Println(result.Port())
	// fmt.Println(result.Path)

	qparams := result.Query()
	for k, v := range qparams {
		fmt.Println(k, ":", v)
	}

	partsOfUrls := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutess",
		RawPath: "user=hitesh",
	}

	anotherURL := partsOfUrls.String()
	fmt.Println(anotherURL)

}
