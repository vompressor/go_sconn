package main_test

import (
	"testing"

	"github.com/vompressor/go_sconn/key_exchanger"
	"github.com/vompressor/go_sconn/sconn/block/aes_cbc_upgrader"
)

const sock_path = "/tmp/test.sock"

func TestAEADSockClient(t *testing.T) {
	kc, err := key_exchanger.ExcDial("unix", sock_path, aes_cbc_upgrader.Upgrade)
	if err != nil {
		t.Fatal(err.Error())
	}

	buf := make([]byte, 250)
	n, err := kc.Read(buf)

	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(string(buf[:n]))
}
