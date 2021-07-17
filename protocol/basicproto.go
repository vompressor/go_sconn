package protocol

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
