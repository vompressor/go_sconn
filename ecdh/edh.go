// Copyright 2021 vompressor. All rights reserved.
// license that can be found in https://github.com/vompressor/go_sconn/blob/master/LICENSE.

// Package ecdh is Elliptic-curve Diffie–Hellman, ECDH lib.
package ecdh

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
)

// DHPubKey is ecdsa.PublicKey wrapper
type DHPubKey struct {
	Pub ecdsa.PublicKey
}

// KeyExchanger is ecdsa.PrivateKey wrapper
type KeyExchanger struct {
	priv *ecdsa.PrivateKey
}

// NewKXchn key create KeyExchanger
// it create a unique elliptic curve pub/priv key pair.
func NewKXchn() (kxc *KeyExchanger, err error) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	kxc = &KeyExchanger{priv: priv}
	return
}

// GeneratePubPem return x509 - pem encoded byte
func (kx *KeyExchanger) GeneratePubPem() (ped []byte, err error) {

	x509EncodedPub, err := x509.MarshalPKIXPublicKey(&kx.priv.PublicKey)
	if err != nil {
		return
	}
	ped = pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})

	return
}

// GeneratePub return public key as []byte
func (kx *KeyExchanger) GeneratePub() (ped []byte, err error) {
	ped, err = x509.MarshalPKIXPublicKey(&kx.priv.PublicKey)
	return
}

// GenerateSharedKey return shared key hashed by sha256
// it require public key of the other party
// and it combine with the other party's public key to return a unique private key
func (kx *KeyExchanger) GenerateSharedKey(remotePub *DHPubKey) []byte {
	a, b := remotePub.Pub.Curve.ScalarMult(remotePub.Pub.X, remotePub.Pub.Y, kx.priv.D.Bytes())

	hasher := sha256.New()

	hasher.Write(a.Bytes())
	hasher.Write(b.Bytes())
	return hasher.Sum(nil)
}

// LoadPemToDHPubKey decode and load pem []byte to pub key
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

// LoadPemToDHPubKey load []byte to pub key
func LoadDHPubKey(data []byte) (*DHPubKey, error) {
	genericPublicKey, err := x509.ParsePKIXPublicKey(data)
	if err != nil {
		return nil, err
	}
	publicKey := genericPublicKey.(*ecdsa.PublicKey)
	return &DHPubKey{*publicKey}, nil
}
