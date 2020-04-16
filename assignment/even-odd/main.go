package main

import "fmt"

func main() {
	numbers := newIntSlice(10)

	for _, value := range numbers {
		if value%2 == 0 {
			fmt.Println(value, "is even")
		} else {
			fmt.Println(value, "is odd")
		}
	}
}

func newIntSlice(length int) []int {
	intSlice := []int{}
	for i := 0; i <= length; i++ {
		intSlice = append(intSlice, i)
	}
	return intSlice
}
