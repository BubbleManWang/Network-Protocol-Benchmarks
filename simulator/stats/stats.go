package stats

import (
	"../core"
)

var IsAlive bool

var XRecvCh chan *core.Packet
var XLossCh chan *core.Packet
var XSendCh chan *core.Packet
var YRecvCh chan *core.Packet
var YLossCh chan *core.Packet
var YSendCh chan *core.Packet

func Spawn() error {
	XRecvCh = make(chan *core.Packet)
	XLossCh = make(chan *core.Packet)
	XSendCh = make(chan *core.Packet)
	YRecvCh = make(chan *core.Packet)
	YLossCh = make(chan *core.Packet)
	YSendCh = make(chan *core.Packet)

	IsAlive = true

	go record()
	go flush()

	return nil
}

func Kill() {
	if !IsAlive {
		return
	}

	IsAlive = false

	close(XRecvCh)
	close(XLossCh)
	close(XSendCh)
	close(YRecvCh)
	close(YLossCh)
	close(YSendCh)
}
