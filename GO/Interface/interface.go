package main

import(
	"fmt"
	"math"
)
// interface
type common_goal interface{
     area() float64
}

type sq struct{
    side float64
}

type ci struct{
	radius float64
}
func (s sq) area()float64{
	return s.side * s.side
}
func (c ci) area()float64{
	return math.Pi * math.Pow(c.radius, 2)
}

func main(){
	 var v common_goal = sq{22.5}
	 var r common_goal = ci{12.5}

	 fmt.Printf("Area of the square is %f & Area of the Circle is %f \n", v.area(), r.area() )
	 
}