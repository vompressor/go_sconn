package block

import (
	"bytes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"net"

	"github.com/vompressor/go_sconn/protocol"
)

type msgProtocol struct {
	Len uint32
}

func (m *msgProtocol) SetBodyLen(b int) {
	m.Len = uint32(b)
}

func (m *msgProtocol) GetBodyLen() int {
	return int(m.Len)
}

type BlockModeGetter func(cipher.Block, []byte) cipher.BlockMode

type BlockSConn struct {
	encrypterGetter BlockModeGetter
	decrypterGetter BlockModeGetter
	cip             cipher.Block
	net.Conn
	buf *bytes.Buffer
}

func Upgrade(conn net.Conn, cip cipher.Block, eg, dg BlockModeGetter) *BlockSConn {
	return &BlockSConn{
		cip:             cip,
		Conn:            conn,
		buf:             new(bytes.Buffer),
		encrypterGetter: eg,
		decrypterGetter: dg,
	}
}

func (s *BlockSConn) Read(b []byte) (int, error) {
	var bufRead int
	var err error
	proto := new(msgProtocol)

	if s.buf.Len() > 0 {
		bufRead, err = s.buf.Read(b)
		//		if bufRead == len(b) {
		return bufRead, err
		//		}
	}

	_, body, err := protocol.ReadProtocol(s.Conn, proto)
	if err != nil {
		return 0, err
	}
	decryptedBody, _ := s.decrypt(body)

	_, err = s.buf.Write(decryptedBody)
	if err != nil {
		return bufRead, err
	}

	return s.buf.Read(b)
}

func (s *BlockSConn) Write(b []byte) (int, error) {
	cipherText := s.encrypt(b)
	proto := new(msgProtocol)

	_, err := protocol.WriteProtocol(s.Conn, proto, cipherText)
	if err != nil {
		return 0, err
	}

	return len(b), nil
}

func (s *BlockSConn) encrypt(b []byte) []byte {
	bSize := s.cip.BlockSize()
	paddingSize := 0

	if mod := len(b) % bSize; mod != 0 {
		paddingSize = bSize - mod
		padding := make([]byte, bSize-mod)
		memsetRepeat(padding, byte(paddingSize))
		b = append(b, padding...)
	}

	cipherText := make([]byte, bSize+len(b))
	iv := cipherText[:bSize]
	rand.Read(iv)

	mode := s.encrypterGetter(s.cip, iv)

	mode.CryptBlocks(cipherText[bSize:], b)

	return cipherText
}

func (s *BlockSConn) decrypt(b []byte) ([]byte, error) {
	bSize := s.cip.BlockSize()
	if mod := len(b) % bSize; mod != 0 {
		return nil, fmt.Errorf("size must be a multiple of %d", bSize)
	}

	iv := b[:bSize]
	cipherText := b[bSize:]

	mode := s.decrypterGetter(s.cip, iv)
	plainText := make([]byte, len(b)-bSize)

	mode.CryptBlocks(plainText, cipherText)

	return paddingChecker(plainText), nil
}

func memsetRepeat(a []byte, v byte) {
	if len(a) == 0 {
		return
	}
	a[0] = v
	for bp := 1; bp < len(a); bp *= 2 {
		copy(a[bp:], a[:bp])
	}
}

func paddingChecker(cipherText []byte) []byte {
	last := cipherText[len(cipherText)-1]

	if !(0 < last && last < 15) {
		return cipherText
	}

	paddingSlice := cipherText[len(cipherText)-int(last):]

	for _, n := range paddingSlice {
		if n != last {
			return nil
		}
	}
	return cipherText[:len(cipherText)-int(last)]
}
