package main

import "fmt"

func foo(c chan int, someVal int) { //here we need 2 parameters 1st is channel c of int type and var someVal of int type 
	c <- someVal * 5 //here we multiply someVal by 5 and put that value in the channel c
}

func main()  {
	fooVal := make(chan int) //this is how we create channels, this one is of int type and chan stands for channel and not var

	go foo(fooVal, 5) //here we specify the channel name and tell the foo() to store that value to that specific channel
					  //and all this in a go routine
	go foo(fooVal, 3)

	v1 := <-fooVal //here we store the value of one of the channels in v1 and other channel's in v2, whithout sequence
	v2 := <-fooVal

	fmt.Println(v1, v2) //it'll mostly return 15 and then 25 but no surprise if the values interchange
}