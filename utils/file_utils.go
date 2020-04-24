package utils

import (
	"os"
)

//FileWrite func
func FileWrite(name string, content []string) (err error) {

	file, err := os.Create("../" + name)
	defer file.Close()

	for _, s := range content {
		file.Write([]byte(s + "\n"))
	}
	err = file.Sync()

	return
}

//FileDelete func
func FileDelete(fileName string) (err error) {
	err = os.Remove("../" + fileName)
	return
}
