package main

import (
	"fmt"
  "math/rand"
  "time"
)

func main() {
	// Add your code below:
  fmt.Println(time.Now().UnixNano())//Prints time elapsed from 1 Jan 1970 till now with nano second accuracy
  rand.Seed(time.Now().UnixNano())//give random number generator a seeding number that changes every nano seconds 
                                  //so that we get new random number every time we run the program
  amountLeft := rand.Intn(10000)//give us a random integer with max value of 10,000
  
  fmt.Println("\nAmountLeft is: ", amountLeft)
  
	if amountLeft > 5000 {
		fmt.Println("What should I spend this on?")
  } else {
    fmt.Println("Where did all my money go?")
  }
}
