package utils

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadFileString(name string) string {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println("File reading error", err)
		panic(err)
	}
	return string(data)
}

func ReadFileSplitLineToArray(name string) []string {
	fileData := ReadFileString(name)
	fileData = strings.ReplaceAll(fileData, "\r\n", "\n")
	return strings.Split(fileData, "\n")
}

func ConvertCommaStringToIntArray(in string) []int {
	values := strings.Split(in, ",")
	returnArray := make([]int, len(values))
	for index, value := range values {
		returnArray[index], _ = strconv.Atoi(value)
	}
	return returnArray
}

func ConvertStringToIntArray(in string) []int {
	values := strings.Split(in, "")
	returnArray := make([]int, len(values))
	for index, value := range values {
		returnArray[index], _ = strconv.Atoi(value)
	}
	return returnArray
}
