package main

import "fmt"

type Employee struct {
	Name string
}
//syntax of the method below goes like func(pointer *Pointer) myMethod()
func (e *Employee) UpdateName(newName string) {
	e.Name = newName
}

func (e *Employee) PrintName() {
	fmt.Println(e.Name)
}

func main() {
	var employee Employee
	employee.Name = "Elliot"
	employee.UpdateName("Forbsey")
	employee.PrintName()
}