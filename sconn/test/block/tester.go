package main

import (
	"log"
	"net"
	"os"

	"github.com/vompressor/go_sconn/key_exchanger"
	"github.com/vompressor/go_sconn/sconn/block/aes_cbc_upgrader"
)

const sock_path = "/tmp/test.sock"

func main() {
	os.RemoveAll(sock_path)
	l, err := net.Listen("unix", sock_path)
	if err != nil {
		log.Fatal(err.Error())
	}
	kl := key_exchanger.NewExcListener(l, aes_cbc_upgrader.Upgrade)
	defer kl.Close()

	sc, err := kl.Accept()

	if err != nil {
		log.Fatal(err.Error())
	}
	defer sc.Close()

	_, err = sc.Write([]byte("woooo"))

	if err != nil {
		log.Fatal(err.Error())
	}

}
