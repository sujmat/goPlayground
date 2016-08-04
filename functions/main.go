package main

import "fmt"

func main(){
	greet("FirstN","LastN")

	s2 := greet2("FirstN","MiddleN","LastN")
	fmt.Println(s2)

	s3,s4 := greet3("FirstN","LastN")
	fmt.Println(s3," ",s4)

	variadicGreet("FirstN", "SecondN", "ThirdN", "FourthN")

	// breaking open the slice
	name := [] string{"FName","MName","LName"}
	variadicGreet(name...)

	//Func expression
	greeting := func(fName,lName string) {
		fmt.Println(fName, lName)
	}
	greeting("FName","LName")

	//Closure
	greeter := makeGreeterClosure()
	fmt.Println(greeter("FName","LName"))

	//Callback examples
	accumulate([]int{1,2,3,4,5,6},func(total int){
		fmt.Println(total)
	})

	// Callback that takes in custom filter functions. Here filtering odd numbers from list
	filteredList := filter([]int{1,3,5,4,2,6},func(val int) bool{
		return val%2==0
	})
	fmt.Println(filteredList)

	//Recursion example
	fmt.Println(factorial(5))

	//Defer example
	deferLastName("FirstN","MiddleN","LastN")

	//Anonymous self executing function
	func(){
		fmt.Println("\nJust executed.")
	}()
}

// Function with no return
func greet(fname, lname string){
	fmt.Println(fname, " ", lname)
}

// Function with return value
func greet2(fname, mname, lname string) string{
	return fmt.Sprint(fname,mname,lname)
}


// Function with multiple returns
func greet3(fname, lname string) (string,string){
	return fmt.Sprint(fname,lname), fmt.Sprint(lname,fname)
}

//Function with varying number of parameters
func variadicGreet(name ...string) {

	for _,element := range name {
		fmt.Print(element," ")
	}
	fmt.Println("")
}

// Closure that returns the function that greet function
func makeGreeterClosure() func(fName,lName string) string{
	return func(fName,lName string) string{
		return fmt.Sprint("Hello",fName,lName)
	}
}

//Callback example
func accumulate(numbers []int, callback func(int)){
	total := 0
	for _,n := range numbers{
		total+=n
	}
	callback(total)
}

//Callback example  with return value
func filter(numbers []int, comparisonFunction func(int) bool) []int{
	filteredArr := []int{}
	for _,val := range numbers{
		if comparisonFunction(val){
			filteredArr = append(filteredArr,val)
		}
	}
	return filteredArr
}

//Recursion example
func factorial(x int) int{
	if x == 0{
		return 1
	}
	return x * factorial(x-1)
}

//Defer printing last name until end of the function
func deferLastName(fName, mName, lName string){
	defer fmt.Print(lName)
	fmt.Print(fName," ")
	fmt.Print(mName," ")
}