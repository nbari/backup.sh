package main

import (
	"crypto"
	"fmt"
	"github.com/nbari/backup.sh/checksum"
)

func main() {
	fmt.Printf("MD5 hash: %s\n", checksum.String("sopas", crypto.MD5))
}
