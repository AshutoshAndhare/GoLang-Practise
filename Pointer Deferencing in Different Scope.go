package main

import "fmt"

func brainwash(saying *string) {//here we tell to func to accept an argument which is pointer to string.
	*saying = "Beep Boop."//here we tell the pointer to change to "beep boop"
}

func main() {
	greeting := "Hello there!"
	
	// Call your brainwash() below:
	brainwash(&greeting)//here we give a argument i.e. address of the var greeting. which in turn becomes "*saying = &greeting" 
	//meaning telling the pointer to pointer to greeting(var saying) to take the address of greeting(&greeting) and then in the brainwash() change the pointer to greeting(var saying) to Beep Boop
	
	fmt.Println("greeting is now:", greeting)
}
