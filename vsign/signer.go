package vsign

import (
	"bytes"
	"crypto/cipher"
	"hash"
	"io"
)

type VSigner struct {
	signer cipher.AEAD
	hasher hash.Hash
}

func NewVSigner(signer cipher.AEAD, hasher hash.Hash) *VSigner {
	return &VSigner{signer: signer, hasher: hasher}
}

func (vs *VSigner) Seal(r io.Reader) ([]byte, error) {
	io.Copy(vs.hasher, r)
	vs.hasher.Reset()
	hash := vs.hasher.Sum(nil)
	nonce := make([]byte, vs.signer.NonceSize())
	return vs.signer.Seal(nil, nonce, hash, nil), nil
}

func (vs *VSigner) Open(r io.Reader, signature []byte) (bool, error) {
	io.Copy(vs.hasher, r)
	vs.hasher.Reset()
	hash := vs.hasher.Sum(nil)
	nonce := make([]byte, vs.signer.NonceSize())
	pd, err := vs.signer.Open(nil, nonce, signature, nil)
	return bytes.Equal(pd, hash), err
}
