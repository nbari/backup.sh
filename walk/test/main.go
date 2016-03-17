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
	//	"sync"
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

func sumFiles(done <-chan struct{}, root string) (<-chan result, <-chan error) {
	c := make(chan result)
	errc := make(chan error, 1)
	go func() { // HL
		var wg sync.WaitGroup
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			wg.Add(1)
			go func() { // HL
				data, err := ioutil.ReadFile(path)
				select {
				case c <- result{path, md5.Sum(data), err}: // HL
				case <-done: // HL
				}
				wg.Done()
			}()
			// Abort the walk if done is closed.
			select {
			case <-done: // HL
				return errors.New("walk canceled")
			default:
				return nil
			}
		})
		// Walk has returned, so all calls to wg.Add are done.  Start a
		// goroutine to close c once all the sends are done.
		go func() { // HL
			wg.Wait()
			close(c) // HL
		}()
		// No select needed here, since errc is buffered.
		errc <- err // HL
	}()
	return c, errc
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
	if err != nil {
		panic(err)
	}

	m, err := MD5All(root)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = filepath.Walk(root, walkFiles(root))
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}
