package wnd//bm:炫彩WinApi窗口类

import (
	"github.com/twgh/xcgui/wapi"
	"strings"
)

// SetTop 窗口_置顶.
//
//	@param hWnd 窗口真实句柄.
//	@param b 是否置顶.
//	@return bool
// ff:窗口置顶
// hWnd:窗口句柄
// b:是否置顶
func SetTop(hWnd uintptr, b bool) bool {
	hWndInsertAfter := wapi.HWND_TOPMOST
	if !b {
		hWndInsertAfter = wapi.HWND_NOTOPMOST
	}
	return wapi.SetWindowPos(hWnd, hWndInsertAfter, 0, 0, 0, 0, wapi.SWP_NOMOVE|wapi.SWP_NOSIZE)
}

// GetTitle 取窗口标题.
//
//	@param hWnd 窗口真实句柄.
//	@return string
// ff:窗口取标题
// hWnd:窗口句柄
func GetTitle(hWnd uintptr) string {
	dwSize := wapi.GetWindowTextLengthW(hWnd)
	if dwSize == 0 {
		return ""
	}
	dwSize++ // 必须算上空字符

	var title string
	wapi.GetWindowTextW(hWnd, &title, dwSize)
	return title
}

// GetHWND 取窗口句柄, 标题支持模糊.
//
//	@param className 窗口类名, 不支持模糊, 可空.
//	@param title 窗口标题, 可输入关键字, 支持模糊, 可空.
//	@return int 窗口真实句柄.
// ff:窗口模糊取句柄
// className:窗口类名
// title:窗口标题
func GetHWND(className, title string) uintptr {
	if className == "" && title == "" {
		return 0
	}
	var hWnd uintptr
	for {
		hWnd = wapi.FindWindowExW(0, hWnd, className, "")
		if hWnd != 0 {
			titleName := strings.ToLower(GetTitle(hWnd))
			if strings.Index(titleName, strings.ToLower(title)) != -1 {
				return hWnd
			}
		} else {
			break
		}
	}
	return 0
}
