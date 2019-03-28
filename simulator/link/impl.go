package link

import (
	"time"

	"../core"
	"../logs"
)

func pass() {
	var pkt *core.Packet
	var toX bool

	for IsAlive {
		select {
		case p, ok := <-XPushCh:
			if !ok {
				logs.LogTrace("x push channel is closed")
				Kill()
				return
			}

			pkt = p
			toX = true
		case p, ok := <-YPushCh:
			if !ok {
				logs.LogTrace("y push channel is closed")
				Kill()
				return
			}

			pkt = p
			toX = false
		}

		// TODO: rate limiting, loss chance, additional delay (XLossCh <- pkt)

		_queueMutex.Lock()
		{
			if toX {
				_xQueue = append(_xQueue, pkt)
			} else {
				_yQueue = append(_yQueue, pkt)
			}
		}
		_queueMutex.Unlock()
	}
}

func tick() {
	for IsAlive {
		time.Sleep(time.Millisecond * 8) // ~10ms, up to 125 ticks

		_queueMutex.Lock()
		{
			// TODO: packet expiry check for additional delay

			for _, pkt := range _xQueue {
				XPullCh <- pkt
			}
			_xQueue = make([]*core.Packet, 0)

			for _, pkt := range _yQueue {
				YPullCh <- pkt
			}
			_yQueue = make([]*core.Packet, 0)
		}
		_queueMutex.Unlock()

		// TODO: update condition modifiers
	}
}
