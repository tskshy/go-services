package log

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"libs/flag"
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

type LoggerConfLvl0 struct {
	Root *LoggerConfLvl1 `json:"root-log"`
}

type LoggerConfLvl1 struct {
	Level  string `json:"level"`
	Output string `json:"output"`
}

var logger *Logger = nil

func level_num(s string) int {
	var ss = strings.ToLower(s)
	switch ss {
	case "debug":
		return _level_debug
	case "info":
		return _level_info
	case "warn":
		return _level_warn
	case "error":
		return _level_error
	default:
		return _level_info
	}
}

func init() {
	var path = flag.Parse("log", "") //log conf file path
	if path == "" {
		/*没有传递日志配置文件，走默认设置*/
		logger = &Logger{
			files: []*os.File{
				os.Stdout,
			},
			level:     _level_debug,
			calldepth: 5,
		}
		return
	}

	var cfg_file, err_cfg = os.Open(path)
	if err_cfg != nil {
		panic(err_cfg)
	}

	var b, err_b = ioutil.ReadAll(cfg_file)
	if err_b != nil {
		panic(err_b)
	}

	var conf LoggerConfLvl0
	var err_conf = json.Unmarshal(b, &conf)
	if err_conf != nil {
		panic(err_conf)
	}

	if conf.Root == nil {
		panic("root-log -> nil")
	}

	var level = level_num(conf.Root.Level)
	var output = conf.Root.Output

	if output == "" {
		logger = &Logger{
			files: []*os.File{
				os.Stdout,
			},
			level:     level,
			calldepth: 5,
		}
		return
	} else {
		var f, err = os.OpenFile(output, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			panic(err)
		}
		logger = &Logger{
			files: []*os.File{
				f,
			},
			level:     level,
			calldepth: 5,
		}
		return
	}
	return
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

	var str = fmt.Sprintf("%s %s %s ▸ %s", p, t, fl, s)
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

	for i := len(file_name) - 1; i > 0; i-- {
		if file_name[i] == '/' {
			file_name = file_name[i+1:]
			break
		}
	}

	return fmt.Sprintf("%s:%d", file_name, line_number)
}

func (l *Logger) Debug(v ...interface{}) {
	if l.level <= _level_debug {
		var s = fmt.Sprintln(v...)
		var _ = l.Output(_color_text_green, "[DEBUG]", s)
	}
}

func (l *Logger) Info(v ...interface{}) {
	if l.level <= _level_info {
		var s = fmt.Sprintln(v...)
		var _ = l.Output(_color_text_white, "[INFO]", s)
	}
}

func (l *Logger) Warn(v ...interface{}) {
	if l.level <= _level_warn {
		var s = fmt.Sprintln(v...)
		var _ = l.Output(_color_text_yellow, "[WARN]", s)
	}
}

func (l *Logger) Error(v ...interface{}) {
	var s = fmt.Sprintln(v...)
	var _ = l.Output(_color_text_red, "[ERROR]", s)
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
