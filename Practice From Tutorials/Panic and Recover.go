package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cleanup() { 
	defer wg.Done()	//5.) then after recovery we set waitgroup to done so that the program could go forward
	if r := recover(); r != nil { //3.) here we set r for recover() so that if recover is != nil, we print the below line
		fmt.Println("Recovere in Cleanup:", r) //4.) and also return the panic message
	}
}
//6.) But, here we're going through a issue where the program stops after panic and recovery and we will
//do something about it after we understand what's the issue with that
func say(s string) {
	defer cleanup() //2.) then, whenever panic occurs, we set the cleanup() to run as soon as panic starts
	for i := 0; i < 10; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond*100)	
		if i == 2 {
			panic("Oh dear, a 2") //1.) here we have set the loop to return a panic if the i = 2 
		}
	}
}

func main() {
	wg.Add(1)
	go say("Hey")
	wg.Add(1)
	go say("There")
	wg.Wait()
}