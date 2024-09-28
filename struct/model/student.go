package model

import "fmt"


type Student struct{
	Id int
	Varsity string
	Person
}


func (s Student) PrintStudent() string{
	return fmt.Sprintf("Id: %d, Varsity: %s, Name: %s, Age: %d", s.Id, s.Varsity, s.Person.Name, s.Person.Age)
}



// interface(kind of polymorphism in OOP) implementation

func (s Student) PrintInterface() {
	 fmt.Println("I am a student")
}