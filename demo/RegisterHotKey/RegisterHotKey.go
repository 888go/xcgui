// 注册热键, 全局的或窗口内的.
package main

import (
	"fmt"

	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/wapi"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xcc"
)

// 注册热键时填的id
const (
	ID_F3 = iota + 1
	ID_F4
)

var (
	a *炫彩App类.App
	w *炫彩窗口基类.Window
)

func main() {
	a = 炫彩App类.X创建(true)
	w = 炫彩窗口基类.X创建窗口(0, 0, 400, 300, "", 0, 炫彩常量类.Window_Style_Default)

	// 全局生效
	one()
	// 只在窗口内生效, 窗口内热键
	two()

	w.X显示(true)
	a.X运行()
	a.X退出()
}

// 全局生效
func one() {
	// 注册热键F3
	if !炫彩WinApi类.X键盘热键注册(w.X取HWND(), ID_F3, 0, 炫彩常量类.VK_F3) {
		fmt.Println("注册热键F3失败")
	}

	// 注册热键F4
	if !炫彩WinApi类.X键盘热键注册(w.X取HWND(), ID_F4, 0, 炫彩常量类.VK_F4) {
		fmt.Println("注册热键F4失败")
	}

	w.X线程_消息过程1(func(hWindow int, message uint32, wParam, lParam uint, pbHandled *bool) int {
		if message == uint32(炫彩常量类.WM_HOTKEY) {
			switch wParam {
			case ID_F3:
				fmt.Println("Event_WINDPROC1 F3键被按下")
			case ID_F4:
				fmt.Println("Event_WINDPROC1 F4键被按下")
			}
		}
		return 0
	})
}

// 只在窗口内生效, 窗口内热键
func two() {
	w.X线程_键盘按键消息1(func(hWindow int, wParam, lParam uint, pbHandled *bool) int {
		switch wParam {
		case 炫彩常量类.VK_F5:
			fmt.Println("Event_KEYDOWN1 F5键被按下")
		case 炫彩常量类.VK_F6:
			fmt.Println("Event_KEYDOWN1 F6键被按下")
		}
		return 0
	})
}
