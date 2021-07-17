package secure_conn

import (
	"crypto/aes"
	"crypto/cipher"
	"net"
	"time"

	"github.com/vompressor/go_sconn/protocol"
)

type SecureConn interface {
	net.Conn
}

type BlockSConn struct {
	net.Conn
	cipher.Block
	Key  []byte
	Type uint16
}

func NewAesSConn(conn net.Conn, key []byte) (*BlockSConn, error) {
	sc := &BlockSConn{}
	sc.Type = 0x04
	cip, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	sc.Block = cip
	sc.Conn = conn
	return sc, nil
}

func (bsc *BlockSConn) Read(b []byte) (n int, err error) {
	h := &SecureConnHeader{}
	l, edata, err := protocol.ReadProtocol(bsc.Conn, h)
	if err != nil {
		return 0, err
	}
	// TODO:: b size err handleing
	data := decrypt(bsc.Block, edata)
	copy(b, data)
	return l, err
}

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
