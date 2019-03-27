package stats

import (
	"../core"
)

var IsAlive bool

var XRecvCh chan *core.Packet
var XPushCh chan *core.Packet
var XLossCh chan *core.Packet
var XPullCh chan *core.Packet
var XSendCh chan *core.Packet
var YRecvCh chan *core.Packet
var YPushCh chan *core.Packet
var YLossCh chan *core.Packet
var YPullCh chan *core.Packet
var YSendCh chan *core.Packet

func Spawn() error {
	// TODO: create log file `stats-<timestamp>.log`

	XRecvCh = make(chan *core.Packet)
	XPushCh = make(chan *core.Packet)
	XLossCh = make(chan *core.Packet)
	XPullCh = make(chan *core.Packet)
	XSendCh = make(chan *core.Packet)
	YRecvCh = make(chan *core.Packet)
	YPushCh = make(chan *core.Packet)
	YLossCh = make(chan *core.Packet)
	YPullCh = make(chan *core.Packet)
	YSendCh = make(chan *core.Packet)

	IsAlive = true

	// TODO: go record()
	// TODO: go flush()

	return nil
}

func Kill() {
	if !IsAlive {
		return
	}

	IsAlive = false

	close(XRecvCh)
	close(XPushCh)
	close(XLossCh)
	close(XPullCh)
	close(XSendCh)
	close(YRecvCh)
	close(YPushCh)
	close(YLossCh)
	close(YPullCh)
	close(YSendCh)
}
