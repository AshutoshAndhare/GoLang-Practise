package main

import "fmt"

func main() {
  fmt.Println("What would you like for lunch?")
  
  var food string
  fmt.Scan(&food)//here we ask for the user input which is stored in the 'food' variable, & represents the address of the variable 
  //as Scan Statement expects address of the variables, we'll learn the 'why?' later
  fmt.Printf("Sure, we can have %v for lunch.", food)//Here we called in the value entered by the user which we stored in the 'food' variable
  //the only thing is that Scan Statements take the value that the user enters before space
  //like here if you enter 'Sabji' as an answer, it'll Print %v as Sabji, but if you enter 'Sabji Roti' as answer, it'll Print: Sabji in place of %v
}
