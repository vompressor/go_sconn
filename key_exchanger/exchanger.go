// Copyright 2021 vompressor. All rights reserved.
// license that can be found in https://github.com/vompressor/go_sconn/blob/master/LICENSE.

// Package key_exchanger is based net package, key exchanged through net.Conn socket communication.
package key_exchanger

import (
	"crypto/sha256"
	"io"
	"net"

	"github.com/vompressor/go_sconn/ecdh"
	"github.com/vompressor/go_sconn/sconn"
)

type ExchangeListner struct {
	net.Listener
	upgrader sconn.ConnUpgrader
}

func NewExcListener(l net.Listener, upgrader sconn.ConnUpgrader) *ExchangeListner {
	chg := ExchangeListner{}
	chg.Listener = l
	chg.upgrader = upgrader
	return &chg
}

// Accept network conn, upgrade conn to secured conn
func (el *ExchangeListner) Accept() (sconn.SConn, error) {
	conn, err := el.Listener.Accept()

	if err != nil {
		return nil, err
	}

	return ServerSideUpgrade(conn, el.upgrader)
}

// ServerSideUpgrade Upgrade network conn to secured conn, master side
func ServerSideUpgrade(c net.Conn, upgrader sconn.ConnUpgrader) (sconn.SConn, error) {
	changer, err := ecdh.NewKXchn()
	if err != nil {
		return nil, err
	}

	remote_pub, err := readClientHello(c)
	if err != nil {
		return nil, err
	}
	lpub, err := changer.GeneratePub()
	if err != nil {
		return nil, err
	}

	err = writeServerHello(c, lpub)

	if err != nil {
		return nil, err
	}

	pub, err := ecdh.LoadDHPubKey(remote_pub)
	if err != nil {
		return nil, err
	}
	sharedKey := changer.GenerateSharedKey(pub, sha256.New())

	return upgrader(c, sharedKey)
}

// ExcDial Dial network conn, upgrade conn to secured conn
func ExcDial(network, address string, upgrader sconn.ConnUpgrader) (sconn.SConn, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}

	return Upgrade(conn, upgrader)
}

// Upgrade network conn to secured conn, slave side
func Upgrade(c net.Conn, upgrader sconn.ConnUpgrader) (sconn.SConn, error) {
	changer, err := ecdh.NewKXchn()
	if err != nil {
		return nil, err
	}

	lpub, err := changer.GeneratePub()
	if err != nil {
		return nil, err
	}

	err = writeClientHello(c, lpub)

	if err != nil {
		return nil, err
	}

	rpub, err := readServerHello(c)
	if err != nil {
		return nil, err
	}

	dhrpub, err := ecdh.LoadDHPubKey(rpub)
	if err != nil {
		return nil, err
	}

	sharedKey := changer.GenerateSharedKey(dhrpub, sha256.New())
	return upgrader(c, sharedKey)
}

func readServerHello(r io.Reader) (d []byte, err error) {
	d, err = readMsg(r, type_handshake, method_server_hello)
	return
}
func writeServerHello(w io.Writer, b []byte) (err error) {
	err = writeMsg(w, type_handshake, method_server_hello, b)
	return
}

func readClientHello(r io.Reader) (d []byte, err error) {
	d, err = readMsg(r, type_handshake, method_client_hello)
	return
}
func writeClientHello(w io.Writer, b []byte) (err error) {
	err = writeMsg(w, type_handshake, method_client_hello, b)
	return
}
