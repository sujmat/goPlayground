package main

import (
	"fmt"
)

func main() {
	chn := make(chan int)

	go func() {
		for i := 0; i < 20; i++ {
			chn <- i
		}
		close(chn)
	}()

	// range loops till channel is closed to exit
	for n := range chn {
		fmt.Println("chn: ", n)
	}

}
