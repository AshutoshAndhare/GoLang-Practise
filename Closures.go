package main

import "fmt"
//Closures are also known as anonymus functions. We use these functions when we don't want to name a function and they are functions inside functions
//which return functions.
func intSeq() func() int {  //the return value here is func() int
	i := 0
	return func() int  {    //defining the return as return here is a function
		i++
		return i     //here this is the return of func()
	}
}

func main()  {
	nextInt := intSeq()  //now saving the return of intSeq() in var nextInt

	fmt.Println(nextInt())  //now the return of intSeq() was a func so now nextInt acts as a func so here we use nextInt()
	fmt.Println(nextInt())  //also if we use nextInt instead of nextInt(), the output comes is in hexadicimal form. *don't know what that means*
	fmt.Println(nextInt())

	nextInts := intSeq()

	fmt.Println(nextInts())  //also here, we can see that if we create and start with another variable, the sequence starts from one again.
	fmt.Println(nextInts())
}