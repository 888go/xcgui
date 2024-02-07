package 炫彩App类

import (
	"github.com/888go/xcgui/xc"
	"github.com/888go/xcgui/xcc"
)

// App 程序. 封装了炫彩的全局API.
type App struct {
}

// New 炫彩_初始化, 失败返回nil.
//
//	@Description 默认会在程序运行目录和系统system32目录寻找并加载xcgui.dll.
//	如果你想要更改xcgui.dll的路径, 那么请在调用本函数之前调用 xc.SetXcguiPath().
//	@param bD2D 是否启用D2D.
//	@return *App
func X创建(是否启用D2D bool) *App {
	moudle := 炫彩基类.XLoadXCGUI()
	if moudle.Handle() == 0 {
		return nil
	}

	p := &App{}
	if !炫彩基类.X初始化(是否启用D2D) {
		return nil
	}
	return p
}

// Run 炫彩_运行. 运行消息循环, 当炫彩窗口数量为0时退出.
//
//	@return int
func (a *App) X运行() int {
	return 炫彩基类.X运行()
}

// Exit 炫彩_退出, 退出界面库释放资源.
//
//	@return int
func (a *App) X退出() int {
	return 炫彩基类.X退出()
}

// ShowAndRun 显示窗口并调用炫彩_运行.
//
//	@param hWindow 炫彩窗口句柄.
func (a *App) X显示窗口并运行(炫彩窗口句柄 int) {
	炫彩基类.X窗口_显示方式(炫彩窗口句柄, 炫彩常量类.SW_SHOW)
	炫彩基类.X运行()
}

// DebugToFileInfo 炫彩_输出调试信息到文件, 打印调试信息到文件xcgui_debug.txt.
//
//	@param pInfo 文本.
func (a *App) X输出调试信息到文件(文本 string) {
	炫彩基类.X输出调试信息到文件(文本)
}

// SetActivateTopWindow 炫彩_激活窗口, 激活当前进程最上层窗口.
//
//	@return bool
func (a *App) X激活窗口() bool {
	return 炫彩基类.X激活窗口()
}

// GetDefaultFont 炫彩_取默认字体.
//
//	@return int 字体句柄.
func (a *App) X取默认字体() int {
	return 炫彩基类.X取默认字体()
}

// MessageBox 炫彩_消息框.
//
//	@param pTitle 标题.
//	@param pText 内容文本.
//	@param nFlags 标识: xcc.MessageBox_Flag_.
//	@param hWndParent 父窗口句柄(真实的窗口句柄).
//	@param XCStyle xcc.Window_Style_.
//	@return xcc.MessageBox_Flag_. 返回: xcc.MessageBox_Flag_Ok: 点击确定按钮退出. xcc.MessageBox_Flag_Cancel: 点击取消按钮退出. xcc.MessageBox_Flag_Other: 其他方式退出.
func (a *App) X消息框(标题, 内容文本 string, 标识 炫彩常量类.MessageBox_Flag_, 父窗口句柄 uintptr, 炫彩样式 炫彩常量类.Window_Style_) 炫彩常量类.MessageBox_Flag_ {
	return 炫彩基类.X消息框(标题, 内容文本, 标识, 父窗口句柄, 炫彩样式)
}

// Msg_Create 消息框_创建, 此窗口是一个模态窗口, 弹出窗口请调用 xc.XModalWnd_DoModal().
//
//	@param pTitle 标题.
//	@param pText 内容文本.
//	@param nFlags 标识: xcc.MessageBox_Flag_.
//	@param hWndParent 父窗口句柄(真实的窗口句柄).
//	@param XCStyle xcc.Window_Style_.
//	@return int 返回消息框窗口句柄.
func (a *App) X消息框_创建(标题, 内容文本 string, 标识 炫彩常量类.MessageBox_Flag_, 父窗口句柄 uintptr, 炫彩样式 炫彩常量类.Window_Style_) int {
	return 炫彩基类.X消息框_创建(标题, 内容文本, 标识, 父窗口句柄, 炫彩样式)
}

// Msg_CreateEx 消息框_创建扩展, 此窗口是一个模态窗口, 弹出窗口请调用 xc.XModalWnd_DoModal().
//
//	@param dwExStyle 窗口扩展样式.
//	@param dwStyle 窗口样式.
//	@param lpClassName 窗口类名.
//	@param pTitle 标题.
//	@param pText 内容文本.
//	@param nFlags 标识: xcc.MessageBox_Flag_.
//	@param hWndParent 父窗口句柄(真实的窗口句柄).
//	@param XCStyle xcc.Window_Style_.
//	@return int 消息框窗口句柄.
func (a *App) X消息框_创建EX(窗口样式EX int, 窗口样式 int, 窗口类名 string, 标题, 内容文本 string, 标识 炫彩常量类.MessageBox_Flag_, 父窗口句柄 uintptr, 炫彩样式 炫彩常量类.Window_Style_) int {
	return 炫彩基类.X消息框_创建EX(窗口样式EX, 窗口样式, 窗口类名, 标题, 内容文本, 标识, 父窗口句柄, 炫彩样式)
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
func (a *App) X发送窗口消息(窗口句柄 int, 消息值 uint32, 参数1, 参数2 uint) uint {
	return 炫彩基类.X发送窗口消息(窗口句柄, 消息值, 参数1, 参数2)
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
func (a *App) X投递窗口消息(窗口句柄 int, 消息值 uint32, 参数1 int32, 参数2 int32) bool {
	return 炫彩基类.X投递窗口消息(窗口句柄, 消息值, 参数1, 参数2)
}

// CallUiThread 炫彩_调用界面线程, 调用UI线程, 设置回调函数, 在回调函数里操作UI.
//
//	@Description: 回调函数尽量不要使用匿名函数, 使用匿名函数意味着你每次都在创建1个新的回调, 超过2000个时, 程序必将panic.
//	如果使用 CallUiThreadEx 和 CallUiThreader 则没有此限制.
//	@param pCall 回调函数.
//	@param data 传进回调函数的用户自定义数据.
//	@return int
func (a *App) X内部_调用界面线程(回调函数 func(data int) int, data int) int {
	return 炫彩基类.X调用界面线程(回调函数, data)
}

// CallUiThreadEx 炫彩_调用界面线程, 调用UI线程, 设置回调函数, 在回调函数里操作UI.
//
//	@Description: 与 CallUiThread 的区别是: 本函数没有2000个回调上限的限制, 回调函数可以直接使用匿名函数.
//	@param f 回调函数.
//	@param data 传进回调函数的用户自定义数据.
//	@return int
func (a *App) X调用界面线程EX(回调函数 func(data int) int, data int) int {
	return 炫彩基类.X炫彩_调用界面线程Ex(回调函数, data)
}

// CallUT 炫彩_调用界面线程, 调用UI线程, 设置回调函数, 在回调函数里操作UI.
//
//	@Description: 与 CallUiThread 的区别是: 本函数没有2000个回调上限的限制, 回调函数可以直接使用匿名函数. 回调函数没有参数也没有返回值.
//	@param f 回调函数, 没有参数也没有返回值, 可以直接使用匿名函数.
func (a *App) X简易调用界面线程(回调函数 func()) {
	炫彩基类.X炫彩_调用界面线程(回调函数)
}

// CallUiThreader 炫彩_调用界面线程, 调用UI线程, 设置回调函数, 在回调函数里操作UI.
//
//	@Description: 与 CallUiThread 的区别是: 本函数没有2000个回调上限的限制, 回调函数可以直接使用匿名函数.
//	@param u xc.UiThreader.
//	@param data 传进回调函数的用户自定义数据.
//	@return int
func (a *App) X调用界面线程(u 炫彩基类.UiThreader, 回调函数数据 int) int {
	return 炫彩基类.X炫彩_调用界面线程1(u, 回调函数数据)
}

// 炫彩_判断元素, 判断是否为元素句柄.
//
// hEle: 元素句柄.
func (a *App) X判断元素(元素句柄 int) bool {
	return 炫彩基类.X判断元素(元素句柄)
}

// 炫彩_判断窗口, 判断是否为窗口句柄.
//
// hWindow: 窗口句柄.
func (a *App) X判断窗口(窗口句柄 int) bool {
	return 炫彩基类.X判断窗口(窗口句柄)
}

// 炫彩_判断形状对象, 判断是否为形状对象.
//
// hShape: 形状对象句柄.
func (a *App) X判断形状对象(形状对象句柄 int) bool {
	return 炫彩基类.X判断形状对象(形状对象句柄)
}

// 炫彩_判断句柄包含类型, 判断句柄是否拥有该类型.
//
// hXCGUI: 炫彩句柄.
//
// nType: 句柄类型, XC_OBJECT_TYPE , 以XC_开头的常量.
func (a *App) X判断句柄包含类型(炫彩句柄 int, 句柄类型 炫彩常量类.XC_OBJECT_TYPE) bool {
	return 炫彩基类.X判断句柄包含类型(炫彩句柄, 句柄类型)
}

// 炫彩_转换HWND到HWINDOW, 通过窗口HWND句柄获取HWINDOW句柄.
//
// hWnd: 窗口真实句柄HWND.
func (a *App) X转换HWND到HWINDOW(窗口真实句柄HWND uintptr) int {
	return 炫彩基类.X转换HWND到HWINDOW(窗口真实句柄HWND)
}

// 炫彩_置属性, 设置对象属性.
//
// hXCGUI: 对象句柄.
//
// pName: 属性名.
//
// pValue: 属性值.
func (a *App) X置属性(对象句柄 int, 属性名 string, 属性值 string) bool {
	return 炫彩基类.X置属性(对象句柄, 属性名, 属性值)
}

// 炫彩_取属性, 获取对象属性, 返回属性值.
//
// hXCGUI: 对象句柄.
//
// pName: 属性名.
func (a *App) X取属性(对象句柄 int, 属性名 string) string {
	return 炫彩基类.X取属性(对象句柄, 属性名)
}

// 炫彩_注册窗口类名, 如果是在DLL中使用, 那么DLL卸载时需要注销窗口类名, 否则DLL卸载后, 类名所指向的窗口过程地址失效.
//
// pClassName: 类名.
func (a *App) X注册窗口类名(类名 string) bool {
	return 炫彩基类.X注册窗口类名(类名)
}

// 炫彩_判断滚动视图扩展元素, 判断元素是否从滚动视图元素扩展的新元素, 包含滚动视图元素.
//
// hEle: 元素句柄.
func (a *App) X判断滚动视图扩展元素(元素句柄 int) bool {
	return 炫彩基类.X判断滚动视图EX元素(元素句柄)
}

// 炫彩_取对象类型, 获取句柄类型, 返回: xcc.XC_OBJECT_TYPE, 以XC_开头的常量.
//
// hXCGUI: 炫彩对象句柄.
func (a *App) X取对象类型(炫彩对象句柄 int) 炫彩常量类.XC_OBJECT_TYPE {
	return 炫彩基类.X取对象类型(炫彩对象句柄)
}

// 炫彩_取对象从ID, 通过ID获取对象句柄, 不包括窗口对象.
//
// hWindow: 所属窗口句柄.
//
// nID: ID值.
func (a *App) X取对象从ID(所属窗口句柄 int, ID值 int) int {
	return 炫彩基类.X取对象从ID(所属窗口句柄, ID值)
}

// 炫彩_取对象从ID名称, 通过ID名称获取对象句柄.
//
// hWindow: 所属窗口句柄.
//
// pName: ID名称.
func (a *App) X取对象从ID名称(所属窗口句柄 int, ID名称 string) int {
	return 炫彩基类.X取对象从ID名称(所属窗口句柄, ID名称)
}

// 炫彩_取对象从UID, 通过UID获取对象句柄, 不包括窗口对象.
//
// nUID: UID值.
func (a *App) X取对象从UID(UID值 int) int {
	return 炫彩基类.X取对象从UID(UID值)
}

// 炫彩_取对象从UID名称, 通过UID名称获取对象句柄.
//
// pName: UID名称.
func (a *App) X取对象从UID名称(UID名称 string) int {
	return 炫彩基类.X取对象从UID名称(UID名称)
}

// 炫彩_取对象从名称, 通过name获取对象句柄.
//
// pName: name名称.
func (a *App) X取对象从名称(name名称 string) int {
	return 炫彩基类.X取对象从名称(name名称)
}

// 炫彩_置绘制频率, 设置UI的最小重绘频率.
//
// nMilliseconds: 重绘最小时间间隔, 单位毫秒.
func (a *App) X置绘制频率(重绘最小时间间隔 int) {
	炫彩基类.X置绘制频率(重绘最小时间间隔)
}

// 炫彩_置文本渲染质量, 设置文本渲染质量.
//
// nType: 参见GDI+ TextRenderingHint 定义.
func (a *App) X置文本渲染质量(nType int) {
	炫彩基类.X置文本渲染质量(nType)
}

// 炫彩_启用GDI绘制文本, 将影响到以下函数: XDraw_TextOut, XDraw_TextOutEx, XDraw_TextOutA.
//
// bEnable: 是否启用.
func (a *App) X启用GDI绘制文本(是否启用 bool) {
	炫彩基类.X启用GDI绘制文本(是否启用)
}

// 炫彩_判断矩形相交, 判断两个矩形是否相交及重叠.
//
// pRect1: 矩形1.
//
// pRect2: 矩形2.
func (a *App) X判断矩形相交(矩形1 *炫彩基类.RECT, 矩形2 *炫彩基类.RECT) bool {
	return 炫彩基类.X判断矩形相交(矩形1, 矩形2)
}

// 炫彩_组合矩形, 组合两个矩形区域.
//
// pDest: 新的矩形区域.
//
// pSrc1: 源矩形1.
//
// pSrc2: 源矩形2.
func (a *App) X组合矩形(新的矩形区域 *炫彩基类.RECT, 源矩形1 *炫彩基类.RECT, 源矩形2 *炫彩基类.RECT) {
	炫彩基类.X组合矩形(新的矩形区域, 源矩形1, 源矩形2)
}

// 炫彩_显示布局边界, 显示布局对象边界.
//
// bShow: 是否显示.
func (a *App) X显示布局边界(是否显示 bool) {
	炫彩基类.X显示布局边界(是否显示)
}

// 炫彩_启用debug文件.
//
// bEnable: 是否启用.
func (a *App) X启用debug文件(是否启用 bool) {
	炫彩基类.X启用debug文件(是否启用)
}

// 炫彩_启用资源监视器.
//
// bEnable: 是否启用.
func (a *App) X启用资源监视器(是否启用 bool) {
	炫彩基类.X启用资源监视器(是否启用)
}

// 炫彩_置布局边界颜色.
//
// color: ABGR 颜色值.
func (a *App) X置布局边界颜色(ABGR int) int {
	return 炫彩基类.X置布局边界颜色(ABGR)
}

// 炫彩_启用错误弹窗, 启用错误弹出, 通过该接口可以设置遇到严重错误时不弹出消息提示框.
//
// bEnabel: 是否启用.
func (a *App) X启用错误弹窗(是否启用 bool) int {
	return 炫彩基类.X启用错误弹窗(是否启用)
}

// 炫彩_启用自动退出程序, 启动或禁用自动退出程序, 当检测到所有用户创建的窗口都关闭时, 自动退出程序; 可调用 XC_PostQuitMessage() 手动退出程序.
//
// bEnabel: 是否启用.
func (a *App) X启用自动退出程序(是否启用 bool) int {
	return 炫彩基类.X启用自动退出程序(是否启用)
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
func (a *App) X取文本绘制大小(字符串 string, 字符串长度 int, 字体 int, 接收返回大小 *炫彩基类.SIZE) int {
	return 炫彩基类.X取文本绘制大小(字符串, 字符串长度, 字体, 接收返回大小)
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
func (a *App) X取文本显示大小(字符串 string, 字符串长度 int, 字体 int, 接收返回大小 *炫彩基类.SIZE) int {
	return 炫彩基类.X取文本显示大小(字符串, 字符串长度, 字体, 接收返回大小)
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
func (a *App) X取文本显示大小EX(字符串 string, 字符串长度 int, 字体 int, 文本对齐方式 炫彩常量类.TextFormatFlag_, 接收返回大小 *炫彩基类.SIZE) int {
	return 炫彩基类.X取文本显示大小EX(字符串, 字符串长度, 字体, 文本对齐方式, 接收返回大小)
}

// GetTextShowRect 炫彩_取文本显示矩形.
//
//	@param pString 字符串.
//	@param length 字符串长度.
//	@param hFontX 字体.
//	@param nTextAlign 文本对齐: xcc.TextFormatFlag_.
//	@param width 最大宽度.
//	@param pOutSize 接收返回大小.
//	@return int
func (a *App) X取文本显示矩形(字符串 string, 字符串长度 int, 字体 int, 文本对齐 炫彩常量类.TextFormatFlag_, 最大宽度 int, 接收返回大小 *炫彩基类.SIZE) int {
	return 炫彩基类.X炫彩_取文本显示矩形(字符串, 字符串长度, 字体, 文本对齐, 最大宽度, 接收返回大小)
}

// 炫彩_置默认字体.
//
// hFontX: 炫彩字体句柄.
func (a *App) X置默认字体(炫彩字体句柄 int) int {
	return 炫彩基类.X置默认字体(炫彩字体句柄)
}

// 炫彩_添加搜索路径, 添加文件搜索路径, 默认路径为exe目录和程序当前运行目录.
//
// pPath: 文件夹.
func (a *App) X添加搜索路径(文件夹 string) int {
	return 炫彩基类.X添加搜索路径(文件夹)
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
func (a *App) X初始化字体(LOGFONTW结构体指针 *炫彩基类.LOGFONTW, 字体名称 string, 字体大小 int, 是否为粗体 bool, 是否为斜体 bool, 是否有下划线 bool, 是否有删除线 bool) int {
	return 炫彩基类.X初始化字体(LOGFONTW结构体指针, 字体名称, 字体大小, 是否为粗体, 是否为斜体, 是否有下划线, 是否有删除线)
}

// 炫彩_分配内存, 在UI库中申请内存, 返回: 内存首地址.
//
// size: 大小, 字节为单位.
func (a *App) X分配内存(大小 int) int {
	return 炫彩基类.X分配内存(大小)
}

// 炫彩_释放内存, 在UI库中释放内存.
//
// p: 内存首地址.
func (a *App) X释放内存(内存首地址 int) int {
	return 炫彩基类.X释放内存(内存首地址)
}

// 炫彩_弹框, 弹出提示框.
//
// pTitle: 提示框标题.
//
// pText: 提示内容.
func (a *App) X弹框(提示框标题 string, 提示内容 string) int {
	return 炫彩基类.X弹框(提示框标题, 提示内容)
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
func (a *App) X对指定文件执行操作(hwnd uintptr, lpOperation string, lpFile string, lpParameters string, lpDirectory string, nShowCmd 炫彩常量类.SW_) uintptr {
	return 炫彩基类.X对指定文件执行操作(hwnd, lpOperation, lpFile, lpParameters, lpDirectory, nShowCmd)
}

// 炫彩_载入动态库, 系统API LoadLibrary, 返回动态库模块句柄.
//
// lpFileName: 文件名.
func (a *App) X载入动态库(文件名 string) uintptr {
	return 炫彩基类.X载入动态库(文件名)
}

// 炫彩_取动态库中函数地址, 系统API GetProcAddress, 返回函数地址.
//
// hModule: 动态库模块句柄.
//
// lpProcName: 函数名.
func (a *App) X取动态库中函数地址(动态库模块句柄 uintptr, 函数名 string) uintptr {
	return 炫彩基类.X取动态库中函数地址(动态库模块句柄, 函数名)
}

// 炫彩_释放动态库, 系统API FreeLibrary.
//
// hModule: 动态库模块句柄.
func (a *App) X释放动态库(动态库模块句柄 uintptr) bool {
	return 炫彩基类.X释放动态库(动态库模块句柄)
}

// 炫彩_加载DLL, 返回DLL模块句柄. 加载指定DLL, 并且调用DLL中函数LoadDll(), DLL中导出函数格式: int WINAPI LoadDll().
//
// pDllFileName: DLL文件名.
func (a *App) X加载DLL(DLL文件名 string) uintptr {
	return 炫彩基类.X加载DLL(DLL文件名)
}

// 炫彩_PostQuitMessage, 发送WM_QUIT消息退出消息循环.
//
// nExitCode: 退出码.
func (a *App) X发送WM_QUIT消息退出消息循环(退出码 int) int {
	return 炫彩基类.X发送WM_QUIT消息退出消息循环(退出码)
}

// 炫彩_加载布局文件, 返回窗口句柄或布局句柄或元素句柄.
//
// pFileName: 布局文件名.
//
// hParent: 父对象句柄.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func (a *App) X加载布局文件(布局文件名 string, 父对象句柄 int, 附加窗口句柄 uintptr) int {
	return 炫彩基类.X炫彩_加载布局文件(布局文件名, 父对象句柄, 附加窗口句柄)
}

// 炫彩_加载布局文件ZIP, 加载布局文件从zip压缩包中, 返回窗口句柄或布局句柄或元素句柄.
//
// pZipFileName: zip文件名.
//
// pFileName: 布局文件名.
//
// pPassword: zip密码.
//
// hParent: 父对象句柄.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func (a *App) X加载布局文件ZIP(zip文件名 string, 布局文件名 string, zip密码 string, 父对象句柄 int, 附加窗口句柄 uintptr) int {
	return 炫彩基类.X炫彩_加载布局文件ZIP(zip文件名, 布局文件名, zip密码, 父对象句柄, 附加窗口句柄)
}

// 炫彩_加载布局文件内存ZIP, 加载布局文件从zip压缩包中, 返回窗口句柄或布局句柄或元素句柄.
//
// data: 内存块指针.
//
// pFileName: 布局文件名.
//
// pPassword: zip密码.
//
// hParent: 父对象句柄.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func (a *App) X加载布局文件内存ZIP(内存块指针 []byte, 布局文件名 string, zip密码 string, 父对象句柄 int, 附加窗口句柄 uintptr) int {
	return 炫彩基类.X炫彩_加载布局文件内存ZIP(内存块指针, 布局文件名, zip密码, 父对象句柄, 附加窗口句柄)
}

// 炫彩_加载布局文件资源ZIP扩展. 加载布局文件从RC资源zip压缩包中, 返回窗口句柄或元素句柄.
//
// id: RC资源ID.
//
// pFileName: 布局文件名.
//
// pPassword: zip密码.
//
// pPrefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄, 窗口句柄或UI元素句柄, 可填0.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用, 可填0.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
//
// hModule: 模块句柄, 可填0.
func (a *App) X加载布局文件资源ZIPEX(RC资源ID int32, 布局文件名 string, zip密码, 名称前缀 string, 父对象句柄 int, 父窗口句柄HWND, 附加窗口句柄, 模块句柄 uintptr) int {
	return 炫彩基类.X炫彩_加载布局文件资源ZIPEX(RC资源ID, 布局文件名, zip密码, 名称前缀, 父对象句柄, 父窗口句柄HWND, 附加窗口句柄, 模块句柄)
}

// 炫彩_加载资源文件资源ZIP. 加载资源文件从RC资源zip压缩包中.
//
// id: RC资源ID.
//
// pFileName: 资源文件名.
//
// pPassword: zip压缩包密码.
//
// hModule: 模块句柄, 可填0.
func (a *App) X加载资源文件资源ZIP(RC资源ID int, 资源文件名 string, zip压缩包密码 string, 模块句柄 uintptr) bool {
	return 炫彩基类.X炫彩_加载资源文件资源ZIP(RC资源ID, 资源文件名, zip压缩包密码, 模块句柄)
}

// 炫彩_加载样式文件从资源ZIP. 从RC资源中的ZIP包中, 加载样式文件.
//
// id: RC资源ID.
//
// pFileName: 文件名.
//
// pPassword: 密码.
//
// hModule: 模块句柄, 可填0.
func (a *App) X加载样式文件从资源ZIP(RC资源ID int, 文件名 string, 密码 string, 模块句柄 uintptr) bool {
	return 炫彩基类.X炫彩_加载样式文件从资源ZIP(RC资源ID, 文件名, 密码, 模块句柄)
}

// 炫彩_加载布局文件从字符串W, 加载布局文件从内存字符串, 返回窗口句柄或布局句柄或元素句柄.
//
// pStringXML: 字符串.
//
// hParent: 父对象.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func (a *App) X加载布局文件从字符串W(字符串 string, 父对象 int, 附加窗口句柄 uintptr) int {
	return 炫彩基类.X炫彩_加载布局文件从字符串W(字符串, 父对象, 附加窗口句柄)
}

// 炫彩_加载布局文件Ex, 返回窗口句柄或布局句柄或元素句柄.
//
// pFileName: 布局文件名.
//
// pPrefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄, 窗口句柄或UI元素句柄.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func (a *App) X加载布局文件Ex(布局文件名, 名称 string, 父对象句柄 int, 父窗口句柄HWND, 附加窗口句柄 uintptr) int {
	return 炫彩基类.X炫彩_加载布局文件Ex(布局文件名, 名称, 父对象句柄, 父窗口句柄HWND, 附加窗口句柄)
}

// 炫彩_加载布局文件ZIPEx, 加载布局文件从zip压缩包中, 返回窗口句柄或布局句柄或元素句柄.
//
// pZipFileName: zip文件名.
//
// pFileName: 布局文件名.
//
// pPassword: zip密码.
//
// pPrefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄, 窗口句柄或UI元素句柄.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func (a *App) X加载布局文件ZIPEx(zip文件名 string, 布局文件名 string, zip密码, 名称 string, 父对象句柄 int, 父窗口句柄HWND, 附加窗口句柄 uintptr) int {
	return 炫彩基类.X炫彩_加载布局文件ZIPEx(zip文件名, 布局文件名, zip密码, 名称, 父对象句柄, 父窗口句柄HWND, 附加窗口句柄)
}

// 炫彩_加载布局文件内存ZIPEx, 加载布局文件从zip压缩包中, 返回窗口句柄或布局句柄或元素句柄.
//
// data: 布局文件数据.
//
// pFileName: 布局文件名.
//
// pPassword: zip密码.
//
// pPrefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄, 窗口句柄或UI元素句柄.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func (a *App) X加载布局文件内存ZIPEx(布局文件数据 []byte, 布局文件名 string, zip密码, 名称 string, 父对象句柄 int, 父窗口句柄HWND, 附加窗口句柄 uintptr) int {
	return 炫彩基类.X炫彩_加载布局文件内存ZIPEx(布局文件数据, 布局文件名, zip密码, 名称, 父对象句柄, 父窗口句柄HWND, 附加窗口句柄)
}

// 炫彩_加载布局文件从字符串WEx, 加载布局文件从内存字符串, 返回窗口句柄或布局句柄或元素句柄.
//
// pStringXML: 字符串.
//
// pPrefixName: 名称(name)前缀, 可选参数; 给当前布局文件中所有name属性增加前缀, 那么name属性值为: 前缀 + name.
//
// hParent: 父对象句柄, 窗口句柄或UI元素句柄.
//
// hParentWnd: 父窗口句柄HWND, 提供给第三方窗口使用.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func (a *App) X加载布局文件从字符串WEx(字符串, 名称 string, 父对象句柄 int, 父窗口句柄HWND, 附加窗口句柄 uintptr) int {
	return 炫彩基类.X炫彩_加载布局文件从字符串WEx(字符串, 名称, 父对象句柄, 父窗口句柄HWND, 附加窗口句柄)
}

/*
// 炫彩_加载布局文件从字符串, 加载布局文件从内存字符串, 返回窗口句柄或布局句柄或元素句柄.
//
// pStringXML: 字符串.
//
// hParent: 父对象.
//
// hAttachWnd: 附加窗口句柄, 附加到指定的窗口, 可填0.
func (a *App) LoadLayoutFromString(pStringXML string, hParent, hAttachWnd int) int {
	return xc.XC_LoadLayoutFromString(pStringXML, hParent, hAttachWnd)
}
*/

// 炫彩_加载样式文件.
//
// pFileName: 样式文件名称.
func (a *App) X加载样式文件(样式文件名称 string) bool {
	return 炫彩基类.X炫彩_加载样式文件(样式文件名称)
}

// 炫彩_加载样式文件ZIP.
//
// pZipFile: ZIP文件名.
//
// pFileName: 文件名.
//
// pPassword: 密码.
func (a *App) X加载样式文件ZIP(ZIP文件名 string, 文件名 string, 密码 string) bool {
	return 炫彩基类.X炫彩_加载样式文件ZIP(ZIP文件名, 文件名, 密码)
}

// 炫彩_加载样式文件从内存ZIP.
//
// data: 样式文件数据.
//
// pFileName: 文件名.
//
// pPassword: 密码.
func (a *App) X加载样式文件从内存ZIP(样式文件数据 []byte, 文件名 string, 密码 string) bool {
	return 炫彩基类.X炫彩_加载样式文件从内存ZIP(样式文件数据, 文件名, 密码)
}

// 炫彩_加载资源文件.
//
// pFileName: 资源文件名.
func (a *App) X加载资源文件(资源文件名 string) bool {
	return 炫彩基类.X炫彩_加载资源文件(资源文件名)
}

// 炫彩_加载资源文件ZIP.
//
// pZipFileName: zip文件名.
//
// pFileName: 资源文件名.
//
// pPassword: zip压缩包密码.
func (a *App) X加载资源文件ZIP(zip文件名 string, 资源文件名 string, zip压缩包密码 string) bool {
	return 炫彩基类.X炫彩_加载资源文件ZIP(zip文件名, 资源文件名, zip压缩包密码)
}

// 炫彩_加载资源文件内存ZIP.
//
// data: 资源文件数据.
//
// pFileName: 资源文件名.
//
// pPassword: zip压缩包密码.
func (a *App) X加载资源文件内存ZIP(资源文件数据 []byte, 资源文件名 string, zip压缩包密码 string) bool {
	return 炫彩基类.X炫彩_加载资源文件内存ZIP(资源文件数据, 资源文件名, zip压缩包密码)
}

// 炫彩_加载资源文件从字符串W.
//
// pStringXML: 字符串.
//
// pFileName: 资源文件名.
func (a *App) X加载资源文件从字符串W(字符串 string, 资源文件名 string) bool {
	return 炫彩基类.X炫彩_加载资源文件从字符串W(字符串, 资源文件名)
}

// 炫彩_W2A.
//
// pValue: 参数.
func (a *App) W2A(参数 string) uintptr {
	return 炫彩基类.W2A(参数)
}

// 炫彩_A2W.
//
// pValue: 参数.
func (a *App) A2W(参数 uintptr) string {
	return 炫彩基类.A2W(参数)
}

// 炫彩_UTF8到文本W.
//
// pUtf8: 参数.
func (a *App) UTF8到文本W(参数 uintptr) string {
	return 炫彩基类.UTF8到文本W(参数)
}

// 炫彩_UTF8到文本W扩展.
//
// pUtf8: utf8字符串指针.
//
// length: utf8字符串长度.
func (a *App) UTF8到文本WEX(utf8字符串指针 uintptr, utf8字符串长度 int) string {
	return 炫彩基类.UTF8到文本WEX(utf8字符串指针, utf8字符串长度)
}

// 炫彩_UTF8到文本A.
//
// pUtf8: utf8字符串指针.
func (a *App) UTF8到文本A(utf8字符串指针 uintptr) uintptr {
	return 炫彩基类.UTF8到文本A(utf8字符串指针)
}

// 炫彩_文本A到UTF8.
//
// pValue: 参数.
func (a *App) X文本A到UTF8(参数 uintptr) uintptr {
	return 炫彩基类.X文本A到UTF8(参数)
}

// 炫彩_文本W到UTF8.
//
// pValue: 字符串.
func (a *App) X文本W到UTF8(字符串 string) uintptr {
	return 炫彩基类.X文本W到UTF8(字符串)
}

// 炫彩_文本W到UTF8扩展.
//
// pValue: 字符串.
//
// length: 字符串长度.
func (a *App) X文本W到UTF8EX(字符串 string, 字符串长度 int) uintptr {
	return 炫彩基类.X文本W到UTF8EX(字符串, 字符串长度)
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
func (a *App) U2A(待转换Unicode文本 string, pIn字符数量 int, 接收转换后缓冲区指针 uintptr, pOut缓冲区大小 int) int {
	return 炫彩基类.U2A(待转换Unicode文本, pIn字符数量, 接收转换后缓冲区指针, pOut缓冲区大小)
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
func (a *App) XA2U(待转换Ansi文本指针 uintptr, pIn字符数量 int, 接收转换后缓冲区指针 *string, pOut缓冲区大小 int) int {
	return 炫彩基类.A2U(待转换Ansi文本指针, pIn字符数量, 接收转换后缓冲区指针, pOut缓冲区大小)
}

// 炫彩_打印调试信息, 打印调试信息到文件xcgui_debug.txt.
//
// level: 级别.
//
// pInfo: 信息.
func (a *App) X打印调试信息(级别 int, 信息 string) int {
	return 炫彩基类.X打印调试信息(级别, 信息)
}

/*
// 炫彩_加载资源文件从字符串.
//
// pStringXML: 字符串.
//
// pFileName: 资源文件名.
func (a *App) LoadResourceFromString(pStringXML string, pFileName string) bool {
	return xc.XC_LoadResourceFromString(pStringXML, pFileName)
}

// 炫彩_加载样式文件从字符串.
//
// pFileName: 样式文件名, 用于打印错误文件和定位关联资源文件位置.
//
// pString: 字符串.
func (a *App) LoadStyleFromString(pFileName string, pString string) bool {
	return xc.XC_LoadStyleFromString(pFileName, pString)
} */

// 炫彩_取D2D工厂, 开启D2D有效, 返回 ID2D1Factory* .
func (a *App) X取D2D工厂() int {
	return 炫彩基类.X取D2D工厂()
}

// 炫彩_取DWrite工厂, 开启D2D有效, 返回 IDWriteFactory* .
func (a *App) X取DWrite工厂() int {
	return 炫彩基类.X取DWrite工厂()
}

// 炫彩_取WIC工厂, 开启D2D有效, 返回 IWICImagingFactory* .
func (a *App) X取WIC工厂() int {
	return 炫彩基类.X取WIC工厂()
}

// 炫彩_置D2D文本渲染模式.
//
// mode: 渲染模式, XC_DWRITE_RENDERING_MODE_ .
func (a *App) X置D2D文本渲染模式(渲染模式 炫彩常量类.XC_DWRITE_RENDERING_MODE_) int {
	return 炫彩基类.X置D2D文本渲染模式(渲染模式)
}

// 炫彩_是否启用了D2D.
func (a *App) X是否启用了D2D() bool {
	return 炫彩基类.X是否启用了D2D()
}

// 炫彩_加载样式文件从字符串W.
//
// pFileName: 样式文件名.
//
// pString: 字符串.
func (a *App) X加载样式文件并按字符串W(样式文件名 string, 字符串 string) bool {
	return 炫彩基类.X炫彩_加载样式文件从字符串W(样式文件名, 字符串)
}

// 炫彩_显示边界.
//
// bShow: 是否显示.
func (a *App) X显示边界(是否显示 bool) int {
	return 炫彩基类.X显示边界(是否显示)
}

// 通知消息_弹出, 未实现, 预留接口.
//
// position: 位置, Position_Flag_.
//
// pTitle: 标题.
//
// pText: 内容.
//
// hIcon: 图标.
//
// skin: 外观类型, NotifyMsg_Skin_.
func (a *App) X通知消息_弹出_未实现(位置 炫彩常量类.Position_Flag_, 标题, 内容 string, 图标 int, 外观类型 炫彩常量类.NotifyMsg_Skin_) int {
	return 炫彩基类.X通知消息_弹出(位置, 标题, 内容, 图标, 外观类型)
}

// 通知消息_弹出扩展, 未实现, 预留接口.
//
// position: 位置, Position_Flag_.
//
// pTitle: 标题.
//
// pText: 内容.
//
// hIcon: 图标.
//
// skin: 外观类型, NotifyMsg_Skin_.
//
// bBtnClose: 是否启用关闭按钮.
//
// bAutoClose: 是否自动关闭.
//
// nWidth: 自定义宽度, -1(使用默认值).
//
// nHeight: 自定义高度, -1(使用默认值).
func (a *App) X通知消息_弹出EX_未实现(位置 炫彩常量类.Position_Flag_, 标题, 内容 string, 图标 int, 外观类型 炫彩常量类.NotifyMsg_Skin_, 是否启用关闭按钮, 是否自动关闭 bool, 自定义宽度, 自定义高度 int) int {
	return 炫彩基类.X通知消息_弹出EX(位置, 标题, 内容, 图标, 外观类型, 是否启用关闭按钮, 是否自动关闭, 自定义宽度, 自定义高度)
}

// 通知消息_置持续时间.
//
// hWindow: 通知消息所属窗口句柄, 如果未指定那么认为是桌面通知消息.
//
// duration: 持续时间.
func (a *App) X通知消息_置持续时间(窗口句柄, 持续时间 int) int {
	return 炫彩基类.X通知消息_置持续时间(窗口句柄, 持续时间)
}

// 通知消息_置父边距 设置通知消息与父对象的四边间隔.
//
// hWindow: 通知消息所属窗口句柄, 如果未指定那么认为是桌面通知消息.
//
// left: 左侧间隔, 未实现, 预留功能.
//
// top: 顶部间隔.
//
// right: 右侧间隔.
//
// bottom: 底部间隔, 未实现, 预留功能.
func (a *App) X通知消息_置父边距(窗口句柄, 左侧间隔, 顶部间隔, 右侧间隔, 底部间隔 int) int {
	return 炫彩基类.X通知消息_置父边距(窗口句柄, 左侧间隔, 顶部间隔, 右侧间隔, 底部间隔)
}

// 通知消息_置标题高度.
//
// hWindow: 通知消息所属窗口句柄, 如果未指定那么认为是桌面通知消息.
//
// nHeight: 高度.
func (a *App) X通知消息_置标题高度(窗口句柄, 高度 int) int {
	return 炫彩基类.X通知消息_置标题高度(窗口句柄, 高度)
}

// 通知消息_置宽度, 设置默认宽度.
//
// hWindow: 通知消息所属窗口句柄, 如果未指定那么认为是桌面通知消息.
//
// nWidth: 宽度.
func (a *App) X通知消息_置宽度(窗口句柄, 宽度 int) int {
	return 炫彩基类.X通知消息_置宽度(窗口句柄, 宽度)
}

// 通知消息_置间隔.
//
// hWindow: 通知消息所属窗口句柄, 如果未指定那么认为是桌面通知消息.
//
// nSpace: 间隔大小.
func (a *App) X通知消息_置间隔(窗口句柄, 间隔大小 int) int {
	return 炫彩基类.X通知消息_置间隔(窗口句柄, 间隔大小)
}

// 通知消息_置边大小, 设置通知消息面板边大小.
//
// hWindow: 通知消息所属窗口句柄, 如果未指定那么认为是桌面通知消息.
//
// left: 左边.
//
// top: 顶边.
//
// right: 右边.
//
// bottom: 底边.
func (a *App) X通知消息_置边大小(窗口句柄, 左边, 顶边, 右边, 底边 int) int {
	return 炫彩基类.X通知消息_置边大小(窗口句柄, 左边, 顶边, 右边, 底边)
}

// 炫彩_启用自动DPI. 当启用后, 创建窗口时自动检测DPI调整UI缩放, 处理DPI改变消息; 禁用后,当DPI改变,需要手动设置窗口DPI.
//
// bEnabel: 是否启用.
func (a *App) X启用自动DPI(是否启用 bool) int {
	return 炫彩基类.X启用自动DPI(是否启用)
}

// 炫彩_置窗口图标. 全局窗口图标, 所有未设置图标的窗口, 都将使用此默认图标.
//
// hImage: 图标句柄.
func (a *App) X置窗口图标(图标句柄 int) int {
	return 炫彩基类.X置窗口图标(图标句柄)
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
func (a *App) X启用DPI(是否启用 bool) bool {
	return 炫彩基类.X启用DPI(是否启用)
}
