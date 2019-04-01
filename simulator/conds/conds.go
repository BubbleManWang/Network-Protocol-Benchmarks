package conds

var IsAlive bool

var RateCh, LossCh, DelayCh chan int

func Spawn(rateMin, rateMax, lossMin, lossMax, delayMin, delayMax int) error {
	// TODO: check args?

	RateCh = make(chan int)
	LossCh = make(chan int)
	DelayCh = make(chan int)

	IsAlive = true

	go update()

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
	if !IsAlive {
		return
	}

	IsAlive = false

	close(RateCh)
	close(LossCh)
	close(DelayCh)
}
