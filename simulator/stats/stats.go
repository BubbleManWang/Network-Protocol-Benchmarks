package stats

import (
	"errors"

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
	// TODO
	return errors.New("not implemented")
}

func Kill() {
	// TODO
}
