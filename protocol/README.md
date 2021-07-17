# protocol

```
    type ProtocolHeader interface {
        GetBodyLen() int
        SetBodyLen(int)
    }
```

## struct definition
Structures must consist only of the following types:
 - uint8
 - uint16
 - uint32
 - uint64

## GetByteLen() int
It should be implemented to return the length of the protocol body in the header struct.

## SetByteLen(int)
It should be implemented to set the body length in header struct at length to the protocol body.

## implementation example
```
    type BasicProtocol struct {
        Type    uint16
        Method  uint16
        Seq     uint32
        BodyLen uint32
    }

    func (bp *BasicProtocol) GetBodyLen() int {
        return int(bp.BodyLen)
    }

    func (bp *BasicProtocol) SetBodyLen(l int) {
        bp.BodyLen = uint32(l)
    }
```

## protocol structure
```
    +- binary.Size(header) -+----------- len(msg) -----------+
    |       header          |              msg               |
    +-----------------------+--------------------------------+
```

## EncodeProtocolByte(head ProtocolHeader, msg []byte) ([]byte, error)

## DecodeProtocolByte(head ProtocolHeader, msg []byte) ([]byte, error)

## DecodeHeader(head ProtocolHeader, headerByte []byte) error

## ReadProtocol(reader io.Reader, head ProtocolHeader) ([]byte, error)

## WriteProtocol(writer io.Writer, head ProtocolHeader, msg []byte) error
