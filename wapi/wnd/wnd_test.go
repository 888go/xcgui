package 炫彩WinApi窗口类_test

import (
	"fmt"
	"github.com/888go/xcgui/wapi"
	"github.com/888go/xcgui/wapi/wnd"
	"testing"
)

func TestGetWindowTitle(t *testing.T) {
	hwnd := 炫彩WinApi类.FindWindowExW(0, 0, "", "任务管理器")
	fmt.Println(炫彩WinApi窗口类.GetTitle(hwnd))
}

func TestSetTop(t *testing.T) {
	hwnd := 炫彩WinApi类.FindWindowExW(0, 0, "", "任务管理器")
	fmt.Println(炫彩WinApi窗口类.SetTop(hwnd, true))
}

func TestGetWindowHWND(t *testing.T) {
	fmt.Println(炫彩WinApi窗口类.GetHWND("", "管理器"))
}
