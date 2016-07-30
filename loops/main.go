package main

import "fmt"

func main() {

	// Variant 1
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println(i, " - ", j)
		}
	}

	// Variant 2
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	//Variant 3
	k := 0
	for {
		if k > 10 {
			break
		}
		fmt.Println(k)
		k++
	}
}
