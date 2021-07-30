// Copyright 2021 vompressor. All rights reserved.
// license that can be found in https://github.com/vompressor/go_sconn/blob/master/LICENSE.

package aes_ctr_upgrader

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"net"

	"github.com/vompressor/go_sconn/sconn"
	"github.com/vompressor/go_sconn/sconn/stream"
)

// Upgrade conn to secured conn, use aes ctr mode
// key argument should be, 16, 24, 32 bytes to select AES-128, AES-192, AES-256
func Upgrade(conn net.Conn, key []byte) (sconn.SConn, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	tempIv := sha256.Sum256(key)

	ctr := cipher.NewCTR(block, tempIv[:block.BlockSize()])

	return stream.Upgrade(conn, ctr), nil
}
