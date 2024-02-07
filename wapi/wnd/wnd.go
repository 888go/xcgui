package 炫彩WinApi窗口类

import (
	"github.com/888go/xcgui/wapi"
	"strings"
)

// SetTop 窗口_置顶.
//
//	@param hWnd 窗口真实句柄.
//	@param b 是否置顶.
//	@return bool
func SetTop(hWnd uintptr, b bool) bool {
	hWndInsertAfter := 炫彩WinApi类.HWND_TOPMOST
	if !b {
		hWndInsertAfter = 炫彩WinApi类.HWND_NOTOPMOST
	}
	return 炫彩WinApi类.SetWindowPos(hWnd, hWndInsertAfter, 0, 0, 0, 0, 炫彩WinApi类.SWP_NOMOVE|炫彩WinApi类.SWP_NOSIZE)
}

// GetTitle 取窗口标题.
//
//	@param hWnd 窗口真实句柄.
//	@return string
func GetTitle(hWnd uintptr) string {
	dwSize := 炫彩WinApi类.GetWindowTextLengthW(hWnd)
	if dwSize == 0 {
		return ""
	}
	dwSize++ // 必须算上空字符

	var title string
	炫彩WinApi类.GetWindowTextW(hWnd, &title, dwSize)
	return title
}

// GetHWND 取窗口句柄, 标题支持模糊.
//
//	@param className 窗口类名, 不支持模糊, 可空.
//	@param title 窗口标题, 可输入关键字, 支持模糊, 可空.
//	@return int 窗口真实句柄.
func GetHWND(className, title string) uintptr {
	if className == "" && title == "" {
		return 0
	}
	var hWnd uintptr
	for {
		hWnd = 炫彩WinApi类.FindWindowExW(0, hWnd, className, "")
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
