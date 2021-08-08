package asym_ed25519

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"errors"

	"github.com/vompressor/go_sconn/asym"
)

type ASymED25519 struct {
	ed25519.PrivateKey
}

type ASymED25519Pub struct {
	ed25519.PublicKey
}

func New() (*ASymED25519, error) {
	_, priv, err := ed25519.GenerateKey(asym.RandReader)

	if err != nil {
		return nil, err
	}
	return &ASymED25519{PrivateKey: priv}, nil
}

func Load(privBytes []byte) (*ASymED25519, error) {
	edPriv, err := x509.ParsePKCS8PrivateKey(privBytes)
	if err != nil {
		return nil, err
	}
	edp := edPriv.(ed25519.PrivateKey)
	return &ASymED25519{PrivateKey: edp}, nil
}

func LoadPub(pubBytes []byte) (*ASymED25519Pub, error) {
	pub, err := x509.ParsePKIXPublicKey(pubBytes)

	if err != nil {
		return nil, err
	}
	return &ASymED25519Pub{PublicKey: pub.(ed25519.PublicKey)}, nil
}

func LoadPEM(pemBytes []byte) (*ASymED25519, error) {
	p, _ := pem.Decode(pemBytes)

	if p == nil {
		return nil, errors.New("invalid input pemBytes")
	}
	return Load(p.Bytes)
}

func LoadPubPEM(pemBytes []byte) (*ASymED25519Pub, error) {
	p, _ := pem.Decode(pemBytes)

	if p == nil {
		return nil, errors.New("invalid input pemBytes")
	}
	return LoadPub(p.Bytes)
}

func (a *ASymED25519) PrivByte() ([]byte, error) {
	data, err := x509.MarshalPKCS8PrivateKey(a.PrivateKey)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (a *ASymED25519) PubByte() ([]byte, error) {
	return a.PubASym().PubByte()
}

func (a *ASymED25519) PrivPEM() ([]byte, error) {
	data, err := a.PrivByte()
	if err != nil {
		return nil, err
	}

	block := pem.Block{Type: "PRIVATE KEY", Bytes: data}
	return pem.EncodeToMemory(&block), nil
}

func (a *ASymED25519) PubPEM() ([]byte, error) {
	return a.PubASym().PubPEM()
}
func (a *ASymED25519Pub) PubByte() ([]byte, error) {
	data, err := x509.MarshalPKIXPublicKey(a.PublicKey)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (a *ASymED25519Pub) PubPEM() ([]byte, error) {
	data, err := a.PubByte()
	if err != nil {
		return nil, err
	}
	block := pem.Block{Type: "PUBLIC KEY", Bytes: data}
	return pem.EncodeToMemory(&block), nil
}

func (a *ASymED25519) PubASym() *ASymED25519Pub {
	pub := a.Public()
	ret := pub.(ed25519.PublicKey)
	return &ASymED25519Pub{ret}
}

func (a *ASymED25519Pub) Verify(msg, sign []byte) bool {
	return ed25519.Verify(a.PublicKey, msg, sign)
}

func (a *ASymED25519) Sign(msg []byte) ([]byte, error) {
	return ed25519.Sign(a.PrivateKey, msg), nil
}
func (a *ASymED25519) Verify(hash, sign []byte) bool {
	return a.PubASym().Verify(hash, sign)
}
