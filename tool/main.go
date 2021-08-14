package main

import (
	"fmt"
	"os"

	"github.com/icedpenguin0504/blog/tool/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v", os.Args[0], err)
		os.Exit(-1)
	}
}
