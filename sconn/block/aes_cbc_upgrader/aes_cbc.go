package aes_cbc_upgrader

import (
	"crypto/aes"
	"crypto/cipher"
	"net"

	"github.com/vompressor/go_sconn/sconn"
	"github.com/vompressor/go_sconn/sconn/block"
)

func Upgrade(conn net.Conn, key []byte) (sconn.SConn, error) {
	cip, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return block.Upgrade(conn, cip, cipher.NewCBCEncrypter, cipher.NewCBCDecrypter), nil
}
