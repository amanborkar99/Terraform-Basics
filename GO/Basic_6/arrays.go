package main

import(
	"fmt"
)

func main(){
	//initialsed the array
	a := [10]int{1,4,5,7,8,9}

	fmt.Println(a)

	//iterating using for loop
	for i:=0;i<len(a);i++{
		fmt.Printf("%v\t",a[i]*2)
	}

	// iterating by using range
	for ind, val := range a{
		fmt.Println(ind, val)
	}

	//slicce of a
	fmt.Println(a[1:4])
	fmt.Println(a[:4])
	fmt.Println(a[3])
	fmt.Println(a[2:])
	fmt.Println(a[:])

	//empty array and slice
	var x []int
	fmt.Println(x)
	//appending elements to slice
	fmt.Println(append(x[:],1,3,4))
	fmt.Println(x)
}