package main

import (
	"github.com/hhhkk/custom_tool/net"
	"github.com/hhhkk/custom_tool/tool"
)

func main() {
	net.SliceDownload("http://192.168.31.107/file/img/t1/photo-1572811185759-6e759e27dcd8",
		tool.GetCwdPath()+"/test.jpg")
}
