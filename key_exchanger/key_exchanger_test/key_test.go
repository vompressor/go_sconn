package key_exchangetrtest

import (
	"testing"

	"github.com/vompressor/go_sconn/key_exchanger"
	"github.com/vompressor/go_sconn/sconn/stream/chacha20_upgrader"
)

func TestExchange(t *testing.T) {
	sc, err := key_exchanger.ExcDial("unix", "/tmp/testsocket", chacha20_upgrader.Upgrade)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer sc.Close()

	sc.Write([]byte("hi server"))

}
