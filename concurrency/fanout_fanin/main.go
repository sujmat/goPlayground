package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3, 5)

	//Fan Out
	c1 := square(in)
	c2 := square(in)

	// Fan In
	for val := range merge(c1, c2) {
		fmt.Println(val)
	}
}

func gen(numbers ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range numbers {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(channels ...chan int) chan int {
	out := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(len(channels))
	for _, channel := range channels {
		go func(chn chan int) {
			for val := range chn {
				out <- val
			}
			wg.Done()
		}(channel)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
