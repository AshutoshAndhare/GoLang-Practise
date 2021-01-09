package main

import "fmt"

type person struct {//this person struct type has name and age fields
	name string
	age int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p//you can safely return the pointer to local variable as a local variable will survive the scope of the function
	}

func main() {
	fmt.Println(person{"Bob", 35})//this syntax creates a new struct
	
	fmt.Println(person{name: "Alice", age: 24})//you can name the field when initializing the struct
	
	fmt.Println(person{name: "Fred"})//omitted fields will be zero valued
	
	fmt.Println(&person{name: "Ann", age: 40})//and & prefix yields a pointer to the struct
	
	fmt.Println(newPerson("Jon"))//it is normal practice to encapsulate new struct creation in constructor functions
	
	s := person{name: "Sean", age: 38}
	fmt.Println(s.name)//access struct fields with a dot

	sp := &s
	fmt.Println(sp.age)//you can also you dot with struct pointers - the pointers are automatically deferenced

	sp.age = 58
	fmt.Println(sp.age)//structs are liable to change
}