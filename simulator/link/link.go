package link

import "errors"

var IsAlive bool

var XPushCh chan []byte
var XPullCh chan []byte
var YPushCh chan []byte
var YPullCh chan []byte

func Spawn() error {
	// TODO
	return errors.New("not implemented")
}

func Kill() {
	// TODO
}
