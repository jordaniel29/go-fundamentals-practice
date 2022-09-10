package main

import (
	"fmt"
	"strconv"
)

func main() {
	numSlice := []int{}

	for i := 0; i <= 10; i++ {
		numSlice = append(numSlice, i)
	}

	for _, num := range numSlice {
		if num%2 == 0 {
			fmt.Println(strconv.Itoa(num) + " is " + "even")
		} else {
			fmt.Println(strconv.Itoa(num) + " is " + "odd")
		}
	}
}