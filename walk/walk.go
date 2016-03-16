package main

import (
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"github.com/nbari/backup.sh/checksum"
	"os"
	"path/filepath"
	"strings"
)

func walkFiles(done <-chan struct{}, root string) (<-chan string, <-chan error) {
	paths := make(chan string)
	errc := make(chan error, 1)
	go func() {
		// Close the paths channel after Walk returns.
		defer close(paths)
		errc <- filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			select {
			case paths <- path:
			case <-done:
				return errors.New("walk canceled")
			}
			return nil
		})
	}()
	return paths, errc
}

func main() {
	flag.Parse()

	root, err := filepath.Abs(flag.Arg(0))

	if err != nil {
		panic(err.Error())
	}

	directories := make(map[string]bool)
	files := make(map[string]string)

	file_ch := make(chan []byte)

	fmt.Printf("root: %s\n", root)

	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		segments := strings.Split(path, string(filepath.Separator))
		if info.IsDir() {
			fmt.Printf("Dir: %s segments: %s\n", path, segments)
			directories[path] = false
		} else {
			//			files[path] = path
			go func() {
				fmt.Printf("File: %s, size: %d bytes\n", path, info.Size())
				cs, err := checksum.File(path, "SHA512")
				if err != nil {
					fmt.Println(err)
				}
				file_ch <- cs
			}()
		}
		return nil
	})

	if err != nil {
		panic(err.Error())
	}

	for i := range files {
		fmt.Println(i)
	}

	for i := range directories {
		fmt.Println(i)
	}

	fmt.Println("-------")
	//	for i := range file_ch {
	for {
		select {
		case i := <-file_ch:
			fmt.Println(base64.URLEncoding.EncodeToString(i))
		}
	}

}
