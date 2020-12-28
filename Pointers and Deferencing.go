package main

import "fmt"

func main() {
	star := "Polaris"//here we added a normal var with a string value
	
	starAddress := &star//here we created a var to save the address of "var star"
	
	*starAddress = "Sirius"//here we told the pointer of the var that has address of star to be equal to "Sirius"
	//that way we indirectly change the value of star without touching the actual variable just by changing the values present at the address of 'var star'
  
  fmt.Println("The actual value of star is", star)//and here we can check that the value of star is changed.
}
