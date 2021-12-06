package utils

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadFile(name string) []byte {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println("File reading error", err)
		panic(err)
	}

	return data
}

func ReadFileString(name string) string {
	return string(ReadFile(name))
}

func ReadFileSplitToArray(name string) []string {
	fileData := ReadFileString(name)
	fileData = strings.ReplaceAll(fileData, "\r\n", "\n")
	return strings.Split(fileData, "\n")
}

func ConvertFormattedStringToIntArray(in string) []int {
	values := strings.Split(in, ",")
	returnArray := make([]int, len(values))
	for index, value := range values {
		returnArray[index], _ = strconv.Atoi(value)
	}
	return returnArray
}