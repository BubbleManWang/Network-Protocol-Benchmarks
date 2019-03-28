package stats

import (
	"time"

	"../logs"
)

func record() {
	for IsAlive {
		select {
		case _, ok := <-XRecvCh:
			if !ok {
				logs.LogTrace("x recv channel is closed")
				Kill()
				return
			}

			// TODO
		case _, ok := <-XLossCh:
			if !ok {
				logs.LogTrace("x loss channel is closed")
				Kill()
				return
			}

			// TODO
		case _, ok := <-XSendCh:
			if !ok {
				logs.LogTrace("x send channel is closed")
				Kill()
				return
			}

			// TODO
		case _, ok := <-YRecvCh:
			if !ok {
				logs.LogTrace("y recv channel is closed")
				Kill()
				return
			}

			// TODO
		case _, ok := <-YLossCh:
			if !ok {
				logs.LogTrace("y loss channel is closed")
				Kill()
				return
			}

			// TODO
		case _, ok := <-YSendCh:
			if !ok {
				logs.LogTrace("y send channel is closed")
				Kill()
				return
			}

			// TODO
		}
	}
}

func flush() {
	for IsAlive {
		time.Sleep(time.Second)

		// TODO: log accumulated stats (logs.LogStats)
	}
}
