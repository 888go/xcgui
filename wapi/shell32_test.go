package wapi_test

import (
	"e.coding.net/gogit/go/xcgui/common"
	"e.coding.net/gogit/go/xcgui/wapi"
	"e.coding.net/gogit/go/xcgui/xcc"
	"fmt"
	"syscall"
	"testing"
)

func TestShellExecuteW(t *testing.T) {
	// 打开指定网址
	wapi.ShellExecuteW(0, "open", "https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shellexecutew", "", "", xcc.I常量_窗口形式_显示且激活_原尺寸位置)
	// 打开指定文件
	wapi.ShellExecuteW(0, "open", "C:\\Windows\\System32\\calc.exe", "", "", xcc.I常量_窗口形式_显示且激活_原尺寸位置)
}

func TestSHBrowseForFolderW(t *testing.T) {
	buf := make([]uint16, 260)
	bi := wapi.BrowseInfoW{
		HwndOwner:      0,
		PidlRoot:       0,
		PszDisplayName: common.Uint16SliceDataPtr(&buf),
		LpszTitle:      common.StrPtr("显示在对话框中树视图控件上方的文本"),
		UlFlags:        wapi.BIF_USENEWUI,
		Lpfn:           0,
		LParam:         0,
		IImage:         0,
	}
	var pszPath string
	ret := wapi.SHGetPathFromIDListW(wapi.SHBrowseForFolderW(&bi), &pszPath)
	fmt.Println(ret)
	fmt.Println("pszPath:", pszPath)                           // 用户选择的文件夹完整路径
	fmt.Println("PszDisplayName:", syscall.UTF16ToString(buf)) // 用户选择的文件夹的名称
}
