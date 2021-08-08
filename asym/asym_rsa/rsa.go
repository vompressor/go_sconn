package asym_rsa

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"

	"github.com/vompressor/go_sconn/asym"
)

type ASymRSA struct {
	*rsa.PrivateKey
}

type ASymRSAPub struct {
	*rsa.PublicKey
}

func New(bits int) (*ASymRSA, error) {
	ret, err := rsa.GenerateKey(asym.RandReader, bits)
	if err != nil {
		return nil, err
	}
	return &ASymRSA{ret}, nil
}

func Load(privBytes []byte) (*ASymRSA, error) {
	rsaPriv, err := x509.ParsePKCS8PrivateKey(privBytes)
	if err != nil {
		return nil, err
	}

	return &ASymRSA{rsaPriv.(*rsa.PrivateKey)}, nil
}

func LoadPub(pubBytes []byte) (*ASymRSAPub, error) {
	rsaPub, err := x509.ParsePKIXPublicKey(pubBytes)

	if err != nil {
		return nil, err
	}
	return &ASymRSAPub{PublicKey: rsaPub.(*rsa.PublicKey)}, nil
}

func LoadPEM(pemBytes []byte) (*ASymRSA, error) {
	p, _ := pem.Decode(pemBytes)

	if p == nil {
		return nil, errors.New("invalid input pemBytes")
	}
	return Load(p.Bytes)
}

func LoadPubPEM(pemBytes []byte) (*ASymRSAPub, error) {
	p, _ := pem.Decode(pemBytes)

	if p == nil {
		return nil, errors.New("invalid input pemBytes")
	}
	return LoadPub(p.Bytes)
}

func (a *ASymRSA) PrivByte() ([]byte, error) {
	data, err := x509.MarshalPKCS8PrivateKey(a.PrivateKey)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (a *ASymRSA) PubByte() ([]byte, error) {
	return a.PubASym().PubByte()
}

func (a *ASymRSA) PrivPEM() ([]byte, error) {
	data, err := a.PrivByte()
	if err != nil {
		return nil, err
	}

	block := pem.Block{Type: "PRIVATE KEY", Bytes: data}
	return pem.EncodeToMemory(&block), nil
}

func (a *ASymRSA) PubPEM() ([]byte, error) {
	return a.PubASym().PubPEM()
}

func (a *ASymRSAPub) PubByte() ([]byte, error) {
	data, err := x509.MarshalPKIXPublicKey(a.PublicKey)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (a *ASymRSAPub) PubPEM() ([]byte, error) {
	data, err := a.PubByte()
	if err != nil {
		return nil, err
	}
	block := pem.Block{Type: "PUBLIC KEY", Bytes: data}
	return pem.EncodeToMemory(&block), nil
}

func (a *ASymRSA) PubASym() *ASymRSAPub {
	return &ASymRSAPub{&a.PublicKey}
}

func (a *ASymRSA) Decrypt(ciphertext []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(asym.RandReader, a.PrivateKey, ciphertext)
}

func (a *ASymRSA) DecryptOAEP(ciphertext []byte) ([]byte, error) {
	return rsa.DecryptOAEP(sha256.New(), asym.RandReader, a.PrivateKey, ciphertext, nil)
}

func (a *ASymRSA) Encrypt(plaintext []byte) ([]byte, error) {
	return a.PubASym().Encrypt(plaintext)
}

func (a *ASymRSAPub) Encrypt(plaintext []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(asym.RandReader, a.PublicKey, plaintext)
}

func (a *ASymRSAPub) EncryptOAEP(plaintext []byte) ([]byte, error) {
	return rsa.EncryptOAEP(sha256.New(), asym.RandReader, a.PublicKey, plaintext, nil)
}

func (a *ASymRSA) EncryptOAEM(plaintext []byte) ([]byte, error) {
	return a.PubASym().EncryptOAEP(plaintext)
}

func (a *ASymRSAPub) Verify(hash, sign []byte, hasher crypto.Hash) bool {
	if err := rsa.VerifyPKCS1v15(a.PublicKey, hasher, hash, sign); err != nil {
		return false
	}
	return true
}

func (a *ASymRSA) Sign(hash []byte, hasher crypto.Hash) ([]byte, error) {
	return rsa.SignPKCS1v15(asym.RandReader, a.PrivateKey, hasher, hash)
}
func (a *ASymRSA) Verify(hash, sign []byte, hasher crypto.Hash) bool {
	return a.PubASym().Verify(hash, sign, hasher)
}
