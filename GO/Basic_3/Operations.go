package main

import(
	"fmt"
	"math"
	"strconv"
)

func main(){
	v := 8
	var x byte = 3
    //Type Casting

	fmt.Println(v + int(x) )
	
	v += int(x)

	fmt.Printf("The value of addition of v %q \n", strconv.Itoa(v))

	fmt.Printf("The standard value of Pi is %v and its datatype is %T\n", math.Pi, math.Pi)

	// Logical as well as Bitwize operator

	r := 16
	m := 22

	fmt.Printf("Bitwise OR of r & m is %v \n", (r | m))

	fmt.Printf("Is r & m logically equivalent ? %v \n", r == m)





}