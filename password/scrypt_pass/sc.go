package scryptpass

import (
	"bytes"

	"github.com/vompressor/go_sconn/password"
	"golang.org/x/crypto/scrypt"
)

type ScryptOption struct {
	Salt            []byte
	KeyLen, N, R, P int
}

func GetPass(passphrase []byte, option *ScryptOption) ([]byte, []byte, error) {
	var len, n, r, p int = 32, 32768, 8, 1
	var salt []byte
	if option == nil {
		salt = make([]byte, 8)
		password.RandReader.Read(salt)
	} else {
		salt = option.Salt
		len = option.KeyLen
		n = option.N
		r = option.R
		p = option.P
	}
	k, err := scrypt.Key(passphrase, salt, n, r, p, len)
	if err != nil {
		return nil, nil, err
	}

	return k, salt, nil
}

func Compare(password []byte, b []byte, option *ScryptOption) (bool, error) {
	p, _, err := GetPass(password, option)
	if err != nil {
		return false, err
	}
	return bytes.Equal(p, b), nil
}
