package main

import(
	"fmt"
)


/*func testfunc(){
	fmt.Println("hello wrold from a function:)")
}
*/
/*
func addnumbers(first int, second int) int{
	return first + second
}
*/

func addNumbers(first, second int) (sum int, errString string){
	// we have added names for return shits...
	sum = first + second
	//errString = "None"  // we will even work without this:)
	return sum, errString
}

func main(){
	//fmt.Println("hello world!:)")
	//testfunc()
	result, some := addNumbers(5, 6)  // we need to use := wtf is the difference?
	fmt.Println(result, some)
}
