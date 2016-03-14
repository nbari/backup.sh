package main

import (
	"fmt"
	"github.com/nbari/backup.sh/checksum"
	"log"
	"time"
)

func main() {
	start := time.Now()
	log.Println("Starting...")
	algo := "SHA512"
	hash, err := checksum.File("/tmp/1GB.raw", algo)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%s: %x\n", algo, hash)
	fmt.Printf("Elapsed time: %s\n", time.Since(start))
}
