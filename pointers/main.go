package main

import "fmt"

func modifyVar(b *int){
	*b = 20
	fmt.Printf("Address being modified in function = %v\n",b)
}

func main() {
	a := 10
	fmt.Printf("original val = %v, address a= %v\n",a,&a)
	modifyVar(&a)
	fmt.Printf("modified val = %v, address a= %v\n",a,&a)
}