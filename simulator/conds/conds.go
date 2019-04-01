package conds

import "errors"

var IsAlive bool

func Spawn(rateMin, rateMax, lossMin, lossMax, delayMin, delayMax int) error {
	// TODO
	return errors.New("not implemented")
}

func Loss(packetSize int) bool {
	// TODO: rate limiting, loss chance
	return false
}

func Delay() int {
	// TODO
	return 0
}

func Kill() {
	// TODO
}
