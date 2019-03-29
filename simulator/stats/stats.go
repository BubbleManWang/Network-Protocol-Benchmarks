package stats

import (
	"../core"
)

var IsAlive bool

var XRecvCh, XLossCh, XSendCh chan *core.Packet
var YRecvCh, YLossCh, YSendCh chan *core.Packet

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
