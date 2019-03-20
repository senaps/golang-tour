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

/*
var programmer  bool// default is false!
var yearsOfGOExperience  int//default is zero

func main(){
	fmt.Printf("%T:%v\n", programmer, programmer)
	fmt.Printf("%T:%v\n", yearsOfGOExperience, yearsOfGOExperience)
	yearsOfGOExperience = 1
	fl := float64(yearsOfGOExperience) // casting the fucking int to float64
	fmt.Printf("%T:%v\n", fl, fl)
}
*/

/*
const likesGo = true  // how to make a const!

func main(){
	// const thing := "something" wont work! const only works with =
	fmt.Println(likesGo)
}
*/

/*
func main(){
	for i := 0; i < 3; i++{
		fmt.Println("looping...")
	}

	i := 0
	for ;i<5;{
		fmt.Println("still looping...")
		i++
	}

	i = 0
	for i<5 {
		fmt.Println("looping while style...")
		i++
	}

	i = 0
	for{
		fmt.Println("looping forever:)))")
		i++
	    fmt.Println(i)
	}
}
*/

/*
func main(){
	name := "maysam"
	
	if name == "maysam"{
		fmt.Println("okay!")
	}
	
	nick := "senaps"
	if nick == name{
		fmt.Println("no nicks...")
	}else{
		fmt.Println("you have a nick!")
	}
}

*/

/*
func main(){
	age := 27 
	switch {
	case age < 10:
		fmt.Println("fuck off!")
	case age < 20:
		fmt.Println("still, fuck off...")
	case age > 20:
		fmt.Println("go fuck yourself!")
	}
}
*/

/*
func main(){
	defer fmt.Println("world?") // runs when function is finished execution...
	fmt.Println("hello")
}
*/

/*
func sayBye(name string){
	fmt.Println("starting bye")
	defer fmt.Printf("Bye bye %v...\n", name)
	fmt.Println("finishing bye")
}

func sayHello(name string){
	fmt.Println("starting sayHello")
	defer fmt.Printf("hello %v\n", name)
	defer sayBye(name)
	fmt.Println("finishing sayHello")
}

func main(){
	fmt.Println("starting main")
	sayHello("maysam")
	fmt.Println("finishing main")
	//executes as:
	    starting main
		starting sayHello
		finishing sayHello
		starting bye
		finishing bye
		Bye bye maysam...
		hello maysam
		finishing main
}
*/

func main(){
	fmt.Println("bye now!")
}