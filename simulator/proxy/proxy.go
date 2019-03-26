package proxy

import (
	"fmt"
	"net"
)

var IsAlive bool

var XRecvCh chan []byte
var XSendCh chan []byte
var YRecvCh chan []byte
var YSendCh chan []byte

var _mConn *net.UDPConn
var _xAddr *net.UDPAddr
var _yAddr *net.UDPAddr

func Initialize(listenPort, leftPort, rightPort int) error {
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

	XRecvCh = make(chan []byte)
	XSendCh = make(chan []byte)
	YRecvCh = make(chan []byte)
	YSendCh = make(chan []byte)

	IsAlive = true

	go recv()
	go send()

	return nil
}

func Shutdown() {
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
