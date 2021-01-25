package main

import "fmt"
//Animal interface is just a type that takes in method Speak() i.e. it will show us animals that speak.
type Animal interface {
	Speak() string
}

//Dog struct is just a declaration of an animal in the Animal interface
type Dog struct {
}
//Speak function here is we telling what does the Animal Dog speak
func (d Dog) Speak() string {
	return "Woof!"
}

//Cat struct declaration
type Cat struct {
}
//Speak function tellin what does the cat say
func (c Cat) Speak() string {
	return "Meow!"
}

//Llama struct declaration
type Llama struct {
}
//Speak function telling Llama speaks "????"
func (l Llama) Speak() string {
	return "????"
}

//JavaProgrammer struct declaration
type JavaProgrammer struct {
}
//Speak function telling JavaProgrammer says "Design Patterns!!"
func (j JavaProgrammer) Speak() string {
	return "Design Patterns!!"
}

//Now we declare a slice of var animals that consists of the structs from the interface Animal
func main() {
	animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{}}
	//Below we use range to call out the Speak function in each struct's second place, hence '_, animal'
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
