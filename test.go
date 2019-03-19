package main

import(
	"fmt"
)

/*
func main(){
	fmt.Println("hello world!")
}
*/

/*
func testfunc(){
	fmt.Println("hello wrold from a function:)")
}
func main(){
    testfunc()
}
*/

/*
func addnumbers(first int, second int) int{
	return first + second
}

func main(){
    fmt.Println(addnumbers(5,6))
}
*/

/*
func addNumbers(first, second int) (sum int, errString string){
	// we have added names for return shits...
	sum = first + second
	//errString = "None"  // we will even work without this:)
	return sum, errString
}

func main(){
	result, some := addNumbers(5, 6)  // we need to use := wtf is the difference?
	fmt.Println(result, some)
}
*/

var name, family string  // this how we make variables!

func main(){
	name = "maysam"  // this how we assign data to a variable
	var age int = 27  // this how we make a variable with data 
	//address = "Tehran"  // this wont work!
	var address = "Tehran"  //this works! so no need of the variable type
	fmt.Println(name, family, age, address)
}