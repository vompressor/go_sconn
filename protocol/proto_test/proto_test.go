package proto_test

import (
	"net"
	"testing"

	"github.com/vompressor/go_sconn/protocol"
)

type Head struct {
	Type uint32
	Len  uint32
}

func (h *Head) SetBodyLen(i int) {
	h.Len = uint32(i)
}

func (h *Head) GetBodyLen() int {
	return int(h.Len)
}

func TestEncodeProto(t *testing.T) {
	h := &Head{}
	h.Type = 1
	b, err := protocol.EncodeProtocolByte(h, []byte("qwertyuiopqwertyuiopqwertyuiopqq"))
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Logf("%x\n", b)
	t.Logf("%d\n", h.Len)

	h2 := &Head{}
	msg, err := protocol.DecodeProtocolByte(h2, b)
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Logf("%s\n", msg)
	t.Logf("%d\n", h2.Len)
	t.Logf("%#v\n", h2)

}

func TestWriteProto(t *testing.T) {
	conn, err := net.Dial("unix", "/tmp/tteesstt.sock")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer conn.Close()

	h := &Head{}
	h.Type = 1
	_, err = protocol.WriteProtocol(conn, h, []byte("hello"))
	if err != nil {
		t.Fatal(err.Error())
	}
}
