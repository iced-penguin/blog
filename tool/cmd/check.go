package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/icedpenguin0504/blog/tool/helper"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use: "check",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := check(); err != nil {
			return err
		}
		return nil
	},
}

const (
	contentDir = "content"
)

func check() error {
	// 全記事ファイル取得
	var filenames []string
	err := filepath.Walk(contentDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			filenames = append(filenames, path)
		}
		return nil
	})
	if err != nil {
		return err
	}

	var warnings []string
	for _, filename := range filenames {
		// draftチェック
		draft, err := isDraft(filename)
		if err != nil {
			return err
		}
		if draft {
			warnings = append(warnings, fmt.Sprintf("warning: %s is draft", filename))
		}
	}

	for _, warning := range warnings {
		fmt.Println(warning)
	}
	return nil
}

func isDraft(filename string) (bool, error) {
	absFilepath, err := filepath.Abs(filename)
	if err != nil {
		return false, fmt.Errorf("failed to get absolute path: %v", err)
	}
	article := helper.NewArticle(absFilepath)
	metadata, err := article.ReadMetadata()
	if err != nil {
		return false, err
	}
	return metadata.Draft, nil
}
