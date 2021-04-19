package main

import (
	"fmt"
	"strings"
	//"reflect"
	//"os"
	"GO/Basic_5"
	"encoding/json"
)

type xyz struct{
	P int `json:"ID"`
	T string `json:"Name"`
}

func main(){
	r := xyz{
		P: 22,
	    T: "Aman Borkar", 
	}

	b, _ := json.MarshalIndent(r,""," ")
	
	fmt.Println(strings.NewReader(string(b)))
    
	//fmt.Println(b)
	fmt.Println(json.Valid(b))

	m := xyz{}

	err := json.Unmarshal(b, &m)

	if(err != nil) {
      fmt.Printf("error\n")
	}else{
		fmt.Println(m)
	}

	fmt.Println(t2.Division(14,5))
}