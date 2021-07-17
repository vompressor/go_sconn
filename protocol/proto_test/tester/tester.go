package main

import (
	"fmt"
	"net"

	"github.com/vompressor/go_sconn/protocol"
)

func main() {
	listener, err := net.Listen("unix", "/tmp/tteesstt.sock")
	if err != nil {
		println(err.Error())
		return
	}
	defer listener.Close()

	conn, _ := listener.Accept()
	defer conn.Close()
	h := &Head{}

	_, msg, err := protocol.ReadProtocol(conn, h)
	if err != nil {
		println(err.Error())
		return
	}

	fmt.Printf("%s\n", msg)
}

type Head struct {
	Type uint32
	Len  uint32
}

func (h *Head) SetBodyLen(i int) {
	h.Len = uint32(i)
}

func (h *Head) GetBodyLen() int {
	return int(h.Len)
}
