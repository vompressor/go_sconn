package key_exchanger

import (
	"errors"
	"io"

	"github.com/vompressor/go_sconn/protocol"
)

// T - type 2byte
// M - Method 2byte
// Seq - sequence number
// Len - len

const header_len = 12

const type_handshake uint16 = 0x08

const method_client_hello uint16 = 0x01
const method_server_hello uint16 = 0x02

// const method_client_chk uint16 = 0x03
// const method_client_ok uint16 = 0x04
// const method_refresh uint16 = 0x03

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
