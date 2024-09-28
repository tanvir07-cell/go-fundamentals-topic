package main

import (
	"fmt"

	"github.com/tanvir07-cell/struct/model"
)


func main(){
	person:= model.Person{Name:"Tanvir Rifat",Age:23}

	fmt.Println(person.GetName())


	student:= model.Student{Id:1,Varsity:"Daffodil International University",Person:person}

	
	fmt.Println(student.PrintStudent())

	
	interFace:= model.Interface(person)

	interFace.PrintInterface()





}