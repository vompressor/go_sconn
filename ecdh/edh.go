package ecdh

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
)

type DHPubKey struct {
	Pub ecdsa.PublicKey
}

type KeyExchanger struct {
	priv *ecdsa.PrivateKey
}

func NewKXchn() (kxc *KeyExchanger, err error) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	kxc = &KeyExchanger{priv: priv}
	return
}

func (kx *KeyExchanger) GeneratePubPem() (ped []byte, err error) {

	x509EncodedPub, err := x509.MarshalPKIXPublicKey(&kx.priv.PublicKey)
	if err != nil {
		return
	}
	ped = pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})

	return
}

func (kx *KeyExchanger) GeneratePub() (ped []byte, err error) {
	ped, err = x509.MarshalPKIXPublicKey(&kx.priv.PublicKey)
	return
}

func (kx *KeyExchanger) GenerateSharedKey(remotePub *DHPubKey) []byte {
	a, b := remotePub.Pub.Curve.ScalarMult(remotePub.Pub.X, remotePub.Pub.Y, kx.priv.D.Bytes())

	hasher := sha256.New()

	hasher.Write(a.Bytes())
	hasher.Write(b.Bytes())
	return hasher.Sum(nil)
}

func LoadPemToDHPubKey(pemData []byte) (*DHPubKey, error) {
	blockPub, _ := pem.Decode(pemData)

	x509EncodedPub := blockPub.Bytes
	genericPublicKey, err := x509.ParsePKIXPublicKey(x509EncodedPub)
	if err != nil {
		return nil, err
	}
	publicKey := genericPublicKey.(*ecdsa.PublicKey)
	return &DHPubKey{*publicKey}, nil
}
func LoadDHPubKey(data []byte) (*DHPubKey, error) {
	genericPublicKey, err := x509.ParsePKIXPublicKey(data)
	if err != nil {
		return nil, err
	}
	publicKey := genericPublicKey.(*ecdsa.PublicKey)
	return &DHPubKey{*publicKey}, nil
}
