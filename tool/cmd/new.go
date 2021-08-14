package cmd

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/icedpenguin0504/blog/tool/helper"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use: "new",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := createNewFile(); err != nil {
			return err
		}
		return nil
	},
}

func createNewFile() error {
	prompt := helper.NewPrompt()

	baseFilename, category := prompt.Input()
	filename := fmt.Sprintf("posts/%s.md", baseFilename)

	out, err := exec.Command("hugo", "new", filename).CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to create new file: %v", err)
	}
	fmt.Println(string(out))

	absFilepath, err := filepath.Abs("./content/" + filename)
	if err != nil {
		return fmt.Errorf("failed to get absolute file path: %v", err)
	}

	fileEditor := helper.NewFileEditor(absFilepath)

	if err := fileEditor.AddCategory(category); err != nil {
		return fmt.Errorf("failed to add category: %v", err)
	}

	return nil
}
