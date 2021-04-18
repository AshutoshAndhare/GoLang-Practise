package main

import "fmt"

func main()  {
	x := 15
	a := &x //&x is the address of x
	fmt.Println(a) //it'll just show the memory address of x
	fmt.Println(*a) //it'll read what is at the memory address of x, i.e. "15". Hence, *a is the pointer of x

	*a = 5 //if we change the value of the pointer of x, we change the value of x
	fmt.Println(x) //this will print 5 since the value of x is changed rn.

	*a = *a**a //just for fun, here we change pointer of x by multiply pointer of x to itself
	fmt.Println(x) //it'll print 25

}