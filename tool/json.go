package tool

import (
	"bytes"
	"encoding/json"
	"custom_tool/log"
)

func ToString(obj *interface{}) *string {
	if obj == nil {
		return nil
	} else {
		if data, err := json.Marshal(*obj); err == nil {
			jsonStr := bytes.NewBuffer(data).String()
			return &jsonStr
		}else{
			log.Fatal(err)
			return nil
		}
	}
}
