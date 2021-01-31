package main

import (
	"github.com/hhhkk/custom_tool/net"
	"github.com/hhhkk/custom_tool/tool"
)

func main() {
	net.SliceDownload("https://images.unsplash.com/photo-1572811185759-6e759e27dcd8?ixid=MXwxMjA3fDB8MXxhbGx8ODMyMDl8fHx8fHwyfA&ixlib=rb-1.2.1",
		tool.GetCwdPath()+"/test.jpg")
}
