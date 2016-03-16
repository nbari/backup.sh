package main

import (
	"crypto/sha512"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	//	"strings"
)

var files = make(map[[sha512.Size]byte]string)

// splitPath returns an slice of the path
func splitPath(p string) []string {
	regex := fmt.Sprintf("[^%c ]+", filepath.Separator)
	split_path_rx := regexp.MustCompile(regex)
	return split_path_rx.FindAllString(p, -1)
}

func walkFiles(root string) filepath.WalkFunc {
	var current_path string
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Print(err)
			return nil
		}
		current_path = path[len(root):]
		parts := splitPath(current_path)
		if info.IsDir() {
			fmt.Printf("Dir: %s Parts: %v\n", current_path, parts)
		} else {
			fmt.Printf("Path: %s file: %s\n", current_path, info.Name())
		}
		return nil
	}
}

func main() {
	flag.Parse()
	root, err := filepath.Abs(flag.Arg(0))
	err = filepath.Walk(root, walkFiles(root))
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}
