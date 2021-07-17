package key_exchanger

import (
	"fmt"
	"io"
	"net"

	"github.com/vompressor/go_sconn/ecdh"
	"github.com/vompressor/go_sconn/secure_conn"
)

/*

 1.

 client - [client hello] + [client pub pem] -> server

 2.

 server - [server hello] + [server pub pem] -> client

*/

type Exchanger struct {
}

type ExchangeListner struct {
	net.Listener
}

func NewExcListener(l net.Listener) *ExchangeListner {
	chg := ExchangeListner{}
	chg.Listener = l
	return &chg
}

func (c *ExchangeListner) Close() error {
	return c.Listener.Close()
}

func (c *ExchangeListner) Accept() (*secure_conn.BlockSConn, error) {
	conn, err := c.Listener.Accept()

	if err != nil {
		return nil, err
	}

	return ServerSideUpgrade(conn)
}

func (c *ExchangeListner) Addr() net.Addr {
	return c.Listener.Addr()
}

func ServerSideUpgrade(c net.Conn) (*secure_conn.BlockSConn, error) {
	changer, err := ecdh.NewKXchn()
	if err != nil {
		return nil, err
	}

	remote_pub, err := readClientHello(c)
	println("hi client!")
	if err != nil {
		return nil, err
	}
	lpub, err := changer.GeneratePub()
	if err != nil {
		return nil, err
	}

	err = writeServerHello(c, lpub)
	println("server hello")

	if err != nil {
		return nil, err
	}

	pub, err := ecdh.LoadDHPubKey(remote_pub)
	if err != nil {
		return nil, err
	}
	key := changer.GenerateSharedKey(pub)

	fmt.Printf("%x\n", key)

	return secure_conn.NewAesSConn(c, key)
}

func Upgrade(c net.Conn) (*secure_conn.BlockSConn, error) {
	changer, err := ecdh.NewKXchn()
	if err != nil {
		return nil, err
	}

	lpub, err := changer.GeneratePub()
	if err != nil {
		return nil, err
	}

	err = writeClientHello(c, lpub)

	if err != nil {
		return nil, err
	}

	rpub, err := readServerHello(c)
	if err != nil {
		return nil, err
	}

	dhrpub, err := ecdh.LoadDHPubKey(rpub)
	if err != nil {
		return nil, err
	}

	sharedKey := changer.GenerateSharedKey(dhrpub)
	fmt.Printf("%x\n", sharedKey)
	return secure_conn.NewAesSConn(c, sharedKey)
}

func readServerHello(r io.Reader) (d []byte, err error) {
	d, err = readMsg(r, type_handshake, method_server_hello)
	return
}
func writeServerHello(w io.Writer, b []byte) (err error) {
	err = writeMsg(w, type_handshake, method_server_hello, b)
	return
}

func readClientHello(r io.Reader) (d []byte, err error) {
	d, err = readMsg(r, type_handshake, method_client_hello)
	return
}
func writeClientHello(w io.Writer, b []byte) (err error) {
	err = writeMsg(w, type_handshake, method_client_hello, b)
	return
}
