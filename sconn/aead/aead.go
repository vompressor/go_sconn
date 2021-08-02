package aead

import (
	"crypto/cipher"
	"crypto/sha256"
	"encoding/binary"
	"net"

	"github.com/vompressor/go_sconn/protocol"
)

func makeNonce(seq uint32, init []byte, nonceLen int) []byte {
	hasher := sha256.New()
	var seqBytes [4]byte
	hasher.Write(init)

	binary.BigEndian.PutUint32(seqBytes[:], seq)

	hasher.Write(seqBytes[:])

	tNonce := hasher.Sum(nil)

	if len(tNonce) < nonceLen {
		padding := make([]byte, nonceLen-len(tNonce))
		tNonce = append(tNonce, padding...)
	}
	return tNonce[:nonceLen]
}

type msgProtocol struct {
	Len uint32
	Seq uint32
}

func (m *msgProtocol) SetBodyLen(b int) {
	m.Len = uint32(b)
}

func (m *msgProtocol) GetBodyLen() int {
	return int(m.Len)
}

type AEADSConn struct {
	cip cipher.AEAD
	net.Conn
	initNon []byte
	seq     uint32
}

func Upgrade(conn net.Conn, cip cipher.AEAD, init []byte) *AEADSConn {
	return &AEADSConn{cip: cip, Conn: conn, initNon: init, seq: 0}
}

func (a *AEADSConn) Read(b []byte) (int, error) {
	msgHead := new(msgProtocol)
	_, body, err := protocol.ReadProtocol(a.Conn, msgHead)
	if err != nil {
		return 0, err
	}
	a.seq = msgHead.Seq + 1
	plain, err := a.cip.Open(nil, makeNonce(msgHead.Seq, a.initNon, a.cip.NonceSize()), body, nil)
	if err != nil {
		return 0, err
	}

	n := copy(b, plain)

	return n, err
}

func (a *AEADSConn) Write(b []byte) (int, error) {
	msgHead := new(msgProtocol)

	msgHead.Seq = a.seq

	d := a.cip.Seal(nil, makeNonce(msgHead.Seq, a.initNon, a.cip.NonceSize()), b, nil)
	_, err := protocol.WriteProtocol(a.Conn, msgHead, d)
	if err != nil {
		return 0, err
	}
	return len(b), err
}
