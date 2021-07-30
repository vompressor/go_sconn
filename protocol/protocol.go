// Copyright 2021 vompressor. All rights reserved.
// license that can be found in https://github.com/vompressor/go_sconn/blob/master/LICENSE.

package protocol

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

type ProtocolHeader interface {
	GetBodyLen() int
	SetBodyLen(int)
}

func EncodeProtocolByte(head ProtocolHeader, msg []byte) ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))

	head.SetBodyLen(len(msg))
	err := binary.Write(buf, binary.BigEndian, head)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.BigEndian, msg)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func DecodeProtocolByte(head ProtocolHeader, msg []byte) ([]byte, error) {
	reader := bytes.NewReader(msg)
	binary.Read(reader, binary.BigEndian, head)
	e := head.GetBodyLen()
	s := binary.Size(head)
	return msg[s : e+s], nil
}

func DecodeHeader(head ProtocolHeader, headerByte []byte) error {
	if len(headerByte) < binary.Size(head) {
		return errors.New("size mismatch")
	}

	reader := bytes.NewReader(headerByte)
	err := binary.Read(reader, binary.BigEndian, head)
	if err != nil {
		return err
	}
	return nil
}

func WriteProtocol(writer io.Writer, head ProtocolHeader, msg []byte) (int, error) {
	d, err := EncodeProtocolByte(head, msg)
	if err != nil {
		return 0, err
	}
	n, err := writer.Write(d)
	if err != nil {
		return n, err
	}
	if head.GetBodyLen()+binary.Size(head) != n {
		return n, errors.New("msg write size err")
	}
	return n, nil
}

func ReadProtocol(reader io.Reader, head ProtocolHeader) (int, []byte, error) {
	headLen := binary.Size(head)
	headByte := make([]byte, headLen)
	rlen, err := reader.Read(headByte)
	if err != nil {
		return rlen, nil, err
	}

	if rlen != headLen {
		return rlen, nil, errors.New("header size err")
	}

	err = DecodeHeader(head, headByte)
	if err != nil {
		return rlen, nil, err
	}

	msgLen := head.GetBodyLen()
	msgByte := make([]byte, msgLen)
	mlem, err := reader.Read(msgByte)
	if err != nil {
		return rlen + mlem, nil, err
	}

	if mlem != msgLen {
		return rlen + mlem, nil, errors.New("msg size err")
	}

	return rlen + mlem, msgByte, nil
}
