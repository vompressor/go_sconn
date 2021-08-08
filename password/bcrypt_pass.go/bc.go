package bcrypt_pass

import (
	"golang.org/x/crypto/bcrypt"
)

func GetPass(passphrase []byte, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword(passphrase, cost)
}

func Compare(password []byte, b []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(b, password)
	if err != nil {
		return false, err
	}
	return true, nil
}
