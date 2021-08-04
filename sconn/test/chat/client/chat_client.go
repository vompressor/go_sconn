package main

import (
	"fmt"
	"log"

	"github.com/vompressor/go_sconn/key_exchanger"
	"github.com/vompressor/go_sconn/sconn/aead/chacha20poly1305_upgrader"
)

const chat_net = "unix"
const chat_addr = "/tmp/test_chat.sock"

func main() {

	kc, err := key_exchanger.ExcDial(chat_net, chat_addr, chacha20poly1305_upgrader.Upgrade)
	errCheck(err)
	fmt.Printf("%s - dialed\n", kc.RemoteAddr())
	defer kc.Close()
	_, err = kc.Write([]byte("hello server!"))
	errCheck(err)

	for {
		go func() {
			buf := make([]byte, 255)
			for {
				n, _ := kc.Read(buf)
				print(string(buf[:n]))
			}
		}()

		var scanBuf string
		fmt.Scanln(&scanBuf)
		kc.Write([]byte(scanBuf))
	}

}

func errCheck(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
