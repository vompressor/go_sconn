// Copyright 2021 vompressor. All rights reserved.
// license that can be found in https://github.com/vompressor/go_sconn/blob/master/LICENSE.

package stream

import (
	"crypto/cipher"
	"net"

	"github.com/vompressor/go_sconn/sconn"
)

type StreamSConn struct {
	cip cipher.Stream
	net.Conn
}

func Upgrade(c net.Conn, cip cipher.Stream) sconn.SConn {
	return &StreamSConn{Conn: c, cip: cip}
}

func (ssc *StreamSConn) Read(b []byte) (n int, err error) {

	n, err = ssc.Conn.Read(b)
	ssc.cip.XORKeyStream(b[:n], b[:n])
	return
}

func (ssc *StreamSConn) Write(b []byte) (n int, err error) {
	buf := make([]byte, len(b))
	ssc.cip.XORKeyStream(buf, b)
	n, err = ssc.Conn.Write(buf)

	return
}
