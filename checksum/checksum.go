/**
* checksum
*
* suported hashes MD5, SHA1, SHA256, SHA512
* return []byte
* encode/hex - hex.EncodeToString
 */

package checksum

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"hash"
	"io"
	"os"
)

func createHash(algo string) (hash.Hash, error) {
	var h hash.Hash
	var err error

	switch algo {
	case "MD5":
		h = md5.New()
	case "SHA1":
		h = sha1.New()
	case "SHA256":
		h = sha256.New()
	case "SHA384":
		h = sha512.New384()
	case "SHA512":
		h = sha512.New()
	default:
		err = errors.New("Method not supported")
	}

	return h, err
}

func String(s, algo string) ([]byte, error) {
	h, err := createHash(algo)
	if err != nil {
		return nil, err
	}
	io.WriteString(h, s)
	return h.Sum(nil), nil
}

func File(file, algo string) ([]byte, error) {
	fh, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer fh.Close()

	h, err := createHash(algo)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(h, fh)

	if err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}
