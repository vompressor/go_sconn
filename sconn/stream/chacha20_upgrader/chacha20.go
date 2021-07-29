package chacha20_upgrader

import (
	"net"

	"github.com/vompressor/go_sconn/sconn"
	"github.com/vompressor/go_sconn/sconn/stream"
	"golang.org/x/crypto/chacha20"
)

func Upgrade(conn net.Conn, key []byte) (sconn.SConn, error) {
	non := make([]byte, 24)
	cip, err := chacha20.NewUnauthenticatedCipher(key, non)
	if err != nil {
		return nil, err
	}

	return stream.Upgrade(conn, cip), nil
}
