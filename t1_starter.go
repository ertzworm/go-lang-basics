package main

//@ Imports all required dependencies for the application
//@ In this case, fmt stands for format
import ("fmt")

//@ func main will always run 
func main() {

	//@ Initialization of variables
	var x int = 5
	var y int = 10
	var sum int = x + y
	
	//@ Array
	var a [5] int

	//@ Short-hand expressions
	// x := 5
	// y := 10
	// sum := x + y
	fmt.Println(x,y,sum)
	fmt.Println(a)
}