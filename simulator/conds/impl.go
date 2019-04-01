package conds

import (
	"math/rand"
	"time"
)

func update() {
	rand.Seed(time.Now().UnixNano())
	for IsAlive {
		time.Sleep(time.Second)

		// delay
		if delayDiff := _delayMax - _delayMin; delayDiff > 0 {
			_condDelay = int64(_delayMin + rand.Intn(delayDiff))
			_condDelay *= time.Millisecond.Nanoseconds()
		}

		// TODO: update _condRate, _condLoss
	}
}
