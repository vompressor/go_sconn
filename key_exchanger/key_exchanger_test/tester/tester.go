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

	excl := key_exchanger.NewExcListener(l)

	c, err := excl.Accept()
	if err != nil {
		println(err.Error())
		return
	}

	sc, err := key_exchanger.ServerSideUpgrade(c, chacha20_upgrader.Upgrade)
	if err != nil {
		panic(err.Error())
	}

	buf := make([]byte, 512)
	n, _ := sc.Read(buf)
	println(string(buf[:n]))
}
