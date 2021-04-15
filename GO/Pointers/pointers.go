package main

import(
	"fmt"
)

func main(){
	// normal declaration
	var p *int
	 t := 32
     p = &t
	 
	 fmt.Println(*p)
    
	 var x = 121
	 var m *int = &x
	 fmt.Println(m, *m)

	//declaration by new keyword

	w := new(string)
	r := `abcd`

	w = &r

	fmt.Println(r, w, *w)

}