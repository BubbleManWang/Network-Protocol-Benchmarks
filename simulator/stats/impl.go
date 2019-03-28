package stats

import "time"

func record() {
	for IsAlive {
		select {
		case _, ok := <-XRecvCh:
			if !ok {
				// TODO: log err - channel closed
				Kill()
				return
			}

			// TODO
		case _, ok := <-XLossCh:
			if !ok {
				// TODO: log err - channel closed
				Kill()
				return
			}

			// TODO
		case _, ok := <-XSendCh:
			if !ok {
				// TODO: log err - channel closed
				Kill()
				return
			}

			// TODO
		case _, ok := <-YRecvCh:
			if !ok {
				// TODO: log err - channel closed
				Kill()
				return
			}

			// TODO
		case _, ok := <-YLossCh:
			if !ok {
				// TODO: log err - channel closed
				Kill()
				return
			}

			// TODO
		case _, ok := <-YSendCh:
			if !ok {
				// TODO: log err - channel closed
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

		// TODO: write-out accumulated stats to log file
	}
}
