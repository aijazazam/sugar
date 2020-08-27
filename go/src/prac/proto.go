package main

import (
	"fmt"
	lib "common"
)

func main() {

	tMovie := lib.Movie{ 20, "Dhoom", "Action", "Its Good Movie", "Sanjay Gadhvi", "aamir Khan", 2010, 130, 6.3, 34000, 141, 381}

	tBytes, err := tMovie.Marshal()
	if err != nil {
		panic(err)
	}

	tUnmarshal := lib.Movie{}
	tUnmarshal.Unmarshal(tBytes)

	fmt.Println(tBytes)
	fmt.Println(tUnmarshal)
}

