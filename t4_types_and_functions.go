package main

import ("fmt")

//@ Format: 
//@ func <function_name> (parameter1 type1, parameter type2) return type
func add(x float64, y float64) float64 {
	return x + y
}

//@ getName accepts two string params and returns 2 strings
func getName(word1, word2 string) (string, string) {
	return word1, word2
}

func main() {

	//@ Regular explicit syntax
	// var first_number float64 = 5.6
	// var second_number float64 = 11.5

	//@ Shorthand
	// first_number := 5.6
	// second_number := 11.5

	//@ Unpacking
	// var first_number, second_number float64 = 5.6, 11.5

	//@ Unpacked shorthand
	first_number, second_number := 5.6, 11.5
	first_word, second_word := "Paul", "Guiao"

	
	result := add(first_number, second_number)
	fmt.Println(result)
	fmt.Println(getName(first_word, second_word))
}