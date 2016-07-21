package main

import "fmt"

func main() {
	a := 10
	b := "matt"
	c := 3.12
	d := true

	fmt.Println(a, b, c, d)
	fmt.Printf("%T %T %T %T",a,b,c,d)
}
