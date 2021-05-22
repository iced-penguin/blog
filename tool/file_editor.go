package tool

import (
	"fmt"
	"os"
	"strings"
)

type FileEditor struct {
	filename string
}

func NewFileEditor(absFilepath string) FileEditor {
	return FileEditor{filename: absFilepath}
}

func (f FileEditor) AddCategory(category string) error {
	buf, err := os.ReadFile(f.filename)
	if err != nil {
		return fmt.Errorf("cannot read file: %v", err)
	}

	oldLine := "categories: []"
	newLine := fmt.Sprintf("categories: [%s]", category)
	s := strings.Replace(string(buf), oldLine, newLine, 1)

	os.WriteFile(f.filename, []byte(s), 0644)

	return nil
}
