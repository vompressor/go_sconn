// Copyright 2021 vompressor. All rights reserved.
// license that can be found in https://github.com/vompressor/go_sconn/blob/master/LICENSE.

// Package key_exchanger is based net package, key exchanged through net.Conn socket communication.
package key_exchanger

import (
	"errors"
	"io"

	"github.com/vompressor/go_sconn/protocol"
)

const type_handshake uint16 = 0x08

const method_client_hello uint16 = 0x01
const method_server_hello uint16 = 0x02

func readMsg(reader io.Reader, t uint16, m uint16) ([]byte, error) {

	proto := &protocol.BasicProtocol{}

	_, msg, err := protocol.ReadProtocol(reader, proto)

	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	if proto.Method != m {
		return nil, errors.New("method mismatch")
	}

	if proto.Type != t {
		return nil, errors.New("type mismatch")
	}

	return msg, nil
}

func writeMsg(writer io.Writer, t uint16, m uint16, msg []byte) error {
	proto := &protocol.BasicProtocol{}
	proto.Method = m
	proto.Type = t
	proto.Seq = 0
	proto.SetBodyLen(len(msg))
	_, err := protocol.WriteProtocol(writer, proto, msg)
	return err
}
