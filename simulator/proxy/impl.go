package proxy

import (
	"../core"
	"../logs"
)

func recv() {
	buf := make([]byte, 1500)
	for IsAlive {
		len, addr, err := _mConn.ReadFromUDP(buf)
		if err != nil {
			logs.LogTrace("cannot read from socket > %s", err)
			Kill()
			return
		}

		pkt := core.GenPacket()
		pkt.Payload = make([]byte, len)
		pkt.Size = len
		copy(pkt.Payload, buf)

		if addr.Port == _xAddr.Port {
			XRecvCh <- pkt
		} else if addr.Port == _yAddr.Port {
			YRecvCh <- pkt
		} else {
			logs.LogTrace("data read from unexpected port > %d", addr.Port)
		}
	}
}

func send() {
	for IsAlive {
		select {
		case pkt, ok := <-XSendCh:
			if !ok {
				logs.LogTrace("x send channel is closed")
				Kill()
				return
			}

			_, err := _mConn.WriteToUDP(pkt.Payload, _xAddr)
			if err != nil {
				logs.LogTrace("cannot write to socket > %s", err)
				Kill()
				return
			}
		case pkt, ok := <-YSendCh:
			if !ok {
				logs.LogTrace("y send channel is closed")
				Kill()
				return
			}

			_, err := _mConn.WriteToUDP(pkt.Payload, _yAddr)
			if err != nil {
				logs.LogTrace("cannot write to socket > %s", err)
				Kill()
				return
			}
		}
	}
}
