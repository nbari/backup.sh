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
	hash = String("sha1", crypto.SHA1)
	e := "415ab40ae9b7cc4e66d6769cb2c08106e8293b48"
	if hash != e {
		t.Error(e, hash)
	}
}

func TestSHA256tring(t *testing.T) {
	var hash string
	hash = String("sha256", crypto.SHA256)
	e := "5d5b09f6dcb2d53a5fffc60c4ac0d55fabdf556069d6631545f42aa6e3500f2e"
	if hash != e {
		t.Error(e, hash)
	}
}

func TestSHA512tring(t *testing.T) {
	var hash string
	hash = String("sha512", crypto.SHA512)
	e := "1f9720f871674c18e5fecff61d92c1355cd4bfac25699fb7ddfe7717c9669b4d085193982402156122dfaa706885fd64741704649795c65b2a5bdec40347e28a"
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
