package main

import "fmt"

func main()  {
	//First of all map is an associate data type, built-in GoLang.
	m := make(map[string]int)//This is how we declare a map using the make statement i.e. make(map[Key type]value type)

	m["k1"] = 17//here k1 is the key which is a string and 17 is the value which is int
	m["k2"] = 56

	fmt.Println("map:", m)//the output here will be in syntax of map[key:value, key:value]

	v1 := m["k1"]//this is how we access the value of the key "k1" of map "m"
	fmt.Println("v1", v1)

	fmt.Println("len:", len(m))//obviously the length of the map which shows how many key:value pairs are present

	delete(m, "k1")//we can directly delete the key this way
	fmt.Println("map:", m)

	_,prs := m["k2"]//here we see if so and so key exists in this map or not, here we didn't want the value of the key so we left blank "_" before the comma 
	fmt.Println("prs:", prs)//if we want to save the key value as well we can write the above line as val, prs := m["k2"] , where val is a variable which will store value of k2

	n := map[string]int{"foo": 1, "bar": 3, "aon": 5}//this is how we can declare a complete map and initialize it in the same line
	fmt.Println("map:", n)//note that the output is sorted alphabetically from a-z
	
}