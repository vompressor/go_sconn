// Copyright 2021 vompressor. All rights reserved.
// license that can be found in https://github.com/vompressor/go_sconn/blob/master/LICENSE.

package sconn

import (
	"net"
)

type SConn interface {
	net.Conn
	Read([]byte) (int, error)
	Write([]byte) (int, error)
}

type ConnUpgrader func(net.Conn, []byte) (SConn, error)
