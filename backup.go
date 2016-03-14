package main

import (
	"fmt"
	"github.com/nbari/backup.sh/checksum"
	"log"
)

func main() {
	log.Println("Starting...")
	hash, err := checksum.File("/tmp/test.raw", "SHA512")
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("MD5 hash: %s\n", hash)
}
