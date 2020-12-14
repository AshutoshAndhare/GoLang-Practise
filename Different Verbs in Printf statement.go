package main

import "fmt"

func main() {
  floatExample := 1.75
  // Edit the following Printf for the FIRST step
  fmt.Printf("Working with a %T.", floatExample) //%T stands for Type of the variable after comma
  
  fmt.Println("\n***") // Added for spacing
  
  yearsOfExp := 3
  reqYearsExp := 15
  // Edit the following Printf for the SECOND step
  fmt.Printf("I have %d years of Go experience and this job is asking for %d years", yearsOfExp, reqYearsExp)//%d interpolates an interger in a string 
  
  fmt.Println("\n***") // Added for spacing
  
  stockPrice := 3.50
  // Edit the following Printf for the THIRD step
  fmt.Printf("Each share of Gopher feed is $%.2f!", stockPrice) //%f stands for float incorporation
  //%.2f gives 2 decimal places of the float hence here it prints 3.50
}