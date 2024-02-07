package 炫彩WinApi类_test

import (
	"fmt"
	"github.com/888go/xcgui/common"
	"github.com/888go/xcgui/wapi"
	"strings"
	"syscall"
	"testing"
	"unsafe"
)

// 多个过滤器, 打开单个文件.
func TestGetOpenFileNameW(t *testing.T) {
	lpstrFilter := strings.Join([]string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}, 炫彩WinApi类.NULL) + 炫彩WinApi类.NULL2

	lpstrFile := make([]uint16, 260)
	lpstrFileTitle := make([]uint16, 260)

	ofn := 炫彩WinApi类.OpenFileNameW{
		LStructSize:       76,
		HwndOwner:         0,
		HInstance:         0,
		LpstrFilter:       炫彩工具类.StringToUint16Ptr(lpstrFilter),
		LpstrCustomFilter: nil,
		NMaxCustFilter:    0,
		NFilterIndex:      1,
		LpstrFile:         &lpstrFile[0],
		NMaxFile:          260,
		LpstrFileTitle:    &lpstrFileTitle[0],
		NMaxFileTitle:     260,
		LpstrInitialDir:   炫彩工具类.StrPtr("D:"),
		LpstrTitle:        炫彩工具类.StrPtr("打开文件"),
		Flags:             炫彩WinApi类.OFN_PATHMUTEXIST, // 用户只能键入有效的路径和文件名
		NFileOffset:       0,
		NFileExtension:    0,
		LpstrDefExt:       0,
		LCustData:         0,
		LpfnHook:          0,
		LpTemplateName:    0,
	}
	ofn.LStructSize = uint32(unsafe.Sizeof(ofn))
	ret := 炫彩WinApi类.X创建打开对话框(&ofn)
	fmt.Println(ret)
	if !ret {
		return
	}
	fmt.Println("完整文件路径:", syscall.UTF16ToString(lpstrFile))
	fmt.Println("文件名:", syscall.UTF16ToString(lpstrFileTitle))
}

// 多个过滤器, 打开多个文件.
func TestGetOpenFileNameW_2(t *testing.T) {
	lpstrFilter := strings.Join([]string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}, 炫彩WinApi类.NULL) + 炫彩WinApi类.NULL2

	lpstrFile := make([]uint16, 512)

	ofn := 炫彩WinApi类.OpenFileNameW{
		LStructSize:       76,
		HwndOwner:         0,
		HInstance:         0,
		LpstrFilter:       炫彩工具类.StringToUint16Ptr(lpstrFilter),
		LpstrCustomFilter: nil,
		NMaxCustFilter:    0,
		NFilterIndex:      1,
		LpstrFile:         &lpstrFile[0],
		NMaxFile:          512,
		LpstrFileTitle:    nil,
		NMaxFileTitle:     0,
		LpstrInitialDir:   炫彩工具类.StrPtr("D:"),
		LpstrTitle:        炫彩工具类.StrPtr("打开文件(可选多个)"),
		Flags:             炫彩WinApi类.OFN_ALLOWMULTISELECT | 炫彩WinApi类.OFN_EXPLORER | 炫彩WinApi类.OFN_PATHMUTEXIST, // 允许文件多选 | 使用新界面 | 用户只能键入有效的路径和文件名
		NFileOffset:       0,
		NFileExtension:    0,
		LpstrDefExt:       0,
		LCustData:         0,
		LpfnHook:          0,
		LpTemplateName:    0,
	}
	ofn.LStructSize = uint32(unsafe.Sizeof(ofn))
	ret := 炫彩WinApi类.X创建打开对话框(&ofn)
	fmt.Println(ret)
	if !ret {
		return
	}
	s := 炫彩工具类.Uint16SliceToStringSlice(lpstrFile)
	fmt.Println("选择的文件个数:", len(s)-1) // -1是因为切片中第一个元素是文件目录, 不是文件
	fmt.Println("文件目录:", lpstrFile[0])
	fmt.Println("文件名:", s)
}

// 多个过滤器, 保存文件.
func TestGetSaveFileNameW(t *testing.T) {
	lpstrFilter := strings.Join([]string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}, 炫彩WinApi类.NULL) + 炫彩WinApi类.NULL2

	lpstrFile := make([]uint16, 260)
	lpstrFileTitle := make([]uint16, 260)

	ofn := 炫彩WinApi类.OpenFileNameW{
		LStructSize:       76,
		HwndOwner:         0,
		HInstance:         0,
		LpstrFilter:       炫彩工具类.StringToUint16Ptr(lpstrFilter),
		LpstrCustomFilter: nil,
		NMaxCustFilter:    0,
		NFilterIndex:      1,
		LpstrFile:         &lpstrFile[0],
		NMaxFile:          260,
		LpstrFileTitle:    &lpstrFileTitle[0],
		NMaxFileTitle:     260,
		LpstrInitialDir:   炫彩工具类.StrPtr("D:"),
		LpstrTitle:        炫彩工具类.StrPtr("保存文件"),
		Flags:             炫彩WinApi类.OFN_OVERWRITEPROMPT | 炫彩WinApi类.OFN_PATHMUTEXIST | 炫彩WinApi类.OFN_PATHMUTEXIST, // 如果所选文件已存在，则使“另存为”对话框生成一个消息框。用户必须确认是否覆盖文件。| 检测文件路径是否合法
		NFileOffset:       0,
		NFileExtension:    0,
		LpstrDefExt:       炫彩工具类.StrPtr("txt"), // 如果用户没有输入文件扩展名, 则默认使用这个
		LCustData:         0,
		LpfnHook:          0,
		LpTemplateName:    0,
	}
	ofn.LStructSize = uint32(unsafe.Sizeof(ofn))
	ret := 炫彩WinApi类.X创建保存对话框(&ofn)
	fmt.Println(ret)
	if !ret {
		return
	}
	fmt.Println("完整文件路径:", syscall.UTF16ToString(lpstrFile))
	fmt.Println("文件名:", syscall.UTF16ToString(lpstrFileTitle))
}

// 打开选择颜色界面
func TestChooseColorW(t *testing.T) {
	var lpCustColors [16]uint32
	cc := 炫彩WinApi类.ChooseColor{
		LStructSize:    36,
		HwndOwner:      0,
		HInstance:      0,
		RgbResult:      0,
		LpCustColors:   &lpCustColors[0],
		Flags:          炫彩WinApi类.CC_FULLOPEN, // 默认打开自定义颜色
		LCustData:      0,
		LpfnHook:       0,
		LpTemplateName: 0,
	}
	cc.LStructSize = uint32(unsafe.Sizeof(cc))
	ret := 炫彩WinApi类.X创建颜色对话框(&cc)
	fmt.Println(ret)
	if !ret {
		return
	}
	fmt.Println(cc.RgbResult) // rgb颜色
	fmt.Println(lpCustColors) // 如果你添加了自定义颜色, 会保存在这个数组里面, 然后只要这个数组还在, 再次打开选择颜色界面时, 之前添加的自定义颜色还会存在
}
