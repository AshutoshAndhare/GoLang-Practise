package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup //here we define the Waitgroup variable

func foo(c chan int, someVal int) { 
	defer wg.Done()
	c <- someVal * 5  
}

func main()  {
	fooVal := make(chan int, 10) //here along with creating a channel, we mention its buffer as 10

	for i := 0; i < 10; i++ {
		wg.Add(1) //tell waitgroup to add 1 each time
		go foo(fooVal, i) //here we specify channels and input i as the value to foo() and take thos values back to the channels
	}
	wg.Wait() //tell wg to wait till all goroutines are completed

	close(fooVal) //close the channel after all the goroutines are completed and buffer of 10 helps the channels to buffer 

	for item := range fooVal { //iterating over the the channels using range after goroutines and channels are closed
	fmt.Println(item) 
	}
}