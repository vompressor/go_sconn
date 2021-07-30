# sconn
sconn is encrypts and decrypts socket communication with a private key.

```
// sconn.go

type SConn interface {

    // net.Conn to be used for encryption
	net.Conn

    // read to net.Conn, decrypt msg
	Read([]byte) (int, error)

    // encrypt msg, write to net.Conn
	Write([]byte) (int, error)
}
```

```
// sconn.go

// get net.Conn and private key, upgrade net.Conn to SConn
type ConnUpgrader func(net.Conn, []byte) (SConn, error)
```

## example
```
// server.go

package main

import (
	"crypto/sha256"
	"net"

	"github.com/vompressor/go_sconn/sconn"
	"github.com/vompressor/go_sconn/sconn/stream/chacha20_upgrader"
)

func main() {
	var sc sconn.SConn

	l, err := net.Listen("tcp", "localhost:54777")
	if err != nil {
		panic(err.Error())
	}

	cc, err := l.Accept()
	if err != nil {
		panic(err.Error())
	}

	k := sha256.Sum256([]byte("hello"))

	sc, err = chacha20_upgrader.Upgrade(cc, k[:])
	if err != nil {
		panic(err.Error())
	}
	sc.Write([]byte("hello"))
}

```
```
//client.go

package main

import (
	"crypto/sha256"
	"net"

	"github.com/vompressor/go_sconn/sconn"
	"github.com/vompressor/go_sconn/sconn/stream/chacha20_upgrader"
)

func main() {
	var sc sconn.SConn
	cc, _ := net.Dial("tcp", "localhost:54777")

	k := sha256.Sum256([]byte("hello"))

	sc, err := chacha20_upgrader.Upgrade(cc, k[:])
	if err != nil {
		panic(err.Error())
	}

	buf := make([]byte, 1024)
	n, _ := sc.Read(buf)

	println(string(buf[:n]))
}

	sc.Close()
```