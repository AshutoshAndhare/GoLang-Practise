package main

import "fmt"

func main()  {
	var a [5]int//here we declare the array simply that contains exact 5 elements and the type of array and length are both part of array's type
	fmt.Println("emp:",a)

	a[4] = 100//we can also access and set value at an index using this syntax, where index numbers start from 0
	fmt.Println("set:",a)
	fmt.Println("get:",a[4])

	fmt.Println("len:",len(a))//we can see the length of an array using len().

	b := [5]int{1,2,3,4,5}//we can also declare and initialize the array in the same line using this syntax
	fmt.Println("dcl:",b)
	
	var twoD [2][3]int//we can also make multi dimentional arrays, like here we made a 2-dimentional array and initialized it using a for loop
	for i:=0;i<2;i++ {
		for j:=0;j<3;j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:",twoD)
}