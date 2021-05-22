package tool

import (
	"fmt"
)

var categories = map[string]string{
	"1": "Programming",
	"2": "DB",
	"3": "PC",
}

type Prompt struct{}

func NewPrompt() Prompt {
	return Prompt{}
}

func (p Prompt) Input() (filename, category string) {
	filename = inputFilename()
	category = inputCategory()
	return
}

func inputFilename() string {
	return inputEach("Enter file name (without extension)")
}

func inputCategory() string {
	key := inputEach("Choose category (1: Programming, 2: DB, 3: PC)")
	category, ok := categories[key]
	if !ok {
		return inputCategory()
	}
	return category
}

func inputEach(msg string) (in string) {
	fmt.Println(msg)
	fmt.Print(">>> ")
	fmt.Scan(&in)
	return
}
