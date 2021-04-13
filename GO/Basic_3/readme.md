# basic Operations in GO

Libraries used :- 
1. **fmt**
2. **math**
3. **strconv**

* Created a variable and typecasted one of them to do the arithmetic operation 
 ``` Go
    v := 8
	var x byte = 3
    fmt.Println(v + int(x) )
	
	v += int(x)

	fmt.Printf("The value of addition of v %q \n", strconv.Itoa(v))
```
here, variable x is typecasted to int and then added to x and then v is converted to string

* Performed `OR` operation on two varaible `r` and `m` and checked their equivalence through `==` operator.
``` Go
    r := 16
	m := 22

	fmt.Printf("Bitwise OR of r & m is %v \n", (r | m))

	fmt.Printf("Is r & m logically equivalent ? %v \n", r == m)
```
***
OUTPUT : -
``` Output
11
The value of addition of v "11"
The standard value of Pi is 3.141592653589793 and its datatype is float64
Bitwise OR of r & m is 22
Is r & m logically equivalent ? false
```
