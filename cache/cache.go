package cache

import (
	"fmt"
	"github.com/hhhkk/custom_tool/log"
	"github.com/hhhkk/custom_tool/tool"
	"io"
	"os"
)

var cacheDir string

func init() {
	cacheDir = tool.GetCwdPath() + "/cache/"
	if !tool.IsExist(cacheDir) {
		if err := os.Mkdir(cacheDir, os.ModePerm); err != nil {
			panic(err)
		}
	}
}

func Save(md5 string, data io.Reader) {
	path := cacheDir + md5
	if !tool.IsExist(path) {
		tool.CreateFile(path, func(file *os.File) {
			if _, err := io.Copy(file, data); err != nil {

				log.LibE(fmt.Errorf("save cache copy file fail"))
			}
		}, func(err error) {
			log.LibE(fmt.Errorf("save cache create file fail"))
		})
	} else {
		log.LibE(fmt.Errorf("save cache fail cache exist"))
	}
}

func Get(md5 string) *os.File {
	return tool.Open(cacheDir+md5, nil, nil)
}

func IsExist(md5 string) bool {
	return tool.IsExist(cacheDir + md5)
}
