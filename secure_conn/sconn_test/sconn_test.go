package sconn_test

import (
	"bufio"
	"net"
	"testing"
	"time"

	"github.com/vompressor/go_sconn/key_exchanger"
)

func TestCreateSconn(t *testing.T) {
	time.Sleep(time.Millisecond * 3)

	// DIAL
	c, err := net.Dial("unix", "/tmp/testsocket")
	if err != nil {
		t.Fatal("CLIENT:: " + err.Error())
	}
	t.Log("CLIENT:: dialed - " + c.RemoteAddr().String())

	// UPGRADE
	sconn, err := key_exchanger.Upgrade(c)
	if err != nil {
		t.Fatal("CLIENT:: " + err.Error())
	}
	t.Log("CLIENT:: key exchanged")

	t.Logf("key - %x\n", sconn.Key)
	reader := bufio.NewReader(sconn)
	str, err := reader.ReadString('\n')
	if err != nil {
		t.Fatal("SERVER:: " + err.Error())
	}
	t.Log(str)
}
