package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var val int
var wg sync.WaitGroup
var mut sync.Mutex

func increment_with_race(s string) {
	for i := 0; i < 20; i++ {
		x := val
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		val = x
		fmt.Println(s, " ", val)
	}
	wg.Done()
}

func increment(s string) {
	for i := 0; i < 20; i++ {
		mut.Lock()
		x := val
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		val = x
		fmt.Println(s, " ", val)
		mut.Unlock()

	}
	wg.Done()
}

func main() {
	// Use go run -race main.go to detect races.
	wg.Add(2)
	go increment("Foo: ")
	go increment("Bar: ")
	wg.Wait()
}
