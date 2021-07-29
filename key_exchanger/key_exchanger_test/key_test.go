package key_exchangetrtest

import (
	"net"
	"testing"

	"github.com/vompressor/go_sconn/key_exchanger"
	"github.com/vompressor/go_sconn/sconn/stream/chacha20_upgrader"
)

func TestExchange(t *testing.T) {
	conn, err := net.Dial("unix", "/tmp/testsocket")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer conn.Close()

	sc, err := key_exchanger.Upgrade(conn, chacha20_upgrader.Upgrade)
	if err != nil {
		t.Fatal(err.Error())
	}

	sc.Write([]byte("hi server"))
	sc.Close()
}
