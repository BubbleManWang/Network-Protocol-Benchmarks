package conds

import (
	"../logs"
)

var IsAlive bool

func Spawn(rateMin, rateMax, lossMin, lossMax, delayMin, delayMax int) error {
	// TODO
	logs.LogTrace("not implemented yet")
	return nil
}

func Delay() int64 {
	// TODO: delay & jitter
	return 0
}

func Loss(packetSize int) bool {
	// TODO: rate limit & loss chance
	return false
}

func Kill() {
	// TODO
}
