package main

import (
	"net"

	"github.com/vompressor/go_sconn/key_exchanger"
)

func main() {
	l, err := net.Listen("unix", "/tmp/testsocket")
	if err != nil {
		println(err.Error())
		return
	}
	defer l.Close()

	excl := key_exchanger.NewExcListener(l)

	_, err = excl.Accept()
	if err != nil {
		println(err.Error())
		return
	}

}
