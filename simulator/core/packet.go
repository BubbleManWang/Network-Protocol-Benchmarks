package core

type Packet struct {
	ID int

	Payload []byte
	Size    int

	RecvTime   int64
	ExpiryTime int64
	SendTime   int64
}
