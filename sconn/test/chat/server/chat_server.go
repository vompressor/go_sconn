package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/vompressor/go_sconn/key_exchanger"
	"github.com/vompressor/go_sconn/sconn/aead/chacha20poly1305_upgrader"
)

const chat_net = "unix"
const chat_addr = "/tmp/test_chat.sock"

func main() {
	os.Remove(chat_addr)
	l, err := net.Listen(chat_net, chat_addr)
	errCheck(err)

	kl := key_exchanger.NewExcListener(l, chacha20poly1305_upgrader.Upgrade)
	defer kl.Close()

	cc, err := kl.Accept()
	errCheck(err)
	fmt.Printf("%s - accepted\n", cc.RemoteAddr())
	defer cc.Close()
	buf := make([]byte, 255)
	n, err := cc.Read(buf)
	errCheck(err)
	println(string(buf[:n]))

	for {
		go func() {
			for {
				n, _ := cc.Read(buf)
				print(string(buf[:n]))
			}
		}()
		var scanBuf string
		fmt.Scanln(&scanBuf)
		cc.Write([]byte(scanBuf))
	}

}

func errCheck(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
