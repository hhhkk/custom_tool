package tool

import (
	"github.com/hhhkk/custom_tool/log"
	"errors"
	"os"
	"path/filepath"
)

var cwdPath string

var hookCwdPathFunc func() string

func HookCwdPath(hook func() string) {
	hookCwdPathFunc = hook
}

func init() {
	cwdPath = os.Args[0]
	if dir, err := filepath.Abs(filepath.Dir(os.Args[0])); err == nil {
		cwdPath = dir
	} else {
		log.Fatal(err)
	}
}

func IsExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		log.Fatal(err)
		return false
	}
}

func WriteFile(path string, data []byte) error {
	if len(data) <= 0 {
		return errors.New("writeFile size == 0")
	}
	if file, err := os.Create(path); err == nil {
		defer file.Close()
		file.Write(data)
		file.Sync()
		return nil
	} else {
		return err
	}
}

func GetCwdPath() string {
	if hookCwdPathFunc != nil {
		return hookCwdPathFunc()
	} else {
		return cwdPath
	}
}

func CreateFile(path string, success func(*os.File), fail func(error)) {
	if file, err := os.Create(path); err == nil {
		defer file.Close()
		if success != nil {
			success(file)
		}
	} else if fail != nil {
		fail(err)
	}
}
func GetFileSize(path string) int64 {
	if info, err := os.Stat(path); err == nil {
		return info.Size()
	} else {
		return -1
	}
	return 0
}

func Open(path string, success func(file *os.File), fail func(err error)) {
	if file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm); err == nil {
		if success != nil {
			success(file)
		}
	} else {
		log.E(err)
		if fail != nil {
			fail(err)
		}
	}
}
