package log

import (
	"fmt"
	"time"
	"runtime"
)

// RFC5424 log message levels.
const (
	LevelEmergency     = iota
	LevelAlert
	LevelCritical
	LevelError
	LevelWarning
	LevelNotice
	LevelInformational
	LevelDebug
)

type brush func(string) string

func newBrush(color int, t string) brush {
	pre := "\033["
	end := "\033[0m"
	return func(text string) string {
		var file string
		var line int
		_, file, line, ok := runtime.Caller(2)
		if !ok {
			file = "???"
			line = 0
		}
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
				break
			}
		}
		file = short

		LastQueryTime := time.Now().Add(time.Hour * -24)
		lastTs := LastQueryTime.Format("2006-01-02 15:04:05")
		return fmt.Sprintf("%s %s:%d %s%d%s[%s]%s ", lastTs, file, line, pre, color, "m", t, end) + text
	}
}

var colors = []brush{
	newBrush(37, "Emergency"), // Emergency          white
	newBrush(36, "Alert"),     // Alert              cyan
	newBrush(35, "Critical"),  // Critical           magenta
	newBrush(31, "Err"),       // Error              red
	newBrush(33, "Warn"),      // Warning            yellow
	newBrush(32, "Notice"),    // Notice             green
	newBrush(34, "Info"),      // Informational      blue
	newBrush(44, "Debug"),     // Debug              Background blue
}

// any

func Info(v ...interface{}) {
	str := fmt.Sprint(v...)
	fmt.Println(colors[LevelInformational](str))
}

func Infof(format string, v ...interface{}) {
	str := fmt.Sprintf(format, v...)
	fmt.Println(colors[LevelInformational](str))
}

func Warn(v ...interface{}) {
	str := fmt.Sprint(v...)
	fmt.Println(colors[LevelWarning](str))
}

func Warnf(format string, v ...interface{}) {
	str := fmt.Sprintf(format, v...)
	fmt.Println(colors[LevelWarning](str))
}

func Error(v ...interface{}) {
	str := fmt.Sprint(v...)
	fmt.Println(colors[LevelError](str))
}

func Errorf(format string, v ...interface{}) {
	str := fmt.Sprintf(format, v...)
	fmt.Println(colors[LevelError](str))
}
