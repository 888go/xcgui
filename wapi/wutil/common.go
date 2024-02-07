package 炫彩WinApi工具类

import (
	"github.com/888go/xcgui/common"
	"github.com/888go/xcgui/wapi"
	"github.com/888go/xcgui/xc"
	"path/filepath"
	"strings"
	"syscall"
	"unsafe"
)

// GetDropFiles 获取拖放进来的文件.
//
// hDropInfo 拖放信息句柄.
func GetDropFiles(hDropInfo uintptr) []string {
	var filePath string
	files := make([]string, 0)
	var i uint32
	for {
		length := 炫彩WinApi类.DragQueryFileW(hDropInfo, i, &filePath, 260)
		if length == 0 { // 返回值为0说明已经检索完所有拖放进来的文件了.
			break
		}
		files = append(files, filePath)
		i++ // 索引+1检索下一个文件
	}
	炫彩WinApi类.DragFinish(hDropInfo)
	return files
}

// OpenDir 打开文件夹.
//
//	@param hParent 炫彩窗口句柄, 可为0.
//	@return string 返回选择的文件夹完整路径.
func OpenDir(hParent int) string {
	buf := make([]uint16, 260)
	var hwnd uintptr
	if hParent > 0 {
		hwnd = xc.XWnd_GetHWND(hParent)
	}
	bi := 炫彩WinApi类.BrowseInfoW{
		HwndOwner:      hwnd,
		PidlRoot:       0,
		PszDisplayName: 炫彩工具类.Uint16SliceDataPtr(&buf),
		LpszTitle:      炫彩工具类.StrPtr("请选择目录"),
		UlFlags:        炫彩WinApi类.BIF_USENEWUI,
		Lpfn:           0,
		LParam:         0,
		IImage:         0,
	}
	var pszPath string
	炫彩WinApi类.SHGetPathFromIDListW(炫彩WinApi类.SHBrowseForFolderW(&bi), &pszPath)
	return pszPath
}

// OpenFile 打开单个文件.
//
//	@param hParent 炫彩窗口句柄, 可为0.
//	@param filters 过滤器数组, 两个成员为一个过滤器, 前面是过滤器描述, 后面是过滤器类型. 填nil则不显示任何过滤器. 例: []string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}
//	@param defaultDir 初始目录, 即默认打开的目录.
//	@return string 返回文件完整路径.
func OpenFile(hParent int, filters []string, defaultDir string) string {
	var hwnd uintptr
	if hParent > 0 {
		hwnd = xc.XWnd_GetHWND(hParent)
	}
	// 拼接过滤器
	var LpstrFilter *uint16 = nil
	if len(filters) > 0 {
		LpstrFilter = 炫彩工具类.StringToUint16Ptr(strings.Join(filters, 炫彩WinApi类.NULL) + 炫彩WinApi类.NULL2)
	}

	lpstrFile := make([]uint16, 260)
	ofn := 炫彩WinApi类.OpenFileNameW{
		LStructSize:       76,
		HwndOwner:         hwnd,
		HInstance:         0,
		LpstrFilter:       LpstrFilter,
		LpstrCustomFilter: nil,
		NMaxCustFilter:    0,
		NFilterIndex:      1,
		LpstrFile:         &lpstrFile[0],
		NMaxFile:          260,
		LpstrFileTitle:    nil,
		NMaxFileTitle:     0,
		LpstrInitialDir:   炫彩工具类.StrPtr(defaultDir),
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
	if !炫彩WinApi类.GetOpenFileNameW(&ofn) {
		return ""
	}
	return syscall.UTF16ToString(lpstrFile)
}

// OpenFiles 打开多个文件.
//
//	@param hParent 炫彩窗口句柄, 可为0.
//	@param filters 过滤器数组, 两个成员为一个过滤器, 前面是过滤器描述, 后面是过滤器类型. 填nil则不显示任何过滤器. 例: []string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}
//	@param defaultDir 初始目录, 即默认打开的目录.
//	@return string 返回文件完整路径数组.
func OpenFiles(hParent int, filters []string, defaultDir string) []string {
	var hwnd uintptr
	if hParent > 0 {
		hwnd = xc.XWnd_GetHWND(hParent)
	}
	// 拼接过滤器
	var LpstrFilter *uint16 = nil
	if len(filters) > 0 {
		LpstrFilter = 炫彩工具类.StringToUint16Ptr(strings.Join(filters, 炫彩WinApi类.NULL) + 炫彩WinApi类.NULL2)
	}

	lpstrFile := make([]uint16, 512)
	ofn := 炫彩WinApi类.OpenFileNameW{
		LStructSize:       76,
		HwndOwner:         hwnd,
		HInstance:         0,
		LpstrFilter:       LpstrFilter,
		LpstrCustomFilter: nil,
		NMaxCustFilter:    0,
		NFilterIndex:      1,
		LpstrFile:         &lpstrFile[0],
		NMaxFile:          512,
		LpstrFileTitle:    nil,
		NMaxFileTitle:     0,
		LpstrInitialDir:   炫彩工具类.StrPtr(defaultDir),
		LpstrTitle:        炫彩工具类.StrPtr("打开文件"),
		Flags:             炫彩WinApi类.OFN_ALLOWMULTISELECT | 炫彩WinApi类.OFN_EXPLORER | 炫彩WinApi类.OFN_PATHMUTEXIST, // 允许文件多选 | 使用新界面 | 用户只能键入有效的路径和文件名
		NFileOffset:       0,
		NFileExtension:    0,
		LpstrDefExt:       0,
		LCustData:         0,
		LpfnHook:          0,
		LpTemplateName:    0,
	}
	ofn.LStructSize = uint32(unsafe.Sizeof(ofn))
	if !炫彩WinApi类.GetOpenFileNameW(&ofn) {
		return nil
	}

	slice := 炫彩工具类.Uint16SliceToStringSlice(lpstrFile)
	if len(slice) < 2 {
		return nil
	}

	dir := slice[0]
	var s []string
	for i := 1; i < len(slice); i++ {
		s = append(s, filepath.Join(dir, slice[i]))
	}
	return s
}

// SaveFile 保存文件.
//
//	@param hParent 炫彩窗口句柄, 可为0.
//	@param filters 过滤器数组, 两个成员为一个过滤器, 前面是过滤器描述, 后面是过滤器类型. 填nil则不显示任何过滤器. 例: []string{"Text Files(*txt)", "*.txt", "All Files(*.*)", "*.*"}
//	@param defaultDir 初始目录, 即默认打开的目录.
//	@param defaultFileName 默认文件名.
//	@return string 返回文件完整路径.
func SaveFile(hParent int, filters []string, defaultDir, defaultFileName string) string {
	var hwnd uintptr
	if hParent > 0 {
		hwnd = xc.XWnd_GetHWND(hParent)
	}
	// 拼接过滤器
	var lpstrFilter *uint16 = nil
	if len(filters) > 0 {
		lpstrFilter = 炫彩工具类.StringToUint16Ptr(strings.Join(filters, 炫彩WinApi类.NULL) + 炫彩WinApi类.NULL2)
	}

	var lpstrFile *uint16 = nil
	if defaultFileName != "" {
		lpstrFile = 炫彩工具类.StringToUint16Ptr(strings.ReplaceAll(defaultFileName, " ", ""))
	} else {
		lpstrFile = &make([]uint16, 260)[0]
	}
	ofn := 炫彩WinApi类.OpenFileNameW{
		LStructSize:       76,
		HwndOwner:         hwnd,
		HInstance:         0,
		LpstrFilter:       lpstrFilter,
		LpstrCustomFilter: nil,
		NMaxCustFilter:    0,
		NFilterIndex:      1,
		LpstrFile:         lpstrFile,
		NMaxFile:          260,
		LpstrFileTitle:    nil,
		NMaxFileTitle:     0,
		LpstrInitialDir:   炫彩工具类.StrPtr(defaultDir),
		LpstrTitle:        炫彩工具类.StrPtr("保存文件"),
		Flags:             炫彩WinApi类.OFN_OVERWRITEPROMPT | 炫彩WinApi类.OFN_PATHMUTEXIST | 炫彩WinApi类.OFN_PATHMUTEXIST, // 如果所选文件已存在，则使“另存为”对话框生成一个消息框。用户必须确认是否覆盖文件。| 检测文件路径是否合法
		NFileOffset:       0,
		NFileExtension:    0,
		LpstrDefExt:       0, // 如果用户没有输入文件扩展名, 则默认使用这个
		LCustData:         0,
		LpfnHook:          0,
		LpTemplateName:    0,
	}
	ofn.LStructSize = uint32(unsafe.Sizeof(ofn))
	if !炫彩WinApi类.GetSaveFileNameW(&ofn) {
		return ""
	}
	return 炫彩工具类.UintPtrToString(uintptr(unsafe.Pointer(lpstrFile)))
}

// ChooseColor 选择颜色.
//
//	@param hParent 炫彩窗口句柄, 可为0.
//	@return int 返回rgb颜色.
func ChooseColor(hParent int) int {
	var hwnd uintptr
	if hParent > 0 {
		hwnd = xc.XWnd_GetHWND(hParent)
	}
	var lpCustColors [16]uint32
	cc := 炫彩WinApi类.ChooseColor{
		LStructSize:    36,
		HwndOwner:      hwnd,
		HInstance:      0,
		RgbResult:      0,
		LpCustColors:   &lpCustColors[0],
		Flags:          炫彩WinApi类.CC_FULLOPEN, // 默认打开自定义颜色
		LCustData:      0,
		LpfnHook:       0,
		LpTemplateName: 0,
	}
	cc.LStructSize = uint32(unsafe.Sizeof(cc))
	if !炫彩WinApi类.ChooseColorW(&cc) {
		return 0
	}
	return int(cc.RgbResult)
}
