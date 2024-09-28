package main

import (
	"fmt"

	"github.com/tanvir07-cell/go/io/another"
)


func printAge(age *int){

	if(*age>=23){
		panic("You are too old")
	}



	

	*age++;
}


func printSecond(){
	println("Hello, I am from second.go")
	println("From second x = ",x)

	age:=22
	printAge(&age)

	fmt.Printf("My age is %d\n",age)
	fmt.Printf("age address is %v\n",&age)

	another.Another()

	


}