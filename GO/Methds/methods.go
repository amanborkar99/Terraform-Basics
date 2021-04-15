package main

import (
	"fmt"
)

type full_name struct {
	f_name, l_name string
}

func (c full_name) cnt() string {
	return c.f_name + " " + c.l_name
}

func main() {

	n := full_name{
		f_name: "Aman",
		l_name: "Borkar",
	}

	fmt.Println(n.cnt())

}
