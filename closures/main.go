package main

import "fmt"

func getIncrementer() func() int {
	x := 0

	return func() int {
		x = x + 1
		return x
	}
}

func main() {
	increment := getIncrementer()

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}
