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
		case xRecv := <-proxy.XRecvCh:
			link.YPushCh <- xRecv
		case yRecv := <-proxy.YRecvCh:
			link.XPushCh <- yRecv
		case xPull := <-link.XPullCh:
			proxy.XSendCh <- xPull
		case yPull := <-link.YPullCh:
			proxy.YSendCh <- yPull
		}
	}
}
