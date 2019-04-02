package conds

import (
	"sync/atomic"
)

var IsAlive bool

var RateCh, LossCh, DelayCh chan int

var _rateMin, _rateMax int
var _lossMin, _lossMax int
var _delayMin, _delayMax int

var _condDelay int64

var _condLossPool []bool
var _lossCounter uint32

func Spawn(rateMin, rateMax, lossMin, lossMax, delayMin, delayMax int) error {
	// TODO: check args?

	_rateMin = rateMin
	_rateMax = rateMax
	_lossMin = lossMin
	_lossMax = lossMax
	_delayMin = delayMin
	_delayMax = delayMax

	_condDelay = int64(delayMin)

	RateCh = make(chan int)
	LossCh = make(chan int)
	DelayCh = make(chan int)

	IsAlive = true

	go update()

	return nil
}

func Delay() int64 {
	return _condDelay
}

func Loss(packetSize int) bool {
	// TODO: rate limit ?
	return _condLossPool[atomic.AddUint32(&_lossCounter, 1)%100]
}

func Kill() {
	if !IsAlive {
		return
	}

	IsAlive = false

	close(RateCh)
	close(LossCh)
	close(DelayCh)
}
