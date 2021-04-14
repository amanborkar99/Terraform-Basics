package main

import(
	"fmt"
)

func main(){
    //Initialsing map
	m := map[int]string{}

    m[1] = `COVID`
	m[2] = `-`
	m[3] = `19`

	//printing the key value pair using for loop
	for key, val := range m{
		fmt.Println(key, val)
	}

	fmt.Println(m)
    //deleting the key from map m
	delete(m, 2)

	fmt.Println(m)

}