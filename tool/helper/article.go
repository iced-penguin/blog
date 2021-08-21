package helper

import (
	"fmt"
	"os"
	"strings"
)

type Article struct {
	filename string
}

func NewArticle(absFilepath string) Article {
	return Article{filename: absFilepath}
}

func (f Article) AddCategory(category string) error {
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
