package 炫彩WinApi类_test

import (
	"fmt"
	"github.com/888go/xcgui/common"
	"github.com/888go/xcgui/wapi"
	"github.com/888go/xcgui/xcc"
	"syscall"
	"testing"
)

func TestShellExecuteW(t *testing.T) {
	// 打开指定网址
	炫彩WinApi类.X对指定文件执行操作(0, "open", "https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shellexecutew", "", "", 炫彩常量类.SW_SHOWNORMAL)
	// 打开指定文件
	炫彩WinApi类.X对指定文件执行操作(0, "open", "C:\\Windows\\System32\\calc.exe", "", "", 炫彩常量类.SW_SHOWNORMAL)
}

func TestSHBrowseForFolderW(t *testing.T) {
	buf := make([]uint16, 260)
	bi := 炫彩WinApi类.BrowseInfoW{
		HwndOwner:      0,
		PidlRoot:       0,
		PszDisplayName: 炫彩工具类.Uint16SliceDataPtr(&buf),
		LpszTitle:      炫彩工具类.StrPtr("显示在对话框中树视图控件上方的文本"),
		UlFlags:        炫彩WinApi类.BIF_USENEWUI,
		Lpfn:           0,
		LParam:         0,
		IImage:         0,
	}
	var pszPath string
	ret := 炫彩WinApi类.X文件夹指针取实际路径(炫彩WinApi类.X对话框打开文件夹(&bi), &pszPath)
	fmt.Println(ret)
	fmt.Println("pszPath:", pszPath)                           // 用户选择的文件夹完整路径
	fmt.Println("PszDisplayName:", syscall.UTF16ToString(buf)) // 用户选择的文件夹的名称
}
