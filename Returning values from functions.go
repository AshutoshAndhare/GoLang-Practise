package main
import (
  "fmt"
  "time"
)

// Add "string" as the return type of this function
func isItLateInNewYork() string {		//here we defined the name of the function and the data type of the return value we expect from this function
  var lateMessage string 
  t := time.Now()
  tz, _ := time.LoadLocation("America/New_York")//I think this one sets the time zone to America/New_York but don't understant the "tz,_"
  nyHour := t.In(tz).Hour()//after a bit of playing around, I've come to know that calculates the number of hrs from 12AM till time right now, but timezome of NY
  if nyHour < 5 {
    lateMessage = "Goodness it is late"
  } else if nyHour < 16 {
    lateMessage = "It's not late at all!"
  } else if nyHour < 19 {
    lateMessage = "I guess it's getting kind of late"
  } else {
    lateMessage = "It's late"
  }
  // Return the string lateMessage
  return lateMessage //this function returns the value of the var "lateMessage" as a string
}

func main() {
  var nyLate string
  nyLate = isItLateInNewYork() //the string of the var "lateMessage" in the isItLateInNewYork() is in turn called and saved in the main function in "nyLate" var.
  
  fmt.Println(nyLate)
}