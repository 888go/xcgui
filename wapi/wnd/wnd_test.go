package 炫彩WinApi窗口类_test//bm:炫彩WinApi窗口类_test

import (
	"fmt"
	"github.com/888go/xcgui/wapi"
	"github.com/888go/xcgui/wapi/wnd"
	"testing"
)

func TestGetWindowTitle(t *testing.T) {
	hwnd := 炫彩WinApi类.X窗口模糊搜索子窗口(0, 0, "", "任务管理器")
	fmt.Println(炫彩WinApi窗口类.X窗口取标题(hwnd))
}

func TestSetTop(t *testing.T) {
	hwnd := 炫彩WinApi类.X窗口模糊搜索子窗口(0, 0, "", "任务管理器")
	fmt.Println(炫彩WinApi窗口类.X窗口置顶(hwnd, true))
}

func TestGetWindowHWND(t *testing.T) {
	fmt.Println(炫彩WinApi窗口类.X窗口模糊取句柄("", "管理器"))
}
