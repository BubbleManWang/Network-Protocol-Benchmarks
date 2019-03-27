package proxy

import (
	"fmt"
	"net"

	"../core"
)

var IsAlive bool

var XRecvCh chan *core.Packet
var XSendCh chan *core.Packet
var YRecvCh chan *core.Packet
var YSendCh chan *core.Packet

var _mConn *net.UDPConn
var _xAddr *net.UDPAddr
var _yAddr *net.UDPAddr

func Spawn(listenPort, leftPort, rightPort int) error {
	mAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("localhost:%d", listenPort))
	if err != nil {
		return err
	}

	xAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("localhost:%d", leftPort))
	if err != nil {
		return err
	}

	yAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("localhost:%d", rightPort))
	if err != nil {
		return nil
	}

	mConn, err := net.ListenUDP("udp", mAddr)
	if err != nil {
		return err
	}

	_mConn = mConn
	_xAddr = xAddr
	_yAddr = yAddr

	XRecvCh = make(chan *core.Packet)
	XSendCh = make(chan *core.Packet)
	YRecvCh = make(chan *core.Packet)
	YSendCh = make(chan *core.Packet)

	IsAlive = true

	go recv()
	go send()

	return nil
}

func Kill() {
	if !IsAlive {
		return
	}

	IsAlive = false

	close(XRecvCh)
	close(XSendCh)
	close(YRecvCh)
	close(YSendCh)

	_mConn.Close()
}
