package core

import "sync/atomic"

var _packetCounter uint64

type Packet struct {
	ID uint64

	Payload []byte
	Size    int

	RecvTime   int64
	ExpiryTime int64
	SendTime   int64
}

func GenPacket() *Packet {
	atomic.AddUint64(&_packetCounter, 1)

	pkt := Packet{}
	pkt.ID = atomic.LoadUint64(&_packetCounter)

	return &pkt
}
