package main

import (
	"crypto/sha512"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	//	"strings"
	"sync"
)

type result struct {
	path string
	sum  [sha512.Size]byte
	err  error
}

var files = make(map[[sha512.Size]byte]string)

// splitPath returns an slice of the path
func splitPath(p string) []string {
	regex := fmt.Sprintf("[^%c ]+", filepath.Separator)
	split_path_rx := regexp.MustCompile(regex)
	return split_path_rx.FindAllString(p, -1)
}

func walkFilesX(root string) filepath.WalkFunc {
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

func walkFiles(done <-chan struct{}, root string) (<-chan string, <-chan string, <-chan error) {
	paths := make(chan string)
	errc := make(chan error, 1) //buffered
	go func() {
		defer close(paths)
		errc <- filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.Mode().IsRegular() {
				files <- path[len(root):]
			}

			if info.Mode().IsDir() {
				paths <- path[len(root):]
			}
			select {
			case <-done:
				return errors.New("walk canceled")
			}
			return nil
		})
	}()
	return paths, errc
}

func buildTree(root string) (map[string]string, error) {

	done := make(chan struct{})
	defer close(done)

	paths, files, errc := walkFiles(done, root)
	print(errc)

	var wg sync.WaitGroup

	wg.Add(1)

	for path := range paths {
		fmt.Printf("Dir: %v\n", path)
	}

	for file := range files {
		fmt.Printf("File: %v\n", file)
	}

	return nil, nil
}

func main() {
	flag.Parse()
	root, err := filepath.Abs(flag.Arg(0))
	if err != nil {
		panic(err)
	}

	tree, err := buildTree(root)

	if err != nil {
		panic(err)
	}

	print(tree)
}
