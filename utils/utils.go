package utils

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ReadFile(name string) string {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println("File reading error", err)
		panic(err)
	}

	return string(data)
}

func ReadFileSplitToArray(name string) []string {
	fileData := ReadFile(name)
	fileData = strings.ReplaceAll(fileData, "\r\n", "\n")
	return strings.Split(fileData, "\n")
}
