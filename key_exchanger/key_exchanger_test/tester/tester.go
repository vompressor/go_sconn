package main

import (
	"net"

	"github.com/vompressor/go_sconn/key_exchanger"
	"github.com/vompressor/go_sconn/sconn/stream/chacha20_upgrader"
)

func main() {
	l, err := net.Listen("unix", "/tmp/testsocket")
	if err != nil {
		println(err.Error())
		return
	}
	defer l.Close()

	excl := key_exchanger.NewExcListener(l, chacha20_upgrader.Upgrade)

	sc, err := excl.Accept()
	if err != nil {
		println(err.Error())
		return
	}
	defer sc.Close()

	buf := make([]byte, 512)
	n, _ := sc.Read(buf)
	println(string(buf[:n]))
}
