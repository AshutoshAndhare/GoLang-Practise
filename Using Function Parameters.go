package main
import "fmt"

func computeMarsYears(earthYears int) int {//we defined earthYears as an int variable as a parameter here 
    
  earthDays := earthYears * 365
  marsYears := earthDays / 687
  return marsYears//now we defined the return var here but the return variable type was defined along with declaration of the function name above
}

func main() {
  myAge := 25
  
  myMartianAge := computeMarsYears(myAge)//now the value of myAge as 25 will be passed to computeMarsYears() as var earthYears being parameter of computeMarsYears()
  //that value of earthYears is now 25, which will be used to calculate marsYears. Which will be the used as return value for computeMarsYears() which will again be transfered to main()
  //which then assigns the value to myMartianAge.
  fmt.Println(myMartianAge)
}