package main

//@ To use functions from packages, you must import them to
//@ the last directory using slashes, in this case, "math/rand"
//@ functions can be used as rand.Intn(100) and not math.rand.Intn(100)
import ("fmt"
			"math/rand")

func main() {
	fmt.Println("A number from 1-100:", rand.Intn(100))
}