package simple_hash

import (
	"bytes"
	"hash"
)

func GetPass(passphrase []byte, hasher hash.Hash) []byte {
	return hasher.Sum(passphrase)
}

func Compare(password []byte, b []byte, hasher hash.Hash) bool {
	p := GetPass(password, hasher)
	return bytes.Equal(p, b)
}
