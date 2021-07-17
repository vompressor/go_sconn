package secure_conn

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"net"
	"time"

	"github.com/vompressor/go_sconn/protocol"
)

type SecureConn interface {
	net.Conn
	SetKey([]byte)
	GetKey() []byte
}

type BlockSConn struct {
	net.Conn
	cipher.Block
	Key  []byte
	Type uint16
	buf  bytes.Buffer
}

func NewAesSConn(conn net.Conn, key []byte) (*BlockSConn, error) {
	sc := &BlockSConn{}
	sc.Type = 0x04
	err := sc.SetKey(key)
	if err != nil {
		return nil, err
	}

	sc.Conn = conn
	return sc, nil
}

func (bsc *BlockSConn) SetKey(key []byte) error {
	cip, err := aes.NewCipher(key)
	if err != nil {
		return err
	}
	bsc.Block = cip

	bsc.Key = make([]byte, len(key))
	copy(bsc.Key, key)
	return nil
}
func (bsc *BlockSConn) GetKey() []byte {
	return bsc.Key
}

func (bsc *BlockSConn) Read(b []byte) (n int, err error) {

	if bsc.buf.Len() == 0 {
		h := &SecureConnHeader{}
		_, edata, err := protocol.ReadProtocol(bsc.Conn, h)
		if err != nil {
			return 0, err
		}
		data := decrypt(bsc.Block, edata)

		bsc.buf.Reset()
		bsc.buf.Write(data)
	}
	return bsc.buf.Read(b)
}

// TODO:: write return - enced byte len, or plain byte len?

func (bsc *BlockSConn) Write(b []byte) (n int, err error) {
	edata, err := encrypt(bsc.Block, b)
	if err != nil {
		return 0, err
	}

	h := &SecureConnHeader{}
	h.SetBodyLen(len(edata))
	h.Type = bsc.Type
	wdata, err := protocol.EncodeProtocolByte(h, edata)
	if err != nil {
		return 0, err
	}
	return bsc.Conn.Write(wdata)
}

func (bsc *BlockSConn) Close() error {
	return bsc.Conn.Close()
}

func (bsc *BlockSConn) LocalAddr() net.Addr {
	return bsc.Conn.RemoteAddr()
}

func (bsc *BlockSConn) RemoteAddr() net.Addr {
	return bsc.Conn.RemoteAddr()
}

func (bsc *BlockSConn) SetDeadline(t time.Time) error {
	return bsc.Conn.SetDeadline(t)
}

func (bsc *BlockSConn) SetReadDeadline(t time.Time) error {
	return bsc.Conn.SetReadDeadline(t)
}

func (bsc *BlockSConn) SetWriteDeadline(t time.Time) error {
	return bsc.Conn.SetWriteDeadline(t)
}
