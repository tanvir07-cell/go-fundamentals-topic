package read

import "os"

func ReadFile(fileName string)(string,error){
	data,err:= os.ReadFile(fileName)

	// if there is an error:
	if err !=nil{
		return "",err
	}

	return string(data),nil

}