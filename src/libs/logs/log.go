package logs

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

const (
	/*color format
	"\x1b[0;%dm%s\x1b[0m"
	*/
	_color_text_black = iota + 30
	_color_text_red
	_color_text_green
	_color_text_yellow
	_color_text_blue
	_color_text_magenta
	_color_text_cyan
	_color_text_white
)

const (
	_level_debug = 0
	_level_info  = 1
	_level_warn  = 2
	_level_error = 3
)

var logger *Logger = nil

func init() {
	logger = &Logger{
		files: []*os.File{
			os.Stdout,
		},
		level:     _level_debug,
		calldepth: 5,
	}
}

type Logger struct {
	mux       sync.Mutex
	files     []*os.File
	level     int
	calldepth int
}

func NewLogger(f []*os.File, level int, calldepth int) *Logger {
	return &Logger{
		files:     f,
		level:     level,
		calldepth: calldepth,
	}
}

func (l *Logger) Writer(color int, s string) (int, error) {
	for i, f := range l.files {
		var fd = f.Fd()
		var name = f.Name()

		var pretty_fmt = "%s"
		if _color_text_black <= color && color <= _color_text_white &&
			((fd == 1 && name == "/dev/stdout") || (fd == 2 && name == "/dev/stderr")) {
			pretty_fmt = fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", color, "%s")
		}

		var ss = fmt.Sprintf(pretty_fmt, s)
		var _, err = f.Write([]byte(ss))
		if err != nil {
			return i + 1, err
		}
	}

	return len(l.files), nil
}

func (l *Logger) Output(color int, prefix string, s string) error {
	var now = time.Now()
	l.mux.Lock()
	defer l.mux.Unlock()
	var str = Format(prefix, now, "", l.calldepth, s)
	var _, err = l.Writer(color, str)
	return err
}

var Format = func(prefix string, time time.Time, timefmt string, calldepth int, s string) string {
	var p = GenPrefixInfo(prefix)
	var t = GenTimeInfo(time, timefmt)
	var fl = GenFileAndLineNumInfo(calldepth)

	var str = fmt.Sprintf("%s %s %s> %s", p, t, fl, s)
	return str
}

var GenPrefixInfo = func(s string) (ss string) {
	ss = s
	return
}

var GenTimeInfo = func(t time.Time, fmt string) string {
	var default_fmt = "2006-01-02 15:04:05.000"
	if fmt != "" {
		default_fmt = fmt
	}
	var fstr = t.Format(default_fmt)
	return fstr
}

var GenFileAndLineNumInfo = func(calldepth int) string {
	var _, file_name, line_number, ok = runtime.Caller(calldepth)
	if !ok {
		file_name = "???"
		line_number = 0
	}

	return fmt.Sprintf("%s:%d", file_name, line_number)
}

func (l *Logger) Debug(v ...interface{}) {
	if l.level <= _level_debug {
		var s = fmt.Sprintln(v...)
		l.Output(_color_text_green, "[DEBUG]", s)
	}
}

func (l *Logger) Info(v ...interface{}) {
	if l.level <= _level_info {
		var s = fmt.Sprintln(v...)
		l.Output(_color_text_white, "[INFO]", s)
	}
}

func (l *Logger) Warn(v ...interface{}) {
	if l.level <= _level_warn {
		var s = fmt.Sprintln(v...)
		l.Output(_color_text_yellow, "[WARN]", s)
	}
}

func (l *Logger) Error(v ...interface{}) {
	var s = fmt.Sprintln(v...)
	l.Output(_color_text_red, "[ERROR]", s)
	panic(s)
}

func Debug(v ...interface{}) {
	logger.Debug(v...)
}

func Info(v ...interface{}) {
	logger.Info(v...)
}

func Warn(v ...interface{}) {
	logger.Warn(v...)
}

func Error(v ...interface{}) {
	logger.Error(v...)
}
