package 炫彩WinApi类_test

import (
	"fmt"
	"github.com/888go/xcgui/tf"
	"testing"
	
	"github.com/888go/xcgui/app"
	"github.com/888go/xcgui/wapi"
	"github.com/888go/xcgui/widget"
	"github.com/888go/xcgui/window"
	"github.com/888go/xcgui/xcc"
)

func TestGetDesktopWindow(t *testing.T) {
	fmt.Println(炫彩WinApi类.X窗口取桌面句柄())
}

func TestMessageBoxW(t *testing.T) {
	id := 炫彩WinApi类.X窗口消息框W(0, "context", "title", 炫彩WinApi类.MB_CanaelTryContinue|炫彩WinApi类.MB_IconInformation)
	switch id {
	case 炫彩WinApi类.ID_Cancel:
		fmt.Println("Cancel")
	case 炫彩WinApi类.ID_TryAgain:
		fmt.Println("TryAgain")
	case 炫彩WinApi类.ID_Continue:
		fmt.Println("Continue")
	default:
		fmt.Println(id)
	}
}

func TestFindWindowExW(t *testing.T) {
	fmt.Println(炫彩WinApi类.X窗口模糊搜索子窗口(0, 0, "", "任务管理器"))
	fmt.Println(炫彩WinApi类.X窗口模糊搜索子窗口(0, 0, "TaskManagerWindow", ""))
	fmt.Println(炫彩WinApi类.X窗口模糊搜索子窗口(0, 0, "TaskManagerWindow", "任务管理器"))
}

func TestFindWindowW(t *testing.T) {
	fmt.Println(炫彩WinApi类.X窗口取顶级句柄("", "任务管理器"))
	fmt.Println(炫彩WinApi类.X窗口取顶级句柄("TaskManagerWindow", ""))
	fmt.Println(炫彩WinApi类.X窗口取顶级句柄("TaskManagerWindow", "任务管理器"))
}

func TestGetWindowTextLengthW(t *testing.T) {
	// 取任务管理器的窗口标题长度
	fmt.Println(炫彩WinApi类.X窗口取标题长度(炫彩WinApi类.X窗口模糊搜索子窗口(0, 0, "TaskManagerWindow", "")))
}

func TestClientToScreen(t *testing.T) {
	tf.TFunc(func(a *炫彩App类.App, w *炫彩窗口基类.Window) {
		pt := 炫彩WinApi类.POINT{X: 0, Y: 0}
		炫彩WinApi类.X窗口取屏幕坐标(w.X取HWND(), &pt)
		fmt.Println(pt)
	})
}

func TestGetCursorPos(t *testing.T) {
	tf.TFunc(func(a *炫彩App类.App, w *炫彩窗口基类.Window) {
		炫彩组件类.X创建按钮(20, 50, 450, 300, "GetCursorPos", w.Handle).X事件_被单击(func(pbHandled *bool) int {
			var pt 炫彩WinApi类.POINT
			炫彩WinApi类.X鼠标取光标坐标(&pt)
			a.X消息框("GetCursorPos", fmt.Sprintf("x: %d  y: %d", pt.X, pt.Y), 炫彩常量类.MessageBox_Flag_Ok, w.X取HWND(), 炫彩常量类.Window_Style_Default)
			return 0
		})
	})
}

func TestIsWindow(t *testing.T) {
	tf.TFunc(func(a *炫彩App类.App, w *炫彩窗口基类.Window) {
		炫彩组件类.X创建按钮(20, 50, 100, 30, "IsWindow", w.Handle).X事件_被单击(func(pbHandled *bool) int {
			w.X消息框("IsWindow", fmt.Sprintf("IsWindow: %v", 炫彩WinApi类.X窗口句柄是否有效(w.X取HWND())), 炫彩常量类.MessageBox_Flag_Ok, 炫彩常量类.Window_Style_Default)
			return 0
		})
	})
}

func TestPostQuitMessage(t *testing.T) {
	tf.TFunc(func(a *炫彩App类.App, w *炫彩窗口基类.Window) {
		炫彩组件类.X创建按钮(20, 50, 200, 30, "PostQuitMessage", w.Handle).X事件_被单击(func(pbHandled *bool) int {
			a.X启用自动退出程序(false)
			w.X关闭()
			炫彩WinApi类.X结束(0)
			return 0
		})
	})
}

func TestSetForegroundWindow(t *testing.T) {
	tf.TFunc(func(a *炫彩App类.App, w *炫彩窗口基类.Window) {
		炫彩组件类.X创建按钮(20, 50, 200, 30, "SetForegroundWindow", w.Handle).X事件_被单击(func(pbHandled *bool) int {
			hwnd := 炫彩WinApi类.X窗口取顶级句柄("TaskManagerWindow", "")
			if !炫彩WinApi类.X窗口句柄是否有效(hwnd) {
				a.X弹框("提示", "请先打开任务管理器, 本按钮功能是把任务管理器窗口激活")
				return 0
			}
			炫彩WinApi类.X窗口激活(hwnd)
			return 0
		})
	})
}
