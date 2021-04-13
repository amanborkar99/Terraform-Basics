package main

import (
	"fmt"
)

func main(){
	//complete declaration
	var x int = 12

	// declaration without type if initialzed during declaration
	var r = 99.99

	// short declaration
    t := "hello Google, can you do mojo?"

	fmt.Printf("This is %T variable with %v value\n",x,x)

	fmt.Printf("The variable r is of type %T and %v value\n",r,r)

	fmt.Printf("This is %T type variable with %v value",t,t)

	fmt.Println(x,r,t)
	 
}