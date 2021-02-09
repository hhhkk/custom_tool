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
}

func Save(md5 string, data io.Reader) {
	path := cacheDir + md5
	if !tool.IsExist(path) {
		tool.CreateFile(path, func(file *os.File) {
			if _, err := io.Copy(file, data); err != nil {
				log.E(fmt.Errorf("save cache copy file fail"))
			}
		}, func(err error) {
			log.E(fmt.Errorf("save cache create file fail"))
		})
	} else {
		log.E(fmt.Errorf("save cache fail cache exist"))
	}
}

func Get(md5 string) io.Reader {
	return tool.Open(cacheDir+md5, nil, nil)
}

func IsExist(md5 string) bool {
	return tool.IsExist(cacheDir + md5)
}
