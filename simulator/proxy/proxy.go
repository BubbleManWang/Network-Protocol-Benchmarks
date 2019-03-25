package proxy

import (
	"fmt"
	"net"
)

var IsAlive bool

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

	IsAlive = true

	// TODO: send/recv routines

	return nil
}

func Shutdown() {
	if !IsAlive {
		return
	}

	IsAlive = false

	_mConn.Close()
}
