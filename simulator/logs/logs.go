package logs

import (
	"fmt"
	"os"
	"time"
)

var IsAlive bool

var _fStats *os.File
var _fTrace *os.File

var _prefixLength int

func Spawn(path string) error {
	subPath := fmt.Sprintf("%s/%d", path, time.Now().Unix())

	err := os.MkdirAll(subPath, os.ModePerm)
	if err != nil {
		return err
	}

	fStats, err := os.Create(fmt.Sprintf("%s/stats.log", subPath))
	if err != nil {
		return err
	}

	fTrace, err := os.Create(fmt.Sprintf("%s/trace.log", subPath))
	if err != nil {
		fStats.Close()
		return err
	}

	_fStats = fStats
	_fTrace = fTrace

	_prefixLength = len(getFuncPrefix())

	return nil
}

func LogStats() {
	// TODO: write to `_fStats`
}

func LogTrace(format string, args ...interface{}) {
	timestamp := time.Now().Unix()
	function := getFuncName(2)
	message := fmt.Sprintf(format, args...)
	line := formatTrace(timestamp, function, message)

	_, err := _fTrace.WriteString(line)
	if err != nil {
		errMsg := fmt.Sprintf("cannot write to trace log file > %s", err)
		fmt.Printf(formatTrace(timestamp, getFuncName(1), errMsg))
		Kill()
		return
	}

	fmt.Printf(line)
}

func Kill() {
	if !IsAlive {
		return
	}

	IsAlive = false

	if _fStats != nil {
		_fStats.Close()
		_fStats = nil
	}

	if _fTrace != nil {
		_fTrace.Close()
		_fTrace = nil
	}
}
