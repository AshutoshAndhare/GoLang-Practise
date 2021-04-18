package main

import "fmt"

func add(a,b float64) float64 {
	return a+b
}

func main()  {
	a := add(3.2, 5.6)
	//above is a simple add function the returns a value in float64 stored in var a.
	//so if we want to change that type in int or something else, we can use "int(a)" as given below to print the int type
	//of the var a 

	fmt.Println(int(a))

	//we can also store this changed int type of a in a variable as below

	b := int(a)
	fmt.Println(b)
}