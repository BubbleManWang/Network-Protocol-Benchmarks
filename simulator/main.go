package main

import (
	"fmt"

	"./link"
	"./logs"
	"./proxy"
	"./stats"
)

func main() {
	var err error

	err = logs.Spawn("logs")
	if err != nil {
		fmt.Printf("cannot spawn logs > %s\n", err)
		return
	}
	defer logs.Kill()

	err = stats.Spawn()
	if err != nil {
		fmt.Printf("cannot spawn stats > %s\n", err)
		return
	}
	defer stats.Kill()

	err = link.Spawn(1024, 2048, 10, 20, 50, 100)
	if err != nil {
		fmt.Printf("cannot spawn link > %s\n", err)
		return
	}
	defer link.Kill()

	err = proxy.Spawn(1337, 9696, 6969)
	if err != nil {
		fmt.Printf("cannot spawn proxy > %s\n", err)
		return
	}
	defer proxy.Kill()

	for {
		select {
		case pkt := <-proxy.XRecvCh:
			link.YPushCh <- pkt
			stats.XRecvCh <- pkt
		case pkt := <-link.XLossCh:
			stats.XLossCh <- pkt
		case pkt := <-link.XPullCh:
			proxy.XSendCh <- pkt
			stats.XSendCh <- pkt
		case pkt := <-proxy.YRecvCh:
			link.XPushCh <- pkt
			stats.YRecvCh <- pkt
		case pkt := <-link.YLossCh:
			stats.YLossCh <- pkt
		case pkt := <-link.YPullCh:
			proxy.YSendCh <- pkt
			stats.YSendCh <- pkt
		}
	}
}
