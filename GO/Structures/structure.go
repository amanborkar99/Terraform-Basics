package main

import(
	"fmt"
)
type LGBT struct{
	marks, total_marks int
	percentage float64
	name string
}
type FoodNameGetter func(string) string
 
type Food struct {
    name string
    getter FoodNameGetter // declare function
}
func main(){
	var t LGBT = LGBT{marks: 178, total_marks: 550, name : "Rohit" }
	t.percentage = (float64(t.marks)/float64(t.total_marks)) * 100

	fmt.Println(t)
    s := t.name
	// function as field
    l := Food{"hello", func(s string) string {
		return s
	}}
	fmt.Println(l, s)
}