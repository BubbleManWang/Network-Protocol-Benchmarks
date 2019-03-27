package proxy

import (
	"../core"
)

func recv() {
	buf := make([]byte, 1500)
	for IsAlive {
		len, addr, err := _mConn.ReadFromUDP(buf)
		if err != nil {
			// TODO: log err - failed to read
			Kill()
			return
		}

		pkt := core.GeneratePacket()
		pkt.Payload = make([]byte, len)
		pkt.Size = len
		copy(pkt.Payload, buf)

		if addr.Port == _xAddr.Port {
			XRecvCh <- pkt
		} else if addr.Port == _yAddr.Port {
			YRecvCh <- pkt
		} else {
			// TODO: log warn = data from unexpected port
		}
	}
}

func send() {
	for IsAlive {
		select {
		case pkt, ok := <-XSendCh:
			if !ok {
				// TODO: log err - channel closed
				Kill()
				return
			}

			_, err := _mConn.WriteToUDP(pkt.Payload, _xAddr)
			if err != nil {
				// TODO: log err - failed to write
				Kill()
				return
			}
		case pkt, ok := <-YSendCh:
			if !ok {
				// TODO: log err - channel closed
				Kill()
				return
			}

			_, err := _mConn.WriteToUDP(pkt.Payload, _yAddr)
			if err != nil {
				// TODO: log err - failed to write
				Kill()
				return
			}
		}
	}
}
