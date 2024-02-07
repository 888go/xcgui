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
func X窗口置顶(窗口句柄 uintptr, 是否置顶 bool) bool {
	hWndInsertAfter := 炫彩WinApi类.HWND_TOPMOST
	if !是否置顶 {
		hWndInsertAfter = 炫彩WinApi类.HWND_NOTOPMOST
	}
	return 炫彩WinApi类.X窗口设置位置(窗口句柄, hWndInsertAfter, 0, 0, 0, 0, 炫彩WinApi类.SWP_NOMOVE|炫彩WinApi类.SWP_NOSIZE)
}

// GetTitle 取窗口标题.
//
//	@param hWnd 窗口真实句柄.
//	@return string
func X窗口取标题(窗口句柄 uintptr) string {
	dwSize := 炫彩WinApi类.X窗口取标题长度(窗口句柄)
	if dwSize == 0 {
		return ""
	}
	dwSize++ // 必须算上空字符

	var title string
	炫彩WinApi类.X窗口取标题(窗口句柄, &title, dwSize)
	return title
}

// GetHWND 取窗口句柄, 标题支持模糊.
//
//	@param className 窗口类名, 不支持模糊, 可空.
//	@param title 窗口标题, 可输入关键字, 支持模糊, 可空.
//	@return int 窗口真实句柄.
func X窗口模糊取句柄(窗口类名, 窗口标题 string) uintptr {
	if 窗口类名 == "" && 窗口标题 == "" {
		return 0
	}
	var hWnd uintptr
	for {
		hWnd = 炫彩WinApi类.X窗口模糊搜索子窗口(0, hWnd, 窗口类名, "")
		if hWnd != 0 {
			titleName := strings.ToLower(X窗口取标题(hWnd))
			if strings.Index(titleName, strings.ToLower(窗口标题)) != -1 {
				return hWnd
			}
		} else {
			break
		}
	}
	return 0
}
