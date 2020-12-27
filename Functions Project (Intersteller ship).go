package main

import "fmt"

// Create the function fuelGauge() here
func fuelGauge(fuel int) {
  fmt.Println("The amount of fuel is:", fuel)
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int {
  var fuel int
  switch planet {
    case "venus":
      fuel = 300000
    case "mercury":
      fuel = 500000
    case "mars":
      fuel = 700000
    default:
      fuel = 0
  } 
  return fuel
}

// Create the function greetPlanet() here
func greetPlanet(planet string) {
  fmt.Println("Hey, welcome to", planet, "!! May you have a comfortable and safe stay here!")
}

// Create the function cantFly() here
func cantFly() {
  fmt.Println("We do not have the available fuel to fly there.")
}

// Create the function flyToPlanet() here
func flyToPlanet(planet string, fuel int) int {
  var fuelRemaining, fuelCost int
  fuelRemaining = fuel
  fuelCost = calculateFuel(planet)
  if fuelRemaining >= fuelCost {
    greetPlanet(planet)
    fuelRemaining -= fuelCost
  } else {
    cantFly()
  }
  return fuelRemaining
}

func main() {
  // Test your functions!
  
  // Create `planetChoice` and `fuel`
  var fuel int
  fuel = 1000000
  planetChoice := "venus"
  // And then liftoff!
  fuel = flyToPlanet(planetChoice, fuel)
  fuelGauge(fuel)

}