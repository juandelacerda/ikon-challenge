package file

import (
	"fmt"
	"os"
	"strings"
)


func WriteLines(path string, fileName string, lines []string) error{
	f, err := os.Create(path+"/"+fileName)
	if err !=nil {
		return err
	}
    checkError(err)

	defer f.Close()

	for _, line := range lines {
		_, err := f.WriteString(line+"\n")
		 checkError(err)
		 if err !=nil {
			 return err
		 }
	}

	return nil
}

func ReadLines(path string, fileName string) (error, []string) {
	fmt.Println(path+"/"+fileName)
	dat, err := os.ReadFile(path+"/"+fileName)
    checkError(err)
	if err !=nil {
		return err, []string{}
	}

	lines:= strings.Split(string(dat),"\n")

	return  nil, lines
}


func checkError(err error) {
	if err != nil {
		fmt.Println("Ocurrió un error %s", err)
    }
}