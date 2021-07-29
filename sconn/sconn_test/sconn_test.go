package main_test

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"net"
	"testing"

	"github.com/vompressor/go_sconn/sconn"
	"github.com/vompressor/go_sconn/sconn/stream/chacha20_upgrader"
)

func TestSconn(t *testing.T) {
	var sc sconn.SConn
	cc, _ := net.Dial("tcp", "localhost:54777")

	k := sha256.Sum256([]byte("hello"))

	sc, err := chacha20_upgrader.Upgrade(cc, k[:])
	if err != nil {
		panic(err.Error())
	}

	buf := make([]byte, 1024)
	n, _ := sc.Read(buf)

	println(string(buf[:n]))

	sc.Close()
}

func TestStream(t *testing.T) {
	k := sha256.Sum256([]byte("hello"))
	v, _ := aes.NewCipher(k[:])
	iv := sha256.Sum256([]byte("bye"))

	x := cipher.NewCTR(v, iv[:v.BlockSize()])

	msg := "wooohahyeah"
	buf := make([]byte, 1024)
	b := []byte(msg)
	x.XORKeyStream(buf[:len(b)], b)

	k1 := sha256.Sum256([]byte("hello"))
	v1, _ := aes.NewCipher(k1[:])

	x1 := cipher.NewCTR(v, iv[:v1.BlockSize()])

	x1.XORKeyStream(buf[:len(b)], buf[:len(b)])

	println(string(buf[:len(b)]))
}
