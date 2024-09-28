package write

import "os"


func WriteFile(fileName string,data string) error{
	err:= os.WriteFile(fileName,[]byte(data),0666)

	if err !=nil{
		return err
	}

	return nil
}