package main

import (
	"fmt"

	"./link"
	"./proxy"
)

func main() {
	var err error

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
		case pkt := <-proxy.YRecvCh:
			link.XPushCh <- pkt
		case pkt := <-link.XPullCh:
			proxy.XSendCh <- pkt
		case pkt := <-link.YPullCh:
			proxy.YSendCh <- pkt
		}
	}
}
