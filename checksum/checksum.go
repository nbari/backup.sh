/**
* checksum
*
* suported hashes MD5, SHA1, SHA256, SHA512
* return []byte
* encode/hex - hex.EncodeToString
 */

package checksum

import (
	"crypto"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"io"
	"os"
)

func createHash(method crypto.Hash) hash.Hash {
	var h hash.Hash

	switch method {
	case crypto.MD5:
		h = md5.New()
	case crypto.SHA1:
		h = sha1.New()
	case crypto.SHA256:
		h = sha256.New()
	case crypto.SHA512:
		h = sha512.New()
	default:
		panic("Method not supported")
	}

	return h
}

func String(s string, method crypto.Hash) []byte {
	h := createHash(method)
	io.WriteString(h, s)
	return h.Sum(nil)
}

func File(file string, method crypto.Hash) []byte {
	fh, err := os.Open(file)

	if err != nil {
		panic(err.Error())
	}

	defer fh.Close()

	h := createHash(method)
	_, err = io.Copy(h, fh)

	if err != nil {
		panic(err.Error())
	}

	return h.Sum(nil)
}
