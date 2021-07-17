package key_exchangetrtest

import (
	"net"
	"testing"

	"github.com/vompressor/go_sconn/key_exchanger"
)

func TestExchange(t *testing.T) {
	conn, err := net.Dial("unix", "/tmp/testsocket")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer conn.Close()

	_, err = key_exchanger.Upgrade(conn)
	if err != nil {
		t.Fatal(err.Error())
	}

}
