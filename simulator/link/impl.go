package link

import (
	"time"

	"../conds"
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

		pkt.ExpiryTime = time.Now().UnixNano() + conds.Delay()
		if toX {
			if conds.Loss(pkt.Size) {
				XLossCh <- pkt
			} else {
				_queueMutex.Lock()
				_xQueue = append(_xQueue, pkt)
				_queueMutex.Unlock()
			}
		} else {
			if conds.Loss(pkt.Size) {
				YLossCh <- pkt
			} else {
				_queueMutex.Lock()
				_yQueue = append(_yQueue, pkt)
				_queueMutex.Unlock()
			}
		}
	}
}

func tick() {
	for IsAlive {
		time.Sleep(time.Millisecond * 8) // ~10ms, up to 125 ticks
		nowNano := time.Now().UnixNano()

		_queueMutex.Lock()
		{
			var xQueueNew []*core.Packet
			for _, pkt := range _xQueue {
				if pkt.ExpiryTime < nowNano {
					XPullCh <- pkt
				} else {
					xQueueNew = append(xQueueNew, pkt)
				}
			}
			_xQueue = xQueueNew

			var yQueueNew []*core.Packet
			for _, pkt := range _yQueue {
				if pkt.ExpiryTime < nowNano {
					YPullCh <- pkt
				} else {
					yQueueNew = append(yQueueNew, pkt)
				}
			}
			_yQueue = yQueueNew
		}
		_queueMutex.Unlock()
	}
}
