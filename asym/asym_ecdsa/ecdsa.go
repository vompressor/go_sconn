package asym_ecdsa

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/x509"
	"encoding/pem"
	"errors"

	"github.com/vompressor/go_sconn/asym"
)

type ASymECDSA struct {
	*ecdsa.PrivateKey
}

type ASymECDSAPub struct {
	*ecdsa.PublicKey
}

func New(c elliptic.Curve) (*ASymECDSA, error) {
	ret, err := ecdsa.GenerateKey(c, asym.RandReader)
	if err != nil {
		return nil, err
	}
	return &ASymECDSA{ret}, nil
}

func Load(privBytes []byte) (*ASymECDSA, error) {
	rsaPriv, err := x509.ParsePKCS8PrivateKey(privBytes)
	if err != nil {
		return nil, err
	}

	return &ASymECDSA{rsaPriv.(*ecdsa.PrivateKey)}, nil
}

func LoadPub(pubBytes []byte) (*ASymECDSAPub, error) {
	pub, err := x509.ParsePKIXPublicKey(pubBytes)

	if err != nil {
		return nil, err
	}
	return &ASymECDSAPub{PublicKey: pub.(*ecdsa.PublicKey)}, nil
}

func LoadPEM(pemBytes []byte) (*ASymECDSA, error) {
	p, _ := pem.Decode(pemBytes)

	if p == nil {
		return nil, errors.New("invalid input pemBytes")
	}
	return Load(p.Bytes)
}

func LoadPubPEM(pemBytes []byte) (*ASymECDSAPub, error) {
	p, _ := pem.Decode(pemBytes)

	if p == nil {
		return nil, errors.New("invalid input pemBytes")
	}
	return LoadPub(p.Bytes)
}

func (a *ASymECDSA) PrivByte() ([]byte, error) {
	data, err := x509.MarshalPKCS8PrivateKey(a.PrivateKey)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (a *ASymECDSA) PubByte() ([]byte, error) {
	return a.PubASym().PubByte()
}

func (a *ASymECDSAPub) PubByte() ([]byte, error) {
	data, err := x509.MarshalPKIXPublicKey(a.PublicKey)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (a *ASymECDSA) PrivPEM() ([]byte, error) {
	data, err := a.PrivByte()
	if err != nil {
		return nil, err
	}

	block := pem.Block{Type: "PRIVATE KEY", Bytes: data}
	return pem.EncodeToMemory(&block), nil
}

func (a *ASymECDSA) PubPEM() ([]byte, error) {
	return a.PubASym().PubPEM()
}

func (a *ASymECDSAPub) PubPEM() ([]byte, error) {
	data, err := a.PubByte()
	if err != nil {
		return nil, err
	}
	block := pem.Block{Type: "PUBLIC KEY", Bytes: data}
	return pem.EncodeToMemory(&block), nil
}

func (a *ASymECDSA) PubASym() *ASymECDSAPub {
	return &ASymECDSAPub{&a.PublicKey}
}

func (a *ASymECDSAPub) Verify(hash, sign []byte) bool {
	return ecdsa.VerifyASN1(a.PublicKey, hash, sign)
}

func (a *ASymECDSA) Sign(hash []byte) ([]byte, error) {
	return a.PrivateKey.Sign(asym.RandReader, hash, nil)
}
func (a *ASymECDSA) Verify(hash, sign []byte) bool {
	return a.PubASym().Verify(hash, sign)
}
