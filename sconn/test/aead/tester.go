package main

import (
	"log"
	"net"
	"os"

	"github.com/vompressor/go_sconn/key_exchanger"
	"github.com/vompressor/go_sconn/sconn/aead"
	"github.com/vompressor/go_sconn/sconn/aead/chacha20poly1305_upgrader"
)

const sock_path = "/tmp/test.sock"

func main() {
	os.RemoveAll(sock_path)
	l, err := net.Listen("unix", sock_path)
	if err != nil {
		log.Fatal(err.Error())
	}
	kl := key_exchanger.NewExcListener(l, chacha20poly1305_upgrader.Upgrade)
	defer kl.Close()

	sc, err := kl.Accept()

	if err != nil {
		log.Fatal(err.Error())
	}
	defer sc.Close()

	ac := sc.(*aead.AEADSConn)

	_, err = ac.Write([]byte("woooo"))

	if err != nil {
		log.Fatal(err.Error())
	}
	
}
