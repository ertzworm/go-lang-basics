package main

import ("fmt")

func main(){

	//@ & - STANDS FOR ADDRESSES
	//@ * - STANDS FOR VALUES
	an_integer := 15
	the_address := &an_integer

	fmt.Println(the_address)
	fmt.Println(*the_address)

	//@ CHANGES THE VALUE STORED IN THE ADDRESS TO 10
	*the_address = 10
	fmt.Println(an_integer)

}