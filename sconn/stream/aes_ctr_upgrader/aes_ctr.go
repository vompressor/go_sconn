package aes_ctr_upgrader

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"net"

	"github.com/vompressor/go_sconn/sconn"
	"github.com/vompressor/go_sconn/sconn/stream"
)

func Upgrade(conn net.Conn, key []byte) (sconn.SConn, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	tempIv := sha256.Sum256(key)

	ctr := cipher.NewCTR(block, tempIv[:block.BlockSize()])

	return stream.Upgrade(conn, ctr), nil
}
