package main

import (
	"crypto"
	"fmt"
	"github.com/nbari/backup.sh/checksum"
	"log"
)

func main() {
	log.Println("Starting...")
	fmt.Printf("MD5 hash: %s\n", checksum.File("/tmp/test.raw", crypto.SHA512))
}
