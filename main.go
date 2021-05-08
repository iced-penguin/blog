package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/icedpenguin0504/blog/tool"
)

func main() {
	if err := createNewFile(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func createNewFile() error {
	baseFilename, category := tool.Prompt()
	filename := fmt.Sprintf("posts/%s.md", baseFilename)
	out, _ := exec.Command("hugo", "new", filename).CombinedOutput()
	fmt.Println(string(out))

	if err := overwriteFile(baseFilename, category); err != nil {
		return err
	}
	return nil
}

func overwriteFile(baseFilename, category string) error {
	filename := fmt.Sprintf("content/posts/%s.md", baseFilename)
	b, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	oldLine := "categories: []"
	newLine := fmt.Sprintf("categories: [%s]", category)
	s := strings.Replace(string(b), oldLine, newLine, 1)

	os.WriteFile(filename, []byte(s), 0644)

	return nil
}
