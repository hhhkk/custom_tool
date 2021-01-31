package tool

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func GetMd5Value(value string) string {
	nameSpase := md5.New()
	io.WriteString(nameSpase, value)
	hexData := nameSpase.Sum(nil)
	return hex.EncodeToString(hexData[0:len(hexData)])
}

