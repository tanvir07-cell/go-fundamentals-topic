package main

import "fmt"

const x int = 10

func init(){
	fmt.Println("Hehe i run first")
}

func main(){
	println("Hello, World!")
	println("x = ",x)
	printSecond()
}