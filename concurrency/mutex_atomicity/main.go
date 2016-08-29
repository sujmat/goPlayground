package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
	"sync/atomic"
)


var wg sync.WaitGroup
var mut sync.Mutex
var valRace int
var val int
var atomicVal int64

func increment_with_race(s string) {
	for i := 0; i < 20; i++ {
		x := valRace
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		valRace = x
		fmt.Println(s, " ", valRace)
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

func increment_with_atomicity(s string){
	for i:=0; i<20; i++{
		atomic.AddInt64(&atomicVal,1)
		fmt.Println(s," ",atomicVal)
	}
	wg.Done()
}

func main() {
	// Use go run -race main.go to detect races.

	// increment with a race conditon
	wg.Add(2)
	go increment_with_race("Foo: ")
	go increment_with_race("Bar: ")
	wg.Wait()
	fmt.Println("----------")

	// increment with a mutex
	wg.Add(2)
	go increment("Foo: ")
	go increment("Bar: ")
	wg.Wait()
	fmt.Println("-----------")

	// increment with atomicity
	// Notice the order of printing is off even though the final value is correct
	wg.Add(2)
	atomicVal = 0
	go increment_with_atomicity("Foo: ")
	go increment_with_atomicity("Bar: ")
	wg.Wait()
	fmt.Println("-----------")


}
