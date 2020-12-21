package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  isHeistOn := true
  eludeGuards := rand.Intn(100)

  if eludeGuards >= 50 {
    fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
  } else {
    isHeistOn = false
    fmt.Println("Plan a better disguise next time?")
  }
  openedVault := rand.Intn(100)
  if openedVault >= 70 && isHeistOn {
    fmt.Println("Grab and Go!!")
  } else if isHeistOn {
    isHeistOn = false
    fmt.Println("The vault can't be opened.")
  }
  leftSafely := rand.Intn(5)
  if isHeistOn {
    switch leftSafely {
      case 0 :
        isHeistOn = false
        fmt.Println("The manager called the poilce while you were in the vault and you got caught")
      case 1 :
        isHeistOn = false
        fmt.Println("The vault alarm went off and while you were in it the police arrived.")
      case 2 :
        isHeistOn = false
        fmt.Println("The vault door closed when you were in it. Turns out vaults can't be opened from inside...")
      case 3 : 
        isHeistOn = false
        fmt.Println("Turns out as soon as the vault is opened unauthorised, 20 special Bank Security guards are sent...")
      default :
        fmt.Println("Start the Getaway Car!") 

    }
  }
  if isHeistOn {
    amtStolen := 10000 + rand.Intn(1000000)
    fmt.Println("The amount Stolen is :", amtStolen)
  }
  fmt.Print("Heist status:",isHeistOn)

}
