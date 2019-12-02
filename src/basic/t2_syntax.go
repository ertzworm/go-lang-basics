package main

//@ IMPORT BOTH format and MATH packages
import("fmt"
			"math")

func main() {				
	//@ DIFFERENCES IN METHODS
	//@ fmt.Println() - If the first letter is in upper-case, it's a package method
	//@ main() - else, custom made or user created
	fmt.Println("The square root of 4 is", math.Sqrt(4))
}