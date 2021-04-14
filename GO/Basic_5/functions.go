package main

import(
	f "fmt"
)
// return variable declared already
func Division(x float64, y int) (r float64){
   r = x /float64(y)
    return
}

func authenticator(name, password string) bool{
	if(name == "aman" && password == "xyz") {
		return true
	}else{
		return false
	}
}
func main(){

	f.Printf("Division of 14 by 3 is %v\n", division(14,3))

	var (
		name, pass string
	)

	f.Scanf("%s %s",&name,&pass)

	if(authenticator(name, pass)){
		f.Println("This is a valid user")
	}else{
		f.Println("Above enter pass and name is incorrect")
	}


}