package main

import (
	"fmt"
)

type Hash struct {
	hash string
	path []string
}

type Key struct {
	Path, Country string
}

func add(m map[string]map[string]int, path, country string) {
	mm, ok := m[path]
	if !ok {
		mm = make(map[string]int)
		m[path] = mm
	}
	mm[country]++
}
func main() {

	//dict := make(map[*root]string)
	var hashes []*Hash

	//	path := make(map[string]string)

	fmt.Println(hashes)

	hits := make(map[string]map[string]int)

	add(hits, "/doc/", "au")
	add(hits, "/doc/", "mx")

	fmt.Println(hits)

	hits2 := make(map[Key]int)

	hits2[Key{"/", "vn"}]++
	hits2[Key{"/", "vn"}]++
	hits2[Key{"/ok", "vn"}]++
	hits2[Key{"/", "mx"}]++
	fmt.Println(hits2)
}
