package helper

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

const (
	metaSeparator = "---"
)

type Article struct {
	filename string
}

func NewArticle(absFilepath string) Article {
	return Article{filename: absFilepath}
}

func (a Article) AddCategory(category string) error {
	buf, err := os.ReadFile(a.filename)
	if err != nil {
		return fmt.Errorf("cannot read file: %v", err)
	}

	oldLine := "categories: []"
	newLine := fmt.Sprintf("categories: [%s]", category)
	s := strings.Replace(string(buf), oldLine, newLine, 1)

	os.WriteFile(a.filename, []byte(s), 0644)

	return nil
}

func (a Article) ReadMetadata() (Metadata, error) {
	f, err := os.Open(a.filename)
	if err != nil {
		return Metadata{}, fmt.Errorf("cannot open file: %v", err)
	}
	defer f.Close()

	var frontMatter []string
	sepCount := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == metaSeparator {
			if sepCount > 0 {
				break
			}
			sepCount += 1
			continue
		}
		frontMatter = append(frontMatter, line)
	}

	buf := []byte(strings.Join(frontMatter, "\n"))
	metadata := Metadata{}
	err = yaml.Unmarshal(buf, &metadata)
	if err != nil {
		return Metadata{}, err
	}
	return metadata, nil
}

type Metadata struct {
	Title      string   `yaml:"title"`
	Date       string   `yaml:"date"`
	Categories []string `yaml:"categories"`
	Tags       []string `yaml:"tags"`
	TOC        bool     `yaml:"toc"`
	Draft      bool     `yaml:"draft"`
}
