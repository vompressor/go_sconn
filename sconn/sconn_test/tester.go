package main

import (
	"crypto/sha256"
	"net"

	"github.com/vompressor/go_sconn/sconn"
	"github.com/vompressor/go_sconn/sconn/stream/chacha20_upgrader"
)

func main() {
	var sc sconn.SConn

	l, err := net.Listen("tcp", "localhost:54777")
	if err != nil {
		panic(err.Error())
	}

	cc, err := l.Accept()
	if err != nil {
		panic(err.Error())
	}

	k := sha256.Sum256([]byte("hello"))

	sc, err = chacha20_upgrader.Upgrade(cc, k[:])
	if err != nil {
		panic(err.Error())
	}
	sc.Write([]byte("hello"))
}
