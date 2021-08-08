package pbkdf2_pass

import (
	"bytes"
	"crypto/sha1"
	"crypto/sha256"
	"hash"

	"github.com/vompressor/go_sconn/password"
	"golang.org/x/crypto/pbkdf2"
)

type PBKDF2Option struct {
	Salt      []byte
	KeyLen    int
	Hasher    func() hash.Hash
	Iteration int
}

func PBKDF2_SHA1_RANDSALT8_LEN32_ITER4096() *PBKDF2Option {
	salt := make([]byte, 8)
	password.RandReader.Read(salt)
	return &PBKDF2Option{
		Salt:      salt,
		KeyLen:    32,
		Hasher:    sha1.New,
		Iteration: 4096,
	}
}

func PBKDF2_SHA256_RANDSALT8_LEN32_ITER4096() *PBKDF2Option {
	salt := make([]byte, 8)
	password.RandReader.Read(salt)
	return &PBKDF2Option{
		Salt:      salt,
		KeyLen:    32,
		Hasher:    sha256.New,
		Iteration: 4096,
	}
}

func GetPass(passphrase []byte, option *PBKDF2Option) ([]byte, []byte, error) {
	var salt []byte
	var keyLen, iter int = 32, 4096
	var hasher func() hash.Hash = sha1.New

	if option == nil {
		salt = make([]byte, 8)
		password.RandReader.Read(salt)
	} else {
		salt = option.Salt
		keyLen = option.KeyLen
		iter = option.Iteration
		hasher = option.Hasher
	}

	return pbkdf2.Key(passphrase, salt, iter, keyLen, hasher), salt, nil
}

func Compare(password []byte, b []byte, option *PBKDF2Option) (bool, error) {
	p, _, err := GetPass(password, option)
	if err != nil {
		return false, err
	}

	return bytes.Equal(p, b), nil
}
