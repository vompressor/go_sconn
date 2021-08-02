package aes_gcm_upgrader

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"net"

	"github.com/vompressor/go_sconn/sconn"
	"github.com/vompressor/go_sconn/sconn/aead"
)

func Upgrade(conn net.Conn, key []byte) (sconn.SConn, error) {
	aes_block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	cip, err := cipher.NewGCM(aes_block)
	if err != nil {
		return nil, err
	}

	return aead.Upgrade(conn, cip, sha256.New().Sum(key)), nil
}
