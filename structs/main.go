package main

import (
	"encoding/json"
	"fmt"
)

type foo int

type Person struct {
	First      string
	Last       string `json:"-"`
	Age        int    `json:"wisdom_score"`
	unexported int
}

func (p Person) fullname() string {
	return p.First + " " + p.Last
}

func (p Person) greeting() string {
	return "Hello regular person"
}

type DoubleZero struct {
	Person
	LicenseToKill bool
}

func (d DoubleZero) agentStatus() string {
	killString := ""
	if d.LicenseToKill {
		killString = " has the license to kill"
	} else {
		killString = " does not have the license to kill"
	}
	return d.First + " " + d.Last + " " + killString
}

func (d DoubleZero) greeting() string {
	return "Greeting Agent " + d.fullname()
}

func main() {

	// Creating a custom type foo
	var myAge foo
	myAge = 44
	fmt.Printf("%T %v \n", myAge, myAge)

	//Creating a new struct object for type person
	p1 := Person{"first", "last", 20, 7}
	fmt.Println("P1 first name:", p1.First, ", last name:", p1.Last, ", age:", p1.Age, " private: ", p1.unexported)
	fmt.Println(p1.fullname())

	//Inheritance example
	d1 := DoubleZero{
		Person: Person{
			First: "James",
			Last:  "Bond",
			Age:   30,
		},
		LicenseToKill: true,
	}

	fmt.Println(d1.agentStatus())

	// Struct Overriding example
	fmt.Println(p1.greeting())
	fmt.Println(d1.greeting())

	//Struct pointers
	p2 := &Person{"James", "Lewis", 20, 8}
	fmt.Println(p2)
	fmt.Println("%T\n", p2)
	fmt.Println(p2.First)
	fmt.Println(p2.Age)

	// JSON marshaling Notice the elimination of unexported field and the json annotations
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))
}
