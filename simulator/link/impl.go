package link

import "time"

func pass() {
	var data []byte
	var size int
	var forX bool

	for IsAlive {
		select {
		case pkt, ok := <-XPushCh:
			if !ok {
				// TODO: log err - channel closed
				Kill()
				return
			}

			data = pkt
			size = len(pkt)
			forX = true
		case pkt, ok := <-YPushCh:
			if !ok {
				// TODO: log err - channel closed
				Kill()
				return
			}

			data = pkt
			size = len(pkt)
			forX = false
		}

		// TODO: rate limiting, loss chance, additional delay

		_queueMutex.Lock()
		{
			pkt := packet{
				data,
				size,
			}

			if forX {
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
				XPullCh <- pkt.data
			}
			_xQueue = make([]packet, 0)

			for _, pkt := range _yQueue {
				YPullCh <- pkt.data
			}
			_yQueue = make([]packet, 0)
		}
		_queueMutex.Unlock()

		// TODO: update condition modifiers
	}
}
