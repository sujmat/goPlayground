package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	hash := md5.New()
	b := []byte("any text")
	hash.Write(b)
	y :=  fmt.Sprintf("%x\n", hash.Sum(nil))
	fmt.Printf(y)
}