package tool

import (
	"time"
)
var cst *time.Location
func init() {
	cst, _ = time.LoadLocation("Asia/Shanghai")
}

func GetTime() time.Time {
	return time.Now().In(cst)
}
