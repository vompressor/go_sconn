package chacha20poly1305_upgrader

import (
	"crypto/sha256"
	"net"

	"github.com/vompressor/go_sconn/sconn"
	"github.com/vompressor/go_sconn/sconn/aead"
	"golang.org/x/crypto/chacha20poly1305"
)

func Upgrade(conn net.Conn, key []byte) (sconn.SConn, error) {
	cip, err := chacha20poly1305.New(key)
	if err != nil {
		return nil, err
	}

	return aead.Upgrade(conn, cip, sha256.New().Sum(key)), nil
}
