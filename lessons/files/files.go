package files

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Lesson struct {
}

func (l Lesson) Run() {
	fileNameToRead := "lessons/files/fileToRead"
	data, error := ioutil.ReadFile(fileNameToRead)
	if error == nil {
		fmt.Println(string(data))
	} else {
		fmt.Println("Cannot read file: " + fileNameToRead)
		fmt.Println(error)
	}
	fileNameToWrite := "lessons/files/fileToWrite"
	e := ioutil.WriteFile(fileNameToWrite, data, 0777)
	if e == nil {
		fmt.Println("File has been written")
	} else {
		fmt.Println(e)
	}
	data, error = ioutil.ReadFile(fileNameToWrite)
	if error == nil {
		fmt.Println(string(data))
	} else {
		fmt.Println("Cannot read file: " + fileNameToWrite)
		fmt.Println(error)
	}
	os.Remove(fileNameToWrite)
	fmt.Println("File has been removed")
}

func New() Lesson {
	return *new(Lesson)
}
