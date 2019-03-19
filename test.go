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

func addnumbers(first, second int) int{
	return first + second
}

func main(){
	//fmt.Println("hello world!:)")
	//testfunc()
	fmt.Println(addnumbers(5,6))
}
