package main

import "fmt"

func main()  {
	type Person struct {
		name string
		age int
	}

	type Team struct {
		name string
		players [2]Person
	}

	var myTeam Team
	fmt.Println(myTeam)

	players1 := [...]Person{Person{name: "Forrest",age: 34}, Person{name: "Gordon", age: 27}}
	
	celtic := Team{name: "Celtic FC", players: players1}
	fmt.Println(celtic)
}