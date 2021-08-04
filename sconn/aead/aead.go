package aead

import (
	"bytes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/binary"
	"errors"
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
	Len               uint32
	Seq               uint32
	AdditionalDataLen uint32
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
	buf     *bytes.Buffer
}

func Upgrade(conn net.Conn, cip cipher.AEAD, init []byte) *AEADSConn {
	return &AEADSConn{cip: cip, Conn: conn, initNon: init, seq: 0, buf: new(bytes.Buffer)}
}

func (a *AEADSConn) Read(b []byte) (int, error) {
	var bufRead int
	var err error
	if a.buf.Len() > 0 {
		bufRead, err = a.buf.Read(b)
		return bufRead, err
	}

	x, _, _, err := a.ReadAndOpen()

	if err != nil {
		return 0, err
	}

	n, err := a.buf.Write(x)
	if err != nil {
		return bufRead, err
	}
	if n != len(x) {
		return bufRead, errors.New("size mismatch")
	}
	return a.buf.Read(b)
}

func (a *AEADSConn) Write(b []byte) (int, error) {
	return a.SealAndWrite(b, nil)
}

func (a *AEADSConn) SealAndWrite(plaintext []byte, additionaldata []byte) (int, error) {
	buf := new(bytes.Buffer)
	msgHead := new(msgProtocol)

	msgHead.Seq = a.seq
	var h []byte
	if additionaldata == nil {
		h = nil
	} else {
		msgHead.AdditionalDataLen = uint32(len(additionaldata))
		h = sha256.New().Sum(additionaldata)
	}
	sealed := a.cip.Seal(nil, makeNonce(msgHead.Seq, a.initNon, a.cip.NonceSize()), plaintext, h)
	protocol.WriteProtocol(buf, msgHead, sealed)

	if additionaldata != nil {
		buf.Write(additionaldata)
	}
	_, err := buf.WriteTo(a.Conn)
	a.seq += 1
	return len(plaintext) + int(msgHead.AdditionalDataLen), err
}

func (a *AEADSConn) ReadAndOpen() ([]byte, []byte, int, error) {
	msgHead := new(msgProtocol)

	_, data, err := protocol.ReadProtocol(a.Conn, msgHead)

	if err != nil {
		return nil, nil, 0, err
	}

	var additional, h []byte
	var n int
	if msgHead.AdditionalDataLen > 0 {
		additional = make([]byte, msgHead.AdditionalDataLen)
		n, err = a.Conn.Read(additional)
		if err != nil {
			return nil, nil, 0, err
		}
		if n != int(msgHead.AdditionalDataLen) {
			return nil, nil, 0, errors.New("additonal data read err")
		}
		h = sha256.New().Sum(additional)
	} else {
		h = nil
	}

	plain, err := a.cip.Open(nil, makeNonce(msgHead.Seq, a.initNon, a.cip.NonceSize()), data, h)
	if err != nil {
		return nil, nil, 0, err
	}

	return plain, additional, len(plain) + n, nil
}
