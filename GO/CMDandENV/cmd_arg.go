package main

import(
	"fmt"
	"os"
)
func main(){
	t := os.Args
	r := t[1:]

	fmt.Print(t, r)

	//environment variables

	os.Setenv("CSGO", "No more played")
	os.Setenv("Valorant", "GG")
	l := os.Getenv("CSGO")
	b := os.Getenv("Valorant")

	fmt.Println(l , b)

	os.Unsetenv("CSGO")
	os.Unsetenv("Valorant")
	 fmt.Println(os.Environ())
}