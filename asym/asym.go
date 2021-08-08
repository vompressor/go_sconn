package asym

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
)

type ASym interface {
	PubByte() ([]byte, error)
	PrivByte() ([]byte, error)
	PubPEM() ([]byte, error)
	PrivPEM() ([]byte, error)
	PubASym() ASymPub
}

type ASymPub interface {
	PubPEM() ([]byte, error)
	PubByte() ([]byte, error)
}

type ASymEncrypter interface {
	Encrypt([]byte) ([]byte, error)
}

type ASymDecrypter interface {
	Decrypt([]byte) ([]byte, error)
}

type ASymSigner interface {
	Sign()
}

type ASymVerifier interface{
	Verify()
}

type AsymWithPassPassphrase interface {
}

var RandReader = rand.Reader

type ASymECDSA struct {
	ecdsa.PrivateKey
}

type ASymRSA struct {
	rsa.PrivateKey
}
