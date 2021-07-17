package secure_conn

type SecureConnHeader struct {
	Type    uint16
	BodyLen uint32
}

func (h *SecureConnHeader) GetBodyLen() int {
	return int(h.BodyLen)
}
func (h *SecureConnHeader) SetBodyLen(len int) {
	h.BodyLen = uint32(len)
}
