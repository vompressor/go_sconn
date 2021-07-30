// Copyright 2021 vompressor. All rights reserved.
// license that can be found in https://github.com/vompressor/go_sconn/blob/master/LICENSE.

package chacha20_upgrader

import (
	"net"

	"github.com/vompressor/go_sconn/sconn"
	"github.com/vompressor/go_sconn/sconn/stream"
	"golang.org/x/crypto/chacha20"
)

// Upgrade conn to secured conn, use chacha20
// key argument should be, 32 bytes
func Upgrade(conn net.Conn, key []byte) (sconn.SConn, error) {
	non := make([]byte, 24)
	cip, err := chacha20.NewUnauthenticatedCipher(key, non)
	if err != nil {
		return nil, err
	}

	return stream.Upgrade(conn, cip), nil
}
