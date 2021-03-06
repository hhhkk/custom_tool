package tool

import (
	"errors"
	"github.com/h2non/filetype"
	"github.com/h2non/filetype/types"
	"github.com/hhhkk/custom_tool/log"
	"os"
	"path/filepath"
)

var cwdPath string

var hookCwdPathFunc func() string

func init() {
	cwdPath = os.Args[0]
	if dir, err := filepath.Abs(filepath.Dir(os.Args[0])); err == nil {
		cwdPath = dir
	} else {
		log.LibFatal(err)
	}
}

func GetFileType(path string) types.Type {
	if file, err := os.Open(path); err == nil {
		bytes := make([]byte, 256)
		file.Read(bytes)
		file.Read(bytes)
		if class, err := filetype.Get(bytes); err == nil {
			return class
		}
	}
	return types.Unknown
}

func HookCwdPath(hook func() string) {
	hookCwdPathFunc = hook
}

func IsExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		log.LibFatal(err)
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

func CreateFile(path string, success func(*os.File), fail func(error)) *os.File {
	if file, err := os.OpenFile(path,os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm); err == nil {
		if success != nil {
			defer file.Close()
			success(file)
		} else {
			return file
		}
	} else if fail != nil {
		fail(err)
	}
	return nil
}
func GetFileSize(path string) int64 {
	if info, err := os.Stat(path); err == nil {
		return info.Size()
	} else {
		return -1
	}
	return 0
}

func Open(path string, success func(file *os.File), fail func(err error)) *os.File {
	if file, err := os.Open(path); err == nil {
		if success != nil {
			defer file.Close()
			success(file)
		} else {
			return file
		}
	} else {
		log.LibE(err)
		if fail != nil {
			fail(err)
		}
	}
	return nil
}
