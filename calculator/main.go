package main

import "fmt"

func main(){
	fmt.Println("Calculator Project 101")
	fmt.Println("===================")
	fmt.Println("Enter Your operation(+, -, *, /): ")

	var operation string
	var firstNumber,secondNumber int

	fmt.Scanf("%s", &operation)

	fmt.Println("Enter Your first number: ")
	fmt.Scanf("%d", &firstNumber)

	fmt.Println("Enter Your second number: ")
	fmt.Scanf("%d", &secondNumber)

	switch operation{
	case "+":
		fmt.Printf("Result: %d\n", firstNumber + secondNumber)
	case "-":
		fmt.Printf("Result: %d\n", firstNumber - secondNumber)
	case "*":
		fmt.Printf("Result: %d\n", firstNumber * secondNumber)
	case "/":
		fmt.Printf("Result: %d\n", firstNumber / secondNumber)
	}


}