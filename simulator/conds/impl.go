package conds

import (
	"math/rand"
	"sync/atomic"
	"time"
)

func update() {
	rand.Seed(time.Now().UnixNano())
	for IsAlive {
		time.Sleep(time.Second)

		updateDelay()
		updateLoss()

		// TODO: rate ?
	}
}

func updateDelay() {
	delayDiff := _delayMax - _delayMin
	delayRand := _delayMin + rand.Intn(delayDiff)
	// TODO: DelayCh <- delayRand
	_condDelay = int64(delayRand) * time.Millisecond.Nanoseconds()
}

func updateLoss() {
	_condLossPool = make([]bool, 100)

	lossDiff := _lossMax - _lossMin
	lossRand := _lossMin + rand.Intn(lossDiff)

	for i := 0; i < lossRand; i++ {
		current := rand.Intn(100)
		for _condLossPool[current] {
			current = rand.Intn(100)
		}
		_condLossPool[current] = true
	}

	atomic.StoreUint32(&_lossCounter, 0)
}
