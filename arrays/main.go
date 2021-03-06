package main

import "fmt"

func main() {

	// array create
	var x [5]int
	x[0] = 10
	x[1] = 11
	fmt.Println(x)

	// array multiline
	y := [5]int{
		10,
		20,
		30,
		40,
		50,
	}
	fmt.Println(y)

	//initialize with ...
	z := [...]int{10, 20, 30, 40, 50}
	fmt.Println(z)

	//initialize values for specific elements
	a := [5]int{2: 10, 4: 40}
	fmt.Println(a)

	// Create a slice with make function
	c := make([]int, 5, 10)
	fmt.Println(c)

	// Create a slice with slice literal
	d := []int{10, 20, 30, 40, 50}
	fmt.Println(d)

	// Slice with initializing of specific elements
	e := []int{4: 3}
	fmt.Println(e)

	// Slice with no elements
	f := []int{}
	fmt.Println(f)

	//slicing examples
	mySlice := [] string{"a","b","c","d","e","f"}
	fmt.Println(mySlice)
	fmt.Println(mySlice[2:4])
	fmt.Println(mySlice[2])
	fmt.Println("myString"[2])

	// Slice capacity length demo
	capacityDemo()

	//Slice append
	sliceAppend()
}

func sliceAppend(){
	slice1 := []string {"sunday","monday","tuesday"}
	slice2 := []string {"wednesday","thursday","friday"}
	sliceNew := append(slice1,slice2...)
	fmt.Println(sliceNew)
}

func capacityDemo(){
	mySlice := make([]int,0,3)
	fmt.Println("--------------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("--------------")

	for i:=0; i<80; i++ {
		mySlice = append(mySlice,i)
		fmt.Println("Len:",len(mySlice)," Capacity:",cap(mySlice)," Value:",i)
	}
}