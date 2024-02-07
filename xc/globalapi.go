package 炫彩基类

import (
	"syscall"
	"unsafe"
	
	"github.com/888go/xcgui/common"
	
	"github.com/888go/xcgui/xcc"
)

// XInitXCGUI 炫彩_初始化.
//
//	@Description 在调用本函数之前请先调用 xc.LoadXCGUI().
//	@param bD2D 是否启用D2D.
//	@return bool
func X初始化(是否启用D2D bool) bool {
	r, _, _ := xInitXCGUI.Call(炫彩工具类.BoolPtr(是否启用D2D))
	return r != 0
}

// XRunXCGUI 炫彩_运行, 运行消息循环, 当炫彩窗口数量为0时退出.
//
//	@return int
func X运行() int {
	r, _, _ := xRunXCGUI.Call()
	return int(r)
}

// XExitXCGUI 炫彩_退出, 退出界面库释放资源.
//
//	@return int
func X退出() int {
	r, _, _ := xExitXCGUI.Call()
	return int(r)
}

// XC_DebugToFileInfo 炫彩_输出调试信息到文件, 打印调试信息到文件xcgui_debug.txt.
//
//	@param pInfo 文本.
func X输出调试信息到文件(文本 string) {
	xC_DebugToFileInfo.Call(W2A(文本))
}

// XC_SetActivateTopWindow 炫彩_激活窗口, 激活当前进程最上层窗口.
//
//	@return bool
func X激活窗口() bool {
	r, _, _ := xC_SetActivateTopWindow.Call()
	return r != 0
}

// XC_GetDefaultFont 炫彩_取默认字体.
//
//	@return int 字体句柄.
func X取默认字体() int {
	r, _, _ := xC_GetDefaultFont.Call()
	return int(r)
}

// XC_MessageBox 炫彩_消息框.
//
//	@param pTitle 标题.
//	@param pText 内容文本.
//	@param nFlags 标识: xcc.MessageBox_Flag_.
//	@param hWndParent 父窗口句柄(真实的窗口句柄).
//	@param XCStyle xcc.Window_Style_.
//	@return xcc.MessageBox_Flag_ , 返回: xcc.MessageBox_Flag_Ok: 点击确定按钮退出. xcc.MessageBox_Flag_Cancel: 点击取消按钮退出. xcc.MessageBox_Flag_Other: 其他方式退出.
func X消息框(标题, 内容文本 string, 标识 炫彩常量类.MessageBox_Flag_, 父窗口句柄 uintptr, 样式 炫彩常量类.Window_Style_) 炫彩常量类.MessageBox_Flag_ {
	r, _, _ := xC_MessageBox.Call(炫彩工具类.StrPtr(标题), 炫彩工具类.StrPtr(内容文本), uintptr(标识), 父窗口句柄, uintptr(样式))
	return 炫彩常量类.MessageBox_Flag_(r)
}

// XMsg_Create 消息框_创建, 此窗口是一个模态窗口, 弹出窗口请调用 xc.XModalWnd_DoModal().
//
//	@param pTitle 标题.
//	@param pText 内容文本.
//	@param nFlags 标识: xcc.MessageBox_Flag_.
//	@param hWndParent 父窗口句柄(真实的窗口句柄).
//	@param XCStyle xcc.Window_Style_.
//	@return int 返回消息框窗口句柄.
func X消息框_创建(标题, 内容文本 string, 标识 炫彩常量类.MessageBox_Flag_, 父窗口句柄 uintptr, 样式 炫彩常量类.Window_Style_) int {
	r, _, _ := xMsg_Create.Call(炫彩工具类.StrPtr(标题), 炫彩工具类.StrPtr(内容文本), uintptr(标识), 父窗口句柄, uintptr(样式))
	return int(r)
}

// XMsg_CreateEx 消息框_创建扩展, 此窗口是一个模态窗口, 弹出窗口请调用 xc.XModalWnd_DoModal().
//
//	@param dwExStyle 窗口扩展样式.
//	@param dwStyle 窗口样式.
//	@param lpClassName 窗口类名.
//	@param pTitle 标题.
//	@param pText 内容文本.
//	@param nFlags 标识: xcc.MessageBox_Flag_ .
//	@param hWndParent 父窗口句柄(真实的窗口句柄).
//	@param XCStyle xcc.Window_Style_ .
//	@return int 消息框窗口句柄.
func X消息框_创建EX(窗口扩展样式 int, 窗口样式 int, 窗口类名, 标题, 内容文本 string, 标识 炫彩常量类.MessageBox_Flag_, 父窗口句柄 uintptr, 样式 炫彩常量类.Window_Style_) int {
	r, _, _ := xMsg_CreateEx.Call(uintptr(窗口扩展样式), uintptr(窗口样式), 炫彩工具类.StrPtr(窗口类名), 炫彩工具类.StrPtr(标题), 炫彩工具类.StrPtr(内容文本), uintptr(标识), 父窗口句柄, uintptr(样式))
	return int(r)
}

// 炫彩_发送窗口消息.
//
// hWindow: 窗口句柄.
//
// msg:.
//
// wParam:.
//
// lParam:.
func X发送窗口消息(窗口句柄 int, 消息值 uint32, 参数1, 参数2 uint) uint {
	r, _, _ := xC_SendMessage.Call(uintptr(窗口句柄), uintptr(消息值), uintptr(参数1), uintptr(参数2))
	return uint(r)
}

// 炫彩_投递窗口消息.
//
// hWindow: 窗口句柄.
//
// msg:.
//
// wParam:.
//
// lParam:.
func X投递窗口消息(窗口句柄 int, 消息值 uint32, 参数1 int32, 参数2 int32) bool {
	r, _, _ := xC_PostMessage.Call(uintptr(窗口句柄), uintptr(消息值), uintptr(参数1), uintptr(参数2))
	return r != 0
}

// XC_CallUiThread 炫彩_调用界面线程, 调用UI线程, 设置回调函数, 在回调函数里操作UI.
//
//	@Description: 回调函数尽量不要使用匿名函数, 使用匿名函数意味着你每次都在创建1个新的回调, 超过2000个时, 程序必将panic.
//	如果使用 xc.XC_CallUiThreadEx 和 xc.XC_CallUiThreader 则没有此限制.
//	@param pCall 回调函数.
//	@param data 传进回调函数的用户自定义数据.
//	@return int
func X调用界面线程(回调函数 func(data int) int, data int) int {
	r, _, _ := xC_CallUiThread.Call(syscall.NewCallback(回调函数), uintptr(data))
	return int(r)
}

// 炫彩_判断元素, 判断是否为元素句柄.
//
// hEle: 元素句柄.
func X判断元素(元素句柄 int) bool {
	r, _, _ := xC_IsHELE.Call(uintptr(元素句柄))
	return r != 0
}

// 炫彩_判断窗口, 判断是否为窗口句柄.
//
// hWindow: 窗口句柄.
func X判断窗口(窗口句柄 int) bool {
	r, _, _ := xC_IsHWINDOW.Call(uintptr(窗口句柄))
	return r != 0
}

// 炫彩_判断形状对象, 判断是否为形状对象.
//
// hShape: 形状对象句柄.
func X判断形状对象(形状对象句柄 int) bool {
	r, _, _ := xC_IsShape.Call(uintptr(形状对象句柄))
	return r != 0
}

// 炫彩_判断句柄包含类型, 判断句柄是否拥有该类型.
//
// hXCGUI: 炫彩句柄.
//
// nType: 句柄类型, XC_OBJECT_TYPE, 以XC_开头的常量.
func X判断句柄包含类型(炫彩句柄 int, 句柄类型 炫彩常量类.XC_OBJECT_TYPE) bool {
	r, _, _ := xC_IsHXCGUI.Call(uintptr(炫彩句柄), uintptr(句柄类型))
	return r != 0
}

// 炫彩_转换HWND到HWINDOW, 通过窗口HWND句柄获取HWINDOW句柄.
//
// hWnd: 窗口真实句柄HWND.
func X转换HWND到HWINDOW(窗口句柄HWND uintptr) int {
	r, _, _ := xC_hWindowFromHWnd.Call(窗口句柄HWND)
	return int(r)
}

// 炫彩_置属性, 设置对象属性.
//
// hXCGUI: 对象句柄.
//
// pName: 属性名.
//
// pValue: 属性值.
func X置属性(对象句柄 int, 属性名 string, 属性值 string) bool {
	r, _, _ := xC_SetProperty.Call(uintptr(对象句柄), 炫彩工具类.StrPtr(属性名), 炫彩工具类.StrPtr(属性值))
	return r != 0
}

// 炫彩_取属性, 获取对象属性, 返回属性值.
//
// hXCGUI: 对象句柄.
//
// pName: 属性名.
func X取属性(对象句柄 int, 属性名 string) string {
	r, _, _ := xC_GetProperty.Call(uintptr(对象句柄), 炫彩工具类.StrPtr(属性名))
	return 炫彩工具类.UintPtrToString(r)
}

// 炫彩_注册窗口类名, 如果是在DLL中使用, 那么DLL卸载时需要注销窗口类名, 否则DLL卸载后, 类名所指向的窗口过程地址失效.
//
// pClassName: 类名.
func X注册窗口类名(类名 string) bool {
	r, _, _ := xC_RegisterWindowClassName.Call(炫彩工具类.StrPtr(类名))
	return r != 0
}

// 炫彩_判断滚动视图扩展元素, 判断元素是否从滚动视图元素扩展的新元素, 包含滚动视图元素.
//
// hEle: 元素句柄.
func X判断滚动视图EX元素(元素句柄 int) bool {
	r, _, _ := xC_IsSViewExtend.Call(uintptr(元素句柄))
	return r != 0
}

// 炫彩_取对象类型, 获取句柄类型, 返回: XC_OBJECT_TYPE.
//
// hXCGUI: 炫彩对象句柄.
func X取对象类型(炫彩对象句柄 int) 炫彩常量类.XC_OBJECT_TYPE {
	r, _, _ := xC_GetObjectType.Call(uintptr(炫彩对象句柄))
	return 炫彩常量类.XC_OBJECT_TYPE(r)
}

// 炫彩_取对象从ID, 通过ID获取对象句柄, 不包括窗口对象.
//
// hWindow: 所属窗口句柄.
//
// nID: ID值.
func X取对象从ID(所属窗口句柄 int, ID值 int) int {
	r, _, _ := xC_GetObjectByID.Call(uintptr(所属窗口句柄), uintptr(ID值))
	return int(r)
}

// 炫彩_取对象从ID名称, 通过ID名称获取对象句柄.
//
// hWindow: 所属窗口句柄.
//
// pName: ID名称.
func X取对象从ID名称(所属窗口句柄 int, ID名称 string) int {
	r, _, _ := xC_GetObjectByIDName.Call(uintptr(所属窗口句柄), 炫彩工具类.StrPtr(ID名称))
	return int(r)
}

// 炫彩_取对象从UID, 通过UID获取对象句柄, 不包括窗口对象.
//
// nUID: UID值.
func X取对象从UID(UID值 int) int {
	r, _, _ := xC_GetObjectByUID.Call(uintptr(UID值))
	return int(r)
}

// 炫彩_取对象从UID名称, 通过UID名称获取对象句柄.
//
// pName: UID名称.
func X取对象从UID名称(UID名称 string) int {
	r, _, _ := xC_GetObjectByUIDName.Call(炫彩工具类.StrPtr(UID名称))
	return int(r)
}

// 炫彩_取对象从名称, 通过name获取对象句柄.
//
// pName: name名称.
func X取对象从名称(名称 string) int {
	r, _, _ := xC_GetObjectByName.Call(炫彩工具类.StrPtr(名称))
	return int(r)
}

// 炫彩_置绘制频率, 设置UI的最小重绘频率.
//
// nMilliseconds: 重绘最小时间间隔, 单位毫秒.
func X置绘制频率(重绘最小时间间隔 int) {
	xC_SetPaintFrequency.Call(uintptr(重绘最小时间间隔))
}

// 炫彩_置文本渲染质量, 设置文本渲染质量.
//
// nType: 参见GDI+ TextRenderingHint 定义.
func X置文本渲染质量(nType int) {
	xC_SetTextRenderingHint.Call(uintptr(nType))
}

// 炫彩_启用GDI绘制文本, 将影响到以下函数: XDraw_TextOut, XDraw_TextOutEx, XDraw_TextOutA.
//
// bEnable: 是否启用.
func X启用GDI绘制文本(是否启用 bool) {
	xC_EnableGdiDrawText.Call(炫彩工具类.BoolPtr(是否启用))
}

// 炫彩_判断矩形相交, 判断两个矩形是否相交及重叠.
//
// pRect1: 矩形1.
//
// pRect2: 矩形2.
func X判断矩形相交(矩形1 *RECT, 矩形2 *RECT) bool {
	r, _, _ := xC_RectInRect.Call(uintptr(unsafe.Pointer(矩形1)), uintptr(unsafe.Pointer(矩形2)))
	return r != 0
}

// 炫彩_组合矩形, 组合两个矩形区域.
//
// pDest: 新的矩形区域.
//
// pSrc1: 源矩形1.
//
// pSrc2: 源矩形2.
func X组合矩形(新的矩形区域 *RECT, 源矩形1 *RECT, 源矩形2 *RECT) {
	xC_CombineRect.Call(uintptr(unsafe.Pointer(新的矩形区域)), uintptr(unsafe.Pointer(源矩形1)), uintptr(unsafe.Pointer(源矩形2)))
}

// 炫彩_显示布局边界, 显示布局对象边界.
//
// bShow: 是否显示.
func X显示布局边界(是否显示 bool) {
	xC_ShowLayoutFrame.Call(炫彩工具类.BoolPtr(是否显示))
}

// 炫彩_启用debug文件.
//
// bEnable: 是否启用.
func X启用debug文件(是否启用 bool) {
	xC_EnableDebugFile.Call(炫彩工具类.BoolPtr(是否启用))
}

// 炫彩_启用资源监视器.
//
// bEnable: 是否启用.
func X启用资源监视器(是否启用 bool) {
	xC_EnableResMonitor.Call(炫彩工具类.BoolPtr(是否启用))
}

// 炫彩_置布局边界颜色.
//
// color: ABGR 颜色值.
func X置布局边界颜色(ABGR颜色值 int) int {
	r, _, _ := xC_SetLayoutFrameColor.Call(uintptr(ABGR颜色值))
	return int(r)
}

// 炫彩_启用错误弹窗, 启用错误弹出, 通过该接口可以设置遇到严重错误时不弹出消息提示框.
//
// bEnabel: 是否启用.
func X启用错误弹窗(是否启用 bool) int {
	r, _, _ := xC_EnableErrorMessageBox.Call(炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 炫彩_启用自动退出程序, 启动或禁用自动退出程序, 当检测到所有用户创建的窗口都关闭时, 自动退出程序; 可调用 XC_PostQuitMessage() 手动退出程序.
//
// bEnabel: 是否启用.
func X启用自动退出程序(是否启用 bool) int {
	r, _, _ := xC_EnableAutoExitApp.Call(炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 炫彩_取文本绘制大小.
//
// pString: 字符串.
//
// length: 字符串长度.
//
// hFontX: 字体.
//
// pOutSize: 接收返回大小.
func X取文本绘制大小(字符串 string, 字符串长度 int, 字体 int, 接收返回大小 *SIZE) int {
	r, _, _ := xC_GetTextSize.Call(炫彩工具类.StrPtr(字符串), uintptr(字符串长度), uintptr(字体), uintptr(unsafe.Pointer(接收返回大小)))
	return int(r)
}

// 炫彩_取文本显示大小.
//
// pString: 字符串.
//
// length: 字符串长度.
//
// hFontX: 字体.
//
// pOutSize: 接收返回大小.
func X取文本显示大小(字符串 string, 字符串长度 int, 字体 int, 接收返回大小 *SIZE) int {
	r, _, _ := xC_GetTextShowSize.Call(炫彩工具类.StrPtr(字符串), uintptr(字符串长度), uintptr(字体), uintptr(unsafe.Pointer(接收返回大小)))
	return int(r)
}

// 炫彩_取文本显示大小扩展.
//
// pString: 字符串.
//
// length: 字符串长度.
//
// hFontX: 字体.
//
// nTextAlign: 文本对齐方式, TextFormatFlag_, TextAlignFlag_, TextTrimming_.
//
// pOutSize: 接收返回大小.
func X取文本显示大小EX(字符串 string, 字符串长度 int, 字体 int, 文本对齐方式 炫彩常量类.TextFormatFlag_, 接收返回大小 *SIZE) int {
	r, _, _ := xC_GetTextShowSizeEx.Call(炫彩工具类.StrPtr(字符串), uintptr(字符串长度), uintptr(字体), uintptr(文本对齐方式), uintptr(unsafe.Pointer(接收返回大小)))
	return int(r)
}

// XC_GetTextShowRect 炫彩_取文本显示矩形.
//
//	@param pString 字符串.
//	@param length 字符串长度.
//	@param hFontX 字体.
//	@param nTextAlign 文本对齐: xcc.TextFormatFlag_.
//	@param width 最大宽度.
//	@param pOutSize 接收返回大小.
//	@return int
func X炫彩_取文本显示矩形(字符串 string, 字符串长度 int, 字体 int, 文本对齐 炫彩常量类.TextFormatFlag_, 最大宽度 int, 接收返回大小 *SIZE) int {
	r, _, _ := xC_GetTextShowRect.Call(炫彩工具类.StrPtr(字符串), uintptr(字符串长度), uintptr(字体), uintptr(文本对齐), uintptr(最大宽度), uintptr(unsafe.Pointer(接收返回大小)))
	return int(r)
}

// 炫彩_置默认字体.
//
// hFontX: 炫彩字体句柄.
func X置默认字体(炫彩字体句柄 int) int {
	r, _, _ := xC_SetDefaultFont.Call(uintptr(炫彩字体句柄))
	return int(r)
}

// 炫彩_添加搜索路径, 添加文件搜索路径, 默认路径为exe目录和程序当前运行目录.
//
// pPath: 文件夹.
func X添加搜索路径(文件夹 string) int {
	r, _, _ := xC_AddFileSearchPath.Call(炫彩工具类.StrPtr(文件夹))
	return int(r)
}

// 炫彩_初始化字体, 初始化LOGFONTW结构体.
//
// pFont: LOGFONTW结构体指针.
//
// pName: 字体名称.
//
// size: 字体大小.
//
// bBold: 是否为粗体.
//
// bItalic: 是否为斜体.
//
// bUnderline: 是否有下划线.
//
// bStrikeOut: 是否有删除线.
func X初始化字体(LOGFONTW结构体指针 *LOGFONTW, 字体名称 string, 字体大小 int, 是否为粗体 bool, 是否为斜体 bool, 是否有下划线 bool, 是否有删除线 bool) int {
	r, _, _ := xC_InitFont.Call(uintptr(unsafe.Pointer(LOGFONTW结构体指针)), 炫彩工具类.StrPtr(字体名称), uintptr(字体大小), 炫彩工具类.BoolPtr(是否为粗体), 炫彩工具类.BoolPtr(是否为斜体), 炫彩工具类.BoolPtr(是否有下划线), 炫彩工具类.BoolPtr(是否有删除线))
	return int(r)
}

// 炫彩_分配内存, 在UI库中申请内存, 返回: 内存首地址.
//
// size: 大小, 字节为单位.
func X分配内存(大小 int) int {
	r, _, _ := xC_Malloc.Call(uintptr(大小))
	return int(r)
}

// 炫彩_释放内存, 在UI库中释放内存.
//
// p: 内存首地址.
func X释放内存(内存首地址 int) int {
	r, _, _ := xC_Free.Call(uintptr(内存首地址))
	return int(r)
}

// 炫彩_弹框, 弹出提示框.
//
// pTitle: 提示框标题.
//
// pText: 提示内容.
func X弹框(提示框标题, 提示内容 string) int {
	r, _, _ := xC_Alert.Call(炫彩工具类.StrPtr(提示框标题), 炫彩工具类.StrPtr(提示内容))
	return int(r)
}

// 对指定文件执行操作. 如果函数成功，则返回大于 32 的值。如果函数失败，则返回指示失败原因的错误值.
//
// hwnd: 用于显示 UI 或错误消息的父窗口的句柄。如果操作与窗口无关，则此值可以为0.
//
// lpOperation: 填“open”则打开lpFlie文档. 其它操作详见: https://docs.microsoft.com/en-us/windows/win32/api/shellapi/nf-shellapi-shellexecutew.
//
// lpFile: 想用关联的程序打印或打开的一个程序名或文件名.
//
// lpParameters: 如果lpFile是一个可执行文件，则这个字串包含了传递给执行程序的参数.
//
// lpDirectory: 想使用的默认路径完整路径.
//
// nShowCmd: 定义了如何显示启动程序的常数值: xcc.SW_.
func X对指定文件执行操作(父窗口句柄 uintptr, 操作类型 string, 文件名 string, 参数 string, 默认完整路径 string, 显示Cmd 炫彩常量类.SW_) uintptr {
	r, _, _ := xC_Sys_ShellExecute.Call(父窗口句柄, 炫彩工具类.StrPtr(操作类型), 炫彩工具类.StrPtr(文件名), 炫彩工具类.StrPtr(参数), 炫彩工具类.StrPtr(默认完整路径), uintptr(显示Cmd))
	return r
}

// 炫彩_载入动态库, 系统API LoadLibrary, 返回动态库模块句柄.
//
// lpFileName: 文件名.
func X载入动态库(文件名 string) uintptr {
	r, _, _ := xC_LoadLibrary.Call(炫彩工具类.StrPtr(文件名))
	return r
}

// 炫彩_取动态库中函数地址, 系统API GetProcAddress, 返回函数地址.
//
// hModule: 动态库模块句柄.
//
// lpProcName: 函数名.
func X取动态库中函数地址(动态库模块句柄 uintptr, 函数名 string) uintptr {
	r, _, _ := xC_GetProcAddress.Call(动态库模块句柄, W2A(函数名))
	return r
}

// 炫彩_释放动态库, 系统API FreeLibrary.
//
// hModule: 动态库模块句柄.
func X释放动态库(动态库模块句柄 uintptr) bool {
	r, _, _ := xC_FreeLibrary.Call(动态库模块句柄)
	return r != 0
}

// 炫彩_加载DLL, 返回DLL模块句柄. 加载指定DLL, 并且调用DLL中函数LoadDll(), DLL中导出函数格式: int WINAPI LoadDll().
//
// pDllFileName: DLL文件名.
func X加载DLL(DLL文件名 string) uintptr {
	r, _, _ := xC_LoadDll.Call(炫彩工具类.StrPtr(DLL文件名))
	return r
}

// 炫彩_PostQuitMessage, 发送WM_QUIT消息退出消息循环.
//
// nExitCode: 退出码.
func X发送WM_QUIT消息退出消息循环(退出码 int) int {
	r, _, _ := xC_PostQuitMessage.Call(uintptr(退出码))
	return int(r)
}

// 炫彩_置D2D文本渲染模式.
//
// mode: 渲染模式, XC_DWRITE_RENDERING_MODE_ .
func X置D2D文本渲染模式(渲染模式 炫彩常量类.XC_DWRITE_RENDERING_MODE_) int {
	r, _, _ := xC_SetD2dTextRenderingMode.Call(uintptr(渲染模式))
	return int(r)
}

// 炫彩_是否启用了D2D.
func X是否启用了D2D() bool {
	r, _, _ := xC_IsEnableD2D.Call()
	return r != 0
}

// 炫彩_W2A.
//
// pValue: 参数.
func W2A(参数 string) uintptr {
	r, _, _ := xC_wtoa.Call(炫彩工具类.StrPtr(参数))
	return r
}

// 炫彩_A2W.
//
// pValue: 参数.
func A2W(参数 uintptr) string {
	r, _, _ := xC_atow.Call(参数)
	return 炫彩工具类.UintPtrToString(r)
}

// 炫彩_UTF8到文本W.
//
// pUtf8: 参数.
func UTF8到文本W(参数 uintptr) string {
	r, _, _ := xC_utf8tow.Call(参数)
	return 炫彩工具类.UintPtrToString(r)
}

// 炫彩_UTF8到文本W扩展.
//
// pUtf8: utf8字符串指针.
//
// length: utf8字符串长度.
func UTF8到文本WEX(utf8字符串指针 uintptr, utf8字符串长度 int) string {
	r, _, _ := xC_utf8towEx.Call(utf8字符串指针, uintptr(utf8字符串长度))
	return 炫彩工具类.UintPtrToString(r)
}

// 炫彩_UTF8到文本A.
//
// pUtf8: utf8字符串指针.
func UTF8到文本A(utf8字符串指针 uintptr) uintptr {
	r, _, _ := xC_utf8toa.Call(utf8字符串指针)
	return r
}

// 炫彩_文本A到UTF8.
//
// pValue: 参数.
func X文本A到UTF8(参数 uintptr) uintptr {
	r, _, _ := xC_atoutf8.Call(参数)
	return r
}

// 炫彩_文本W到UTF8.
//
// pValue: 字符串.
func X文本W到UTF8(字符串 string) uintptr {
	r, _, _ := xC_wtoutf8.Call(炫彩工具类.StrPtr(字符串))
	return r
}

// 炫彩_文本W到UTF8扩展.
//
// pValue: 字符串.
//
// length: 字符串长度.
func X文本W到UTF8EX(字符串 string, 字符串长度 int) uintptr {
	r, _, _ := xC_wtoutf8Ex.Call(炫彩工具类.StrPtr(字符串), uintptr(字符串长度))
	return r
}

// 炫彩_U2A, 返回写入接收缓冲区字节数量.
//
// pIn: 待转换的Unicode字符串.
//
// inLen: pIn字符数量.
//
// pOut: 指向接收转换后的Ansi字符串缓冲区指针.
//
// outLen: pOut缓冲区大小, 字节单位.
func U2A(待转换的Unicode字符串 string, pIn字符数量 int, 转换后缓冲区指针 uintptr, pOut缓冲区大小 int) int {
	r, _, _ := xC_UnicodeToAnsi.Call(炫彩工具类.StrPtr(待转换的Unicode字符串), uintptr(pIn字符数量), 转换后缓冲区指针, uintptr(pOut缓冲区大小))
	return int(r)
}

// 炫彩_A2U, 返回写入接收缓冲区字符数量.
//
// pIn: 指向待转换的Ansi字符串指针.
//
// inLen: pIn字符数量.
//
// pOut: 指向接收转换后的Unicode字符串缓冲区指针.
//
// outLen: pOut缓冲区大小,字符wchar_t单位.
func A2U(指向待转换的Ansi字符串指针 uintptr, pIn字符数量 int, 转换后缓冲区指针 *string, pOut缓冲区大小 int) int {
	buf := make([]uint16, pOut缓冲区大小)
	r, _, _ := xC_AnsiToUnicode.Call(指向待转换的Ansi字符串指针, uintptr(pIn字符数量), 炫彩工具类.Uint16SliceDataPtr(&buf), uintptr(pOut缓冲区大小))
	*转换后缓冲区指针 = syscall.UTF16ToString(buf[0:])
	return int(r)
}

// 炫彩_打印调试信息, 打印调试信息到文件xcgui_debug.txt.
//
// level: 级别.
//
// pInfo: 信息.
func X打印调试信息(级别 int, 信息 string) int {
	r, _, _ := xDebug_Print.Call(uintptr(级别), W2A(信息))
	return int(r)
}

// 炫彩_显示边界.
//
// bShow: 是否显示.
func X显示边界(是否显示 bool) int {
	r, _, _ := xC_ShowSvgFrame.Call(炫彩工具类.BoolPtr(是否显示))
	return int(r)
}

// 炫彩_启用自动DPI. 当启用后, 创建窗口时自动检测DPI调整UI缩放, 处理DPI改变消息; 禁用后,当DPI改变,需要手动设置窗口DPI.
//
// bEnabel: 是否启用.
func X启用自动DPI(是否启用 bool) int {
	r, _, _ := xC_EnableAutoDPI.Call(炫彩工具类.BoolPtr(是否启用))
	return int(r)
}

// 炫彩_启用DPI.
//
// 为go程序启用DPI的几种方式:
//  1. 使用程序清单文件.
//  2. 调用此函数.
//  3. 自己调用DPI函数.
//
// 参考[MSDN](https://learn.microsoft.com/zh-cn/windows/win32/hidpi/setting-the-default-dpi-awareness-for-a-process)
//
// bEnabel: 是否启用.
func X启用DPI(是否启用 bool) bool {
	r, _, _ := xC_EnableDPI.Call(炫彩工具类.BoolPtr(是否启用))
	return r != 0
}

// 炫彩_置窗口图标. 全局窗口图标, 所有未设置图标的窗口, 都将使用此默认图标.
//
// hImage: 图标句柄.
func X置窗口图标(图标句柄 int) int {
	r, _, _ := xC_SetWindowIcon.Call(uintptr(图标句柄))
	return int(r)
}

/* // 炫彩_打印调试信息, 打印调试信息到文件xcgui_debug.txt.[无效]
//
// pString: 字符串.
func XDebug_OutputDebugStringW(pString string) int {
	r, _, _ := xDebug_OutputDebugStringW.Call(common.StrPtr(pString))
	return int(r)
}

// 炫彩_设置debug输出编码方式为utf8.[无效]
//
// bUTF8: 是否开启utf8编码.
func XDebug_Set_OutputDebugString_UTF8(bUTF8 bool) int {
	r, _, _ := xDebug_Set_OutputDebugString_UTF8.Call(common.BoolPtr(bUTF8))
	return int(r)
}

// 炫彩_整数到文本A.
//
// nValue: 参数.
func XC_itoa(nValue int) int {
	r, _, _ := xC_itoa.Call(uintptr(nValue))
	return int(r)
}

// 炫彩_整数到文本W.
//
// nValue: 参数.
func XC_itow(nValue int) int {
	r, _, _ := xC_itow.Call(uintptr(nValue))
	return int(r)
}

// 炫彩_浮点数到文本A.
//
// fValue: 参数.
func XC_ftoa(fValue int) int {
	r, _, _ := xC_ftoa.Call(uintptr(fValue))
	return int(r)
}

// 炫彩_浮点数到文本W.
//
// fValue: 参数.
func XC_ftow(fValue int) int {
	r, _, _ := xC_ftow.Call(uintptr(fValue))
	return int(r)
}
*/
