package logs

import (
	"fmt"
	"runtime"
	"strings"
)

func getFuncPrefix() string {
	return strings.Replace(getFuncName(1), "logs.getFuncPrefix", "", -1)
}

func getFuncName(skipFrames int) string {
	pc, _, _, _ := runtime.Caller(skipFrames)
	fn := runtime.FuncForPC(pc)
	if fn != nil {
		name := fn.Name()
		if name[0] == '_' {
			return name[_prefixLength:]
		}

		return name
	}

	return "<unknown>"
}

func formatTrace(timestamp int64, function, message string) string {
	return fmt.Sprintf("@%d [%s] %s\n", timestamp, function, message)
}
