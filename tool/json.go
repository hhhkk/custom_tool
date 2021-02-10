package tool

import (
	"bytes"
	"encoding/json"
	"github.com/hhhkk/custom_tool/log"
	"io"
	"io/ioutil"
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

func Encoding(obj interface{},data *[]byte) interface{} {
	if json.Unmarshal(*data,obj)==nil{
		return obj
	}else{
		return nil
	}
}

func EncodingByReader(obj interface{},data io.Reader)interface{} {
	if byteData,err := ioutil.ReadAll(data);err ==nil{
		return Encoding(obj,&byteData)
	}else{
		return nil
	}
}
