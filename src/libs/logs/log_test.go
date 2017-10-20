package logs

import (
	"os"
	"testing"
)

func Test_Output(t *testing.T) {
	Debug("123", "sss")
	Info("123", "sss")
	Warn("123", "sss")
	//Error("123", "sss")

	var logger = &Logger{
		files: []*os.File{
			os.Stdout,
		},
		level:     _level_debug,
		calldepth: 4,
	}

	logger.Debug(12311)
}
