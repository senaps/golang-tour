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

/*
var name, family string  // this how we make variables!

func main(){
	name = "maysam"  // this how we assign data to a variable
	var age int = 27  // this how we make a variable with data 
	//address = "Tehran"  // this wont work!
	var address = "Tehran"  //this works! so no need of the variable type
	fmt.Println(name, family, age, address)
}
*/

/*
var name, nick, programmer = "maysam", "senaps", true  // this how we will do multiple variables

func main(){
	fmt.Println(name, nick, programmer)
}

*/

/*
// name := "maysam" //wont work, cause we don't have := outside of functions!
var name = "maysam"

func main(){
	nick := "senaps"
	fmt.Println(name, nick)
}
*/

/*
var(
	name string = "maysam"
	age = 27
	programmer = true
)
// shit! :) we can do one liner stuff too with var
func main(){
	fmt.Printf("concating with T as type: %T, value is: %v\n", name, name)
	fmt.Printf("type: %T, value is: %v", programmer, programmer)
	fmt.Printf("what the heck is the \n doing at the end? we forced to use it?")
	//if stupid `\n` is not used, next print comes appended to current line
	fmt.Println("do we have this effect here too?")
	fmt.Println("next line to check!")
	// so `Println` goes to next line automatically, but `Printf` does not!
}
*/


var programmer  bool// default is false!
var yearsOfGOExperience  int//default is zero

func main(){
	fmt.Printf("%T:%v\n", programmer, programmer)
	fmt.Printf("%T:%v\n", yearsOfGOExperience, yearsOfGOExperience)
	yearsOfGOExperience = 1
	fl := float64(yearsOfGOExperience) // casting the fucking int to float64
	fmt.Printf("%T:%v\n", fl, fl)
}