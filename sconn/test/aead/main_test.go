package main_test

import (
	"testing"

	"github.com/vompressor/go_sconn/key_exchanger"
	"github.com/vompressor/go_sconn/sconn/aead"
	"github.com/vompressor/go_sconn/sconn/aead/chacha20poly1305_upgrader"
)

const sock_path = "/tmp/test.sock"

func TestAEADSockClient(t *testing.T) {
	kc, err := key_exchanger.ExcDial("unix", sock_path, chacha20poly1305_upgrader.Upgrade)
	if err != nil {
		t.Fatal(err.Error())
	}

	sc := kc.(*aead.AEADSConn)

	buf := make([]byte, 250)
	n, err := sc.Read(buf)

	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(string(buf[:n]))
}
