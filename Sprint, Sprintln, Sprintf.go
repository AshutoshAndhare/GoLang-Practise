package main

import "fmt"

func main() {
  step1 := "Breathe in..."
  step2 := "Breathe out..."
  
  // Add your code below:
  meditation := fmt.Sprint(step1,step2)//works like Print(), didn't add spaces between the two
  fmt.Print(meditation)//Prints: Breath in...Breath out...
  
  fmt.Print("\n")
 
  meditation = fmt.Sprintln(step1,step2)//works like Println(), adds space between step1 and step2
  fmt.Print(meditation)//Prints: Breath in... Breath out...


  template := "I wish I had a %v."//%v here is a verb
  pet := "puppy"
  
  var wish string
  
  wish = fmt.Sprintf(template,pet)//works like Printf(), uses the %v as a verb and places pet="puppy" in the %v placeholder
  
  fmt.Println(wish)//Prints: I wish I had a Puppy.
}
