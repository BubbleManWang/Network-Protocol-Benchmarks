package logs

import "errors"

var IsAlive bool

func Spawn() error {
	// TODO
	return errors.New("not implemented")
}

func LogStats() {
	// TODO: write to `/<timestamp>/stats.log`
}

func LogTrace() {
	// TODO: write to `/<timestamp>/trace.log`
}

func Kill() {
	// TODO
}
