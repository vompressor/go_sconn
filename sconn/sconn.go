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
