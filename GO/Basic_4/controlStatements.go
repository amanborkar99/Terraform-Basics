package main

import(
	"fmt"
)

func main(){
    
	//declaration with condition and if short declaration
	x:=12
	if x%2 == 0 {
	  fmt.Printf("%d is even number. \n", x)
	}else{
       fmt.Printf("%d is an odd number.\n", x)
	}

	//switch short declaration
	switch x:=3; x {
	case 1: fmt.Println("value of x is 1")

	case 2: fmt.Println("value of x is 2")

	default : fmt.Println("value of x is not 2 as well as not 3.")
		
	}

	//for loop declaration
	fmt.Printf("\nTable of 12 is as follows :\n")



	for i:=0;i<=10;i++{
		
        fmt.Printf("12 %v cha %v\n", i, i*12)
	}
}