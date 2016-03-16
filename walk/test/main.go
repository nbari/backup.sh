package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func visit(path string, f os.FileInfo, err error) error {
	parts := strings.Split(path, string(filepath.Separator))
	if f.IsDir() {
		fmt.Printf("Path: %s %s\n", path, parts)
	} else {
		fmt.Printf("File: %s Path: %s\n", f.Name(), path)
	}
	return nil
}

func main() {
	flag.Parse()
	root := flag.Arg(0)
	err := filepath.Walk(root, visit)
	fmt.Printf("filepath.Walk() returned %v\n", err)
}
