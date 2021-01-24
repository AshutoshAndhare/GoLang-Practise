package main

import "fmt"

//Employee is just the name of the struct which will have the name 
type Employee struct {
	Name string
}
//UpdateName is a method where syntax goes like func(pointer *Pointer) myMethod()
func (e *Employee) UpdateName(newName string) {
	e.Name = newName
}

//PrintName just prints the name
func (e *Employee) PrintName() {
	fmt.Println(e.Name)
}

func main() {
	var employee Employee
	employee.Name = "Elliot"
	employee.UpdateName("Forbsey")
	employee.PrintName()
}