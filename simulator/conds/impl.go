package conds

import (
	"math/rand"
	"time"
)

func update() {
	rand.Seed(time.Now().UnixNano())
	for IsAlive {
		time.Sleep(time.Second)

		updateDelay()

		// TODO: loss, rate ?
	}
}

func updateDelay() {
	delayDiff := _delayMax - _delayMin
	delayRand := _delayMin + rand.Intn(delayDiff)
	// TODO: DelayCh <- delayRand
	_condDelay = int64(delayRand) * time.Millisecond.Nanoseconds()
}
