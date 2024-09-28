package model

import "fmt"


type Person struct{
	Name string
	Age int
}


func (p Person) GetName() string{
	return fmt.Sprintf("Name: %s", p.Name)
}


// interface(kind of polymorphism in OOP) implementation
func (p Person) PrintInterface() {
	 fmt.Println("I am a person")
}