package log

import "fmt"

var enableLibLog bool

func EnableLibraryLog() {
	enableLibLog = true
}

func IsEnableLibraryLog() bool {
	return enableLibLog
}

func DisableLibraryLog() {
	enableLibLog = false
}

func E(error ...error) {
	fmt.Println(error)
}

func Fatal(error ...error) {
	panic(error)
}

func LibE(error ...error) {
	if !enableLibLog {
		return
	}
	E(error...)
}

func LibFatal(error ...error) {
	if !enableLibLog {
		return
	}
	Fatal(error...)
}
