package tool

import (
	"fmt"
)

var categories = map[string]string{
	"1": "Programming",
	"2": "DB",
	"3": "PC",
}

func Prompt() (filename, category string) {
	filename = inputFilename()
	category = inputCategory()
	return
}

func inputFilename() string {
	return input("Enter file name (without extension)")
}

func inputCategory() string {
	key := input("Choose category (1: Programming, 2: DB, 3: PC)")
	category, ok := categories[key]
	if !ok {
		return inputCategory()
	}
	return category
}

func input(msg string) (in string) {
	fmt.Println(msg)
	fmt.Print(">>> ")
	fmt.Scan(&in)
	return
}
