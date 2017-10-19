package log

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
	"time"
)

var logger *Logger = nil

func init() {
	var writers = []io.Writer{
		os.Stdout,
	}
	logger = &Logger{
		out:      io.MultiWriter(writers...),
		out_type: terminal,
	}
}

const (
	/*color format
	"\x1b[0;%dm%s\x1b[0m"
	*/
	color_text_black = iota + 30
	color_text_red
	color_text_green
	color_text_yellow
	color_text_blue
	color_text_magenta
	color_text_cyan
	color_text_white
)

const (
	/*out type*/
	terminal = 0
	file     = 1
)

type Logger struct {
	mu       sync.Mutex
	out      io.Writer
	out_type uint
}

func New(out io.Writer) *Logger {
	var l = &Logger{
		out:      out,
		out_type: file,
	}
	return l
}

func (l *Logger) Output(pretty string, prefix string, calldepth int, s string) error {
	var now = time.Now()
	l.mu.Lock()
	defer l.mu.Unlock()
	var str = fmt.Sprintf(pretty, Format(prefix, calldepth, now, s))
	var _, err = l.out.Write([]byte(str))
	return err
}

var Format = func(prefix string, calldepth int, t time.Time, s string) string {
	var _, file_name, line_number, ok = runtime.Caller(calldepth)
	if !ok {
		file_name = "???"
		line_number = 0
	}
	var buf bytes.Buffer

	buf.WriteString(prefix)
	var fstr = t.Format("2006-01-02 15:04:05.000000000")
	buf.WriteString(fstr)
	buf.Write([]byte(" "))
	buf.WriteString(file_name)
	buf.Write([]byte(":"))
	buf.WriteString(fmt.Sprintf("%d", line_number))
	buf.Write([]byte(" > "))
	buf.WriteString(s)

	return buf.String()
}

func color_format(color int) string {
	if runtime.GOOS == "windows" {
		return "%s"
	}

	var format = "\x1b[0;%dm%s\x1b[0m"
	switch color {
	case color_text_black,
		color_text_red,
		color_text_green,
		color_text_yellow,
		color_text_blue,
		color_text_magenta,
		color_text_cyan,
		color_text_white:
		return fmt.Sprintf(format, color, "%s")
	default:
		return "%s"
	}
}

func Debug(v interface{}) {
	if logger != nil {
		var pretty = "%s"
		if logger.out_type == terminal {
			pretty = color_format(color_text_yellow)
		}

		logger.Output(pretty, "[DEBUG]▸ ", 3, fmt.Sprintln(v))
	}
}

func Info(v interface{}) {
	if logger != nil {
		var pretty = "%s"
		if logger.out_type == terminal {
			pretty = color_format(color_text_white)
		}

		logger.Output(pretty, "[INFO]▸ ", 3, fmt.Sprintln(v))
	}
}

func Error(v interface{}) {
	if logger != nil {
		var pretty = "%s"
		if logger.out_type == terminal {
			pretty = color_format(color_text_red)
		}

		logger.Output(pretty, "[ERROR]▸ ", 3, fmt.Sprintln(v))
	}
}
