package main

import (
	"fmt"
	"github.com/sujmat/goPlayground/packages/stringutil"
)

func main() {
	fmt.Println("Original:" + stringutil.MyName)
	fmt.Println("Reversed:" + stringutil.Reverse(stringutil.MyName))
}
