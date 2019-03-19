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

func addNumbers(first, second int) (int, string){
	return first + second, "No error:)"
}

func main(){
	//fmt.Println("hello world!:)")
	//testfunc()
	res, errorString := addNumbers(5,6)
	fmt.Println(res, errorString)
}
