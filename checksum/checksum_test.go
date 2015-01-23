package checksum

import (
	"crypto"
	"testing"
)

func TestMD5String(t *testing.T) {
	var hash string
	hash = String("md5", crypto.MD5)
	e := "1bc29b36f623ba82aaf6724fd3b16718"
	if hash != e {
		t.Error(e, hash)
	}
}

func TestSHA1tring(t *testing.T) {
	var hash string
	hash = String("The quick brown fox jumps over the lazy dog", crypto.SHA1)
	e := "2fd4e1c67a2d28fced849ee1bb76e7391b93eb12"
	if hash != e {
		t.Error(e, hash)
	}
}

func TestSHA256tring(t *testing.T) {
	var hash string
	hash = String("The quick brown fox jumps over the lazy dog", crypto.SHA256)
	e := "d7a8fbb307d7809469ca9abcb0082e4f8d5651e46d3cdb762d02d0bf37c9e592"
	if hash != e {
		t.Error(e, hash)
	}
}

func TestSHA512tring(t *testing.T) {
	var hash string
	hash = String("The quick brown fox jumps over the lazy dog", crypto.SHA512)
	e := "07e547d9586f6a73f73fbac0435ed76951218fb7d0c8d788a309d785436bbb642e93a252a954f23912547d1e8a3b5ed6e1bfd7097821233fa0538f3db854fee6"
	if hash != e {
		t.Error(e, hash)
	}
}

func TestFile(t *testing.T) {
	var hash string
	hash = File("./checksum.go", crypto.SHA512)
	e := "f8df78487e3d7465d5cd16cb161c468e4e8a08fd5ad556b8b96b221e0d15fc82d1aeade788aa5e6ceb67a854cc381b114d6fde2251ef29b69c677d36226c62bf"
	if hash != e {
		t.Error(e, hash)
	}
}
