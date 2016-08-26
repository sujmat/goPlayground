package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Helps in composition of independently executing processes. Dealing with lots of things at once.
// Usually might interleave the processes. Different from processes.
// Concurrency runs on all cores as indicated in init

var wg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func foo() {
	for i := 0; i < 400; i++ {
		fmt.Println("i= ", i)
		time.Sleep(3 * time.Millisecond)
	}
	fmt.Println("i is done")
	wg.Done()
}

func bar() {
	for j := 0; j < 200; j++ {
		fmt.Println("j= ", j)
		time.Sleep(20 * time.Millisecond)
	}
	fmt.Println("j is done")
	wg.Done()
}

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
	fmt.Println("Both done")
}
