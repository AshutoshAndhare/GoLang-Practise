package main

import "fmt"
//Varadic Functions can be called with any number of trailing arguments.
func sum(nums ...int) {//like here we just declared one var with a data type int but ... as the prefix of the int the shows any number of arguments are valid
	fmt.Print(nums, " ")//this part just tells the function to print whatever arguments were called.
	total:= 0
	for _,num := range nums {//this is the calculation part where we use the range function so that n number of arguments are made possible
		total += num
	}
	fmt.Println(total)
}
func main()  {
	sum(1,2,7,8)//varadic functions are called in the usual way with individual arguments seperated by commas
	sum(3,4,5)

	nums:= []int{1,2,3,4,5,7,8}//if we have a slice or an array, we can also input that to the function as an argument but with the func (slice...)
	sum(nums...)//like this
}