package main

import (
	"fmt"
	"os"

	read "github.com/tanvir07-cell/files/readFile"
)

func main(){
	// get the directory path
	path,_:= os.Getwd()

	content,err:= read.ReadFile(path+"/data/data.txt")

	

	if err !=nil{
		fmt.Printf("Error: %v\n",err)
	}

	fmt.Println(content)




}