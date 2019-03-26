package main

import (
	"fmt"

	"./proxy"
)

func main() {
	err := proxy.Spawn(1337, 9696, 6969)
	if err != nil {
		fmt.Printf("cannot spawn proxy > %s", err)
		return
	}

	for proxy.IsAlive {
		select {
		case data := <-proxy.XRecvCh:
			proxy.YSendCh <- data
		case data := <-proxy.YRecvCh:
			proxy.XSendCh <- data
		}
	}
}
