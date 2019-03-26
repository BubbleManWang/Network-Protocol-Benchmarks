package proxy

func recv() {
	buf := make([]byte, 1500)
	for IsAlive {
		len, addr, err := _mConn.ReadFromUDP(buf)
		if err != nil {
			// TODO: log err - failed to read
			Shutdown()
			return
		}

		data := make([]byte, len)
		copy(data, buf)

		if addr.Port == _xAddr.Port {
			XRecvCh <- data
		} else if addr.Port == _yAddr.Port {
			YRecvCh <- data
		} else {
			// TODO: log warn = data from unexpected port
		}
	}
}

func send() {
	for IsAlive {
		select {
		case data, ok := <-XSendCh:
			if !ok {
				// TODO: log err - channel closed
				Shutdown()
				return
			}

			_, err := _mConn.WriteToUDP(data, _xAddr)
			if err != nil {
				// TODO: log err - failed to write
				Shutdown()
				return
			}
		case data, ok := <-YSendCh:
			if !ok {
				// TODO: log err - channel closed
				Shutdown()
				return
			}

			_, err := _mConn.WriteToUDP(data, _yAddr)
			if err != nil {
				// TODO: log err - failed to write
				Shutdown()
				return
			}
		}
	}
}
