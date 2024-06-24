package tests

import (
	"os"
	"strconv"

	"github.com/reatang/go-tapd/tapd"
)

// 抽象重复获取客户端的代码
func getClient() *tapd.Client {
	w, _ := strconv.Atoi(os.Getenv("TAPD_API_WorkspaceID"))
	return tapd.NewTAPDClient(tapd.WithBasicAuth(os.Getenv("TAPD_API_USER"), os.Getenv("TAPD_API_PASSWORD")), tapd.WithWorkspaceID(w))
}
