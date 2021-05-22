package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/icedpenguin0504/blog/tool"
)

func main() {
	if err := createNewFile(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func createNewFile() error {
	prompt := tool.NewPrompt()

	baseFilename, category := prompt.Input()
	filename := fmt.Sprintf("posts/%s.md", baseFilename)

	out, err := exec.Command("hugo", "new", filename).CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to create new file: %v", err)
	}
	fmt.Println(string(out))

	absFilepath, err := filepath.Abs("./conten/" + filename)
	if err != nil {
		return fmt.Errorf("failed to get absolute file path: %v", err)
	}

	fileEditor := tool.NewFileEditor(absFilepath)

	if err := fileEditor.AddCategory(category); err != nil {
		return fmt.Errorf("failed to add category: %v", err)
	}

	return nil
}
