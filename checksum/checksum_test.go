package checksum

import (
	"crypto"
	"encoding/hex"
	"testing"
)

func TestMD5String(t *testing.T) {
	var hash string
	hash = hex.EncodeToString(String("md5", crypto.MD5))
	e := "1bc29b36f623ba82aaf6724fd3b16718"
	if hash != e {
		t.Error(e, hash)
	}
}

func TestSHA1tring(t *testing.T) {
	var hash string
	hash = hex.EncodeToString(String("The quick brown fox jumps over the lazy dog", crypto.SHA1))
	e := "2fd4e1c67a2d28fced849ee1bb76e7391b93eb12"
	if hash != e {
		t.Error(e, hash)
	}
}

func TestSHA256tring(t *testing.T) {
	var hash string
	hash = hex.EncodeToString(String("The quick brown fox jumps over the lazy dog", crypto.SHA256))
	e := "d7a8fbb307d7809469ca9abcb0082e4f8d5651e46d3cdb762d02d0bf37c9e592"
	if hash != e {
		t.Error(e, hash)
	}
}

func TestSHA512tring(t *testing.T) {
	var hash string
	hash = hex.EncodeToString(String("The quick brown fox jumps over the lazy dog", crypto.SHA512))
	e := "07e547d9586f6a73f73fbac0435ed76951218fb7d0c8d788a309d785436bbb642e93a252a954f23912547d1e8a3b5ed6e1bfd7097821233fa0538f3db854fee6"
	if hash != e {
		t.Error(e, hash)
	}
}

func TestFile(t *testing.T) {
	var hash string
	hash = hex.EncodeToString(File("./checksum.go", crypto.SHA512))
	e := "891b8e793ce6c2a9cfa732856d0fc7dfb1db7b1497c30f0577ce732143aa2be4c0af05fc312a7da69f52038222ea59bd270bbf9380ae98be2e7a4482b9ff4d87"
	if hash != e {
		t.Error(e, hash)
	}
}
