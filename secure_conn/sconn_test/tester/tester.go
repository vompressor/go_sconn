package main

import (
	"bufio"
	"fmt"
	"net"

	"github.com/vompressor/go_sconn/key_exchanger"
)

func main() {

	// LISTEN
	l, err := net.Listen("unix", "/tmp/testsocket")
	if err != nil {
		println("SERVER:: " + err.Error())
		return
	}
	defer l.Close()
	println("SERVER:: created listener")

	// ACCEPT
	c, err := l.Accept()
	if err != nil {
		println("SERVER:: " + err.Error())
		return
	}
	defer c.Close()
	println("SERVER:: accepted - " + c.RemoteAddr().String())

	// UPGRADE
	sconn, err := key_exchanger.ServerSideUpgrade(c)
	if err != nil {
		println("SERVER:: " + err.Error())
		return
	}
	println("SERVER:: key exchanged")

	fmt.Printf("key - %x", sconn.Key)

	writer := bufio.NewWriter(sconn)
	_, err = writer.WriteString("hello\n")
	if err != nil {
		println("SERVER:: " + err.Error())
		return
	}
	writer.Flush()
}
