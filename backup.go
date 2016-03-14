package backup

import (
	"github.com/nbari/backup.sh/checksum"
	"log"
	"time"
)

type Backup struct {
	src   string
	dst   string
	debug bool
}

func New() {
	start := time.Now()
	log.Println("Starting...")
	algo := "SHA512"
	hash, err := checksum.File("/tmp/1GB.raw", algo)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%s: %x\n", algo, hash)
	log.Printf("Elapsed time: %s\n", time.Since(start))
}
