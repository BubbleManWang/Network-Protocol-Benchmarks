package link

import (
	"sync"

	"../core"
)

var IsAlive bool

var XPushCh, XLossCh, XPullCh chan *core.Packet
var YPushCh, YLossCh, YPullCh chan *core.Packet

var _queueMutex sync.Mutex
var _xQueue, _yQueue []*core.Packet

func Spawn() error {
	XPushCh = make(chan *core.Packet)
	XLossCh = make(chan *core.Packet)
	XPullCh = make(chan *core.Packet)
	YPushCh = make(chan *core.Packet)
	YLossCh = make(chan *core.Packet)
	YPullCh = make(chan *core.Packet)

	IsAlive = true

	go pass()
	go tick()

	return nil
}

func Kill() {
	if !IsAlive {
		return
	}

	IsAlive = false

	close(XPushCh)
	close(XLossCh)
	close(XPullCh)
	close(YPushCh)
	close(YLossCh)
	close(YPullCh)
}
