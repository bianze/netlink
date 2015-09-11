package nl

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"testing"
)

func (msg *TcMsg) write(b []byte) {
	native := NativeEndian()
	b[0] = msg.Family
	copy(b[1:4], msg.Pad[:])
	native.PutUint32(b[4:8], uint32(msg.Ifindex))
	native.PutUint32(b[8:12], msg.Handle)
	native.PutUint32(b[12:16], msg.Parent)
	native.PutUint32(b[16:20], msg.Info)
}

func (msg *TcMsg) serializeSafe() []byte {
	length := SizeofTcMsg
	b := make([]byte, length)
	msg.write(b)
	return b
}

func deserializeTcMsgSafe(b []byte) *TcMsg {
	var msg = TcMsg{}
	binary.Read(bytes.NewReader(b[0:SizeofTcMsg]), NativeEndian(), &msg)
	return &msg
}

func TestTcMsgDeserializeSerialize(t *testing.T) {
	var orig = make([]byte, SizeofTcMsg)
	rand.Read(orig)
	safemsg := deserializeTcMsgSafe(orig)
	msg := DeserializeTcMsg(orig)
	testDeserializeSerialize(t, orig, safemsg, msg)
}

func (msg *TcActionMsg) write(b []byte) {
	b[0] = msg.Family
	copy(b[1:4], msg.Pad[:])
}

func (msg *TcActionMsg) serializeSafe() []byte {
	length := SizeofTcActionMsg
	b := make([]byte, length)
	msg.write(b)
	return b
}

func deserializeTcActionMsgSafe(b []byte) *TcActionMsg {
	var msg = TcActionMsg{}
	binary.Read(bytes.NewReader(b[0:SizeofTcActionMsg]), NativeEndian(), &msg)
	return &msg
}

func TestTcActionMsgDeserializeSerialize(t *testing.T) {
	var orig = make([]byte, SizeofTcActionMsg)
	rand.Read(orig)
	safemsg := deserializeTcActionMsgSafe(orig)
	msg := DeserializeTcActionMsg(orig)
	testDeserializeSerialize(t, orig, safemsg, msg)
}

func (msg *TcRateSpec) write(b []byte) {
	native := NativeEndian()
	b[0] = msg.CellLog
	b[1] = msg.Linklayer
	native.PutUint16(b[2:4], msg.Overhead)
	native.PutUint16(b[4:6], uint16(msg.CellAlign))
	native.PutUint16(b[6:8], msg.Mpu)
	native.PutUint32(b[8:12], msg.Rate)
}

func (msg *TcRateSpec) serializeSafe() []byte {
	length := SizeofTcRateSpec
	b := make([]byte, length)
	msg.write(b)
	return b
}

func deserializeTcRateSpecSafe(b []byte) *TcRateSpec {
	var msg = TcRateSpec{}
	binary.Read(bytes.NewReader(b[0:SizeofTcRateSpec]), NativeEndian(), &msg)
	return &msg
}

func TestTcRateSpecDeserializeSerialize(t *testing.T) {
	var orig = make([]byte, SizeofTcRateSpec)
	rand.Read(orig)
	safemsg := deserializeTcRateSpecSafe(orig)
	msg := DeserializeTcRateSpec(orig)
	testDeserializeSerialize(t, orig, safemsg, msg)
}
