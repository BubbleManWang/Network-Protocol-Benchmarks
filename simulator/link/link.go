package link

var IsAlive bool

var XPushCh chan []byte
var XPullCh chan []byte
var YPushCh chan []byte
var YPullCh chan []byte

var _rateMin, _rateMax int
var _lossMin, _lossMax int
var _delayMin, _delayMax int

func Spawn(rateMin, rateMax, lossMin, lossMax, delayMin, delayMax int) error {
	// TODO: check args

	_rateMin = rateMin
	_rateMax = rateMax
	_lossMin = lossMin
	_lossMax = lossMax
	_delayMin = delayMin
	_delayMax = delayMax

	XPushCh = make(chan []byte)
	XPullCh = make(chan []byte)
	YPushCh = make(chan []byte)
	YPullCh = make(chan []byte)

	IsAlive = true

	// TODO: routines

	return nil
}

func Kill() {
	if !IsAlive {
		return
	}

	IsAlive = false

	close(XPushCh)
	close(XPullCh)
	close(YPushCh)
	close(YPullCh)
}
