// Copyright 2021 vompressor. All rights reserved.
// license that can be found in https://github.com/vompressor/go_sconn/blob/master/LICENSE.

package protocol

type BasicProtocol struct {
	Type    uint16
	Method  uint16
	Seq     uint32
	BodyLen uint32
}

// GetBodyLen return size of BasicProtocol
func (bp *BasicProtocol) GetBodyLen() int {
	return int(bp.BodyLen)
}

// GetBodyLen set size of BasicProtocol
func (bp *BasicProtocol) SetBodyLen(l int) {
	bp.BodyLen = uint32(l)
}
