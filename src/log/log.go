package log

import (
	"fmt"
	"time"
	"runtime"
	"os"
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
		if false {
			fmt.Println(lastTs)
		}
		return fmt.Sprintf("%20s:%-3d %s%d%s%-6s%s ", file, line, pre, color, "m", t, end) + text
	}
}

var colors = []brush{
	newBrush(35, "Emergency"), // Emergency          white
	newBrush(36, "Alert"),     // Alert              cyan
	newBrush(37, "io"),        // Critical           magenta
	newBrush(31, "Error"),     // Error              red
	newBrush(33, "Warn"),      // Warning            yellow
	newBrush(32, "Notice"),    // Notice             green
	newBrush(34, "Info"),      // Informational      blue
	newBrush(44, "Debug"),     // Debug              Background blue
}

// any

func Debug(v ...interface{}) {
	str := fmt.Sprintln(v...)
	fmt.Print(colors[LevelNotice](str))
}

func Debugf(format string, v ...interface{}) {
	str := fmt.Sprintf(format, v...)
	fmt.Println(colors[LevelNotice](str))
}

func Info(v ...interface{}) {
	str := fmt.Sprintln(v...)
	fmt.Print(colors[LevelInformational](str))
}

func Infof(format string, v ...interface{}) {
	str := fmt.Sprintf(format, v...)
	fmt.Println(colors[LevelInformational](str))
}

func Change(v ...interface{}) {
	str := fmt.Sprintln(v...)
	fmt.Print(colors[LevelCritical](str))
}

func Changef(format string, v ...interface{}) {
	str := fmt.Sprintf(format, v...)
	fmt.Println(colors[LevelCritical](str))
}

func Warn(v ...interface{}) {
	str := fmt.Sprintln(v...)
	fmt.Print(colors[LevelWarning](str))
}

func Warnf(format string, v ...interface{}) {
	str := fmt.Sprintf(format, v...)
	fmt.Println(colors[LevelWarning](str))
}

func Error(v ...interface{}) {
	str := fmt.Sprintln(v...)
	fmt.Print(colors[LevelError](str))
}

func Errorf(format string, v ...interface{}) {
	str := fmt.Sprintf(format, v...)
	fmt.Println(colors[LevelError](str))
}

func Fatal(v ...interface{}) {
	str := fmt.Sprintln(v...)
	fmt.Print(colors[LevelEmergency](str))
}

func Fatalf(format string, v ...interface{}) {
	str := fmt.Sprintf(format, v...)
	fmt.Println(colors[LevelEmergency](str))
	os.Exit(1)
}
