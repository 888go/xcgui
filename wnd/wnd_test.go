package wnd_test

import (
	"e.coding.net/gogit/go/xcgui/wapi"
	"e.coding.net/gogit/go/xcgui/wnd"
	"fmt"
	"testing"
)

func TestGetWindowTitle(t *testing.T) {
	hwnd := wapi.FindWindowExW(0, 0, "", "任务管理器")
	fmt.Println(wnd.GetTitle(hwnd))
}

func TestSetTop(t *testing.T) {
	hwnd := wapi.FindWindowExW(0, 0, "", "任务管理器")
	fmt.Println(wnd.SetTop(hwnd, true))
}

func TestGetWindowHWND(t *testing.T) {
	fmt.Println(wnd.GetHWND("", "管理器"))
}
