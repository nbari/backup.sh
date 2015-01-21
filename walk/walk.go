package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
}

type Folder struct {
	Name    string
	Files   []File
	Folders map[string]*Folder


func newFolder(name string) *Folder {
	return &Folder{
		name,
		[]File{},
		make(map[string]*Folder)}
}

func (f *Folder) getFolder(name string) *Folder {
	if nextF, ok := f.Folders[name]; ok {
		return nextF
	} else {
		log.Fatalf("Expected nested folder %v in %v\n", name, f.Name)
	}
	return &Folder{} // cannot happen
}

func (f *Folder) addFolder(path []string) {
	for i, segment := range path {
		if i == len(path)-1 { // last segment == new folder
			f.Folders[segment] = newFolder(segment)
		} else {
			f.getFolder(segment).addFolder(path[1:])
		}
	}
}

func (f *Folder) addFile(path []string) {
	for i, segment := range path {
		if i == len(path)-1 { // last segment == file
			f.Files = append(f.Files, File{segment})
		} else {
			f.getFolder(segment).addFile(path[1:])
			return
		}
	}
}

func (f *Folder) String() string {
	var str string
	for _, file := range f.Files {
		str += f.Name + string(filepath.Separator) + file.Name + "\n"
	}
	for _, folder := range f.Folders {
		str += folder.String()
	}
	return str
}

func main() {
	flag.Parse()

	root, err := filepath.Abs(flag.Arg(0))

	if err != nil {
		panic(err.Error())
	}

	rootFolder := newFolder(root)

	fmt.Printf("%v", rootFolder)
	return

	walkFn := func(path string, info os.FileInfo, err error) error {
		segments := strings.Split(path, string(filepath.Separator))
		if info.IsDir() {
			fmt.Printf("Dir: %s\n", path)
			rootFolder.addFolder(segments)
		} else {
			fmt.Printf("File: %s, size: %d bytes\n", path, info.Size())
			rootFolder.addFile(segments)
		}
		return nil
	}

	fmt.Printf("root: %s\n", root)

	err = filepath.Walk(root, walkFn)
	if err != nil {
		panic(err.Error())
	}

	log.Printf("%v\n", rootFolder)
}
