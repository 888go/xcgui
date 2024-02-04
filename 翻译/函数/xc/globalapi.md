
# <翻译开始>
func XInitXCGUI(bD2D
是否启用D2D
# <翻译结束>

# <翻译开始>
func XInitXCGUI
X初始化
# <翻译结束>


# <翻译开始>
func XRunXCGUI
X运行
# <翻译结束>


# <翻译开始>
func XExitXCGUI
X退出
# <翻译结束>


# <翻译开始>
func XC_DebugToFileInfo(pInfo
文本
# <翻译结束>

# <翻译开始>
func XC_DebugToFileInfo
X输出调试信息到文件
# <翻译结束>


# <翻译开始>
func XC_SetActivateTopWindow
X激活窗口
# <翻译结束>


# <翻译开始>
func XC_GetDefaultFont
X取默认字体
# <翻译结束>


# <翻译开始>
func XC_MessageBox(pTitle, pText string, nFlags xcc.MessageBox_Flag_, hWndParent uintptr, XCStyle
样式
# <翻译结束>

# <翻译开始>
func XC_MessageBox(pTitle, pText string, nFlags xcc.MessageBox_Flag_, hWndParent
父窗口句柄
# <翻译结束>

# <翻译开始>
func XC_MessageBox(pTitle, pText string, nFlags
标识
# <翻译结束>

# <翻译开始>
func XC_MessageBox(pTitle, pText
内容文本
# <翻译结束>

# <翻译开始>
func XC_MessageBox(pTitle
标题
# <翻译结束>

# <翻译开始>
func XC_MessageBox
X消息框
# <翻译结束>


# <翻译开始>
func XMsg_Create(pTitle, pText string, nFlags xcc.MessageBox_Flag_, hWndParent uintptr, XCStyle
样式
# <翻译结束>

# <翻译开始>
func XMsg_Create(pTitle, pText string, nFlags xcc.MessageBox_Flag_, hWndParent
父窗口句柄
# <翻译结束>

# <翻译开始>
func XMsg_Create(pTitle, pText string, nFlags
标识
# <翻译结束>

# <翻译开始>
func XMsg_Create(pTitle, pText
内容文本
# <翻译结束>

# <翻译开始>
func XMsg_Create(pTitle
标题
# <翻译结束>

# <翻译开始>
func XMsg_Create
X消息框_创建
# <翻译结束>


# <翻译开始>
func XMsg_CreateEx(dwExStyle int, dwStyle int, lpClassName, pTitle, pText string, nFlags xcc.MessageBox_Flag_, hWndParent uintptr, XCStyle
样式
# <翻译结束>

# <翻译开始>
func XMsg_CreateEx(dwExStyle int, dwStyle int, lpClassName, pTitle, pText string, nFlags xcc.MessageBox_Flag_, hWndParent
父窗口句柄
# <翻译结束>

# <翻译开始>
func XMsg_CreateEx(dwExStyle int, dwStyle int, lpClassName, pTitle, pText string, nFlags
标识
# <翻译结束>

# <翻译开始>
func XMsg_CreateEx(dwExStyle int, dwStyle int, lpClassName, pTitle, pText
内容文本
# <翻译结束>

# <翻译开始>
func XMsg_CreateEx(dwExStyle int, dwStyle int, lpClassName, pTitle
标题
# <翻译结束>

# <翻译开始>
func XMsg_CreateEx(dwExStyle int, dwStyle int, lpClassName
窗口类名
# <翻译结束>

# <翻译开始>
func XMsg_CreateEx(dwExStyle int, dwStyle
窗口样式
# <翻译结束>

# <翻译开始>
func XMsg_CreateEx(dwExStyle
窗口EX样式
# <翻译结束>

# <翻译开始>
func XMsg_CreateEx
X消息框_创建EX
# <翻译结束>


# <翻译开始>
func XC_SendMessage(hWindow int, msg uint32, wParam, lParam
参数2
# <翻译结束>

# <翻译开始>
func XC_SendMessage(hWindow int, msg uint32, wParam
参数1
# <翻译结束>

# <翻译开始>
func XC_SendMessage(hWindow int, msg
消息值
# <翻译结束>

# <翻译开始>
func XC_SendMessage(hWindow
窗口句柄
# <翻译结束>

# <翻译开始>
func XC_SendMessage
X发送窗口消息
# <翻译结束>


# <翻译开始>
func XC_PostMessage(hWindow int, msg uint32, wParam int32, lParam
参数2
# <翻译结束>

# <翻译开始>
func XC_PostMessage(hWindow int, msg uint32, wParam
参数1
# <翻译结束>

# <翻译开始>
func XC_PostMessage(hWindow int, msg
消息值
# <翻译结束>

# <翻译开始>
func XC_PostMessage(hWindow
窗口句柄
# <翻译结束>

# <翻译开始>
func XC_PostMessage
X投递窗口消息
# <翻译结束>


# <翻译开始>
func XC_CallUiThread(pCall
回调函数
# <翻译结束>

# <翻译开始>
func XC_CallUiThread
X调用界面线程
# <翻译结束>


# <翻译开始>
func XC_IsHELE(hEle
元素句柄
# <翻译结束>

# <翻译开始>
func XC_IsHELE
X判断元素
# <翻译结束>


# <翻译开始>
func XC_IsHWINDOW(hWindow
窗口句柄
# <翻译结束>

# <翻译开始>
func XC_IsHWINDOW
X判断窗口
# <翻译结束>


# <翻译开始>
func XC_IsShape(hShape
形状对象句柄
# <翻译结束>

# <翻译开始>
func XC_IsShape
X判断形状对象
# <翻译结束>


# <翻译开始>
func XC_IsHXCGUI(hXCGUI int, nType
句柄类型
# <翻译结束>

# <翻译开始>
func XC_IsHXCGUI(hXCGUI
炫彩句柄
# <翻译结束>

# <翻译开始>
func XC_IsHXCGUI
X判断句柄包含类型
# <翻译结束>


# <翻译开始>
func XC_hWindowFromHWnd(hWnd
窗口句柄HWND
# <翻译结束>

# <翻译开始>
func XC_hWindowFromHWnd
X转换HWND到HWINDOW
# <翻译结束>


# <翻译开始>
func XC_SetProperty(hXCGUI int, pName string, pValue
属性值
# <翻译结束>

# <翻译开始>
func XC_SetProperty(hXCGUI int, pName
属性名
# <翻译结束>

# <翻译开始>
func XC_SetProperty(hXCGUI
对象句柄
# <翻译结束>

# <翻译开始>
func XC_SetProperty
X置属性
# <翻译结束>


# <翻译开始>
func XC_GetProperty(hXCGUI int, pName
属性名
# <翻译结束>

# <翻译开始>
func XC_GetProperty(hXCGUI
对象句柄
# <翻译结束>

# <翻译开始>
func XC_GetProperty
X取属性
# <翻译结束>


# <翻译开始>
func XC_RegisterWindowClassName(pClassName
类名
# <翻译结束>

# <翻译开始>
func XC_RegisterWindowClassName
X注册窗口类名
# <翻译结束>


# <翻译开始>
func XC_IsSViewExtend(hEle
元素句柄
# <翻译结束>

# <翻译开始>
func XC_IsSViewExtend
X判断滚动视图EX元素
# <翻译结束>


# <翻译开始>
func XC_GetObjectType(hXCGUI
炫彩对象句柄
# <翻译结束>

# <翻译开始>
func XC_GetObjectType
X取对象类型
# <翻译结束>


# <翻译开始>
func XC_GetObjectByID(hWindow int, nID
ID值
# <翻译结束>

# <翻译开始>
func XC_GetObjectByID(hWindow
所属窗口句柄
# <翻译结束>

# <翻译开始>
func XC_GetObjectByID
X取对象从ID
# <翻译结束>


# <翻译开始>
func XC_GetObjectByIDName(hWindow int, pName
ID名称
# <翻译结束>

# <翻译开始>
func XC_GetObjectByIDName(hWindow
所属窗口句柄
# <翻译结束>

# <翻译开始>
func XC_GetObjectByIDName
X取对象从ID名称
# <翻译结束>


# <翻译开始>
func XC_GetObjectByUID(nUID
UID值
# <翻译结束>

# <翻译开始>
func XC_GetObjectByUID
X取对象从UID
# <翻译结束>


# <翻译开始>
func XC_GetObjectByUIDName(pName
UID名称
# <翻译结束>

# <翻译开始>
func XC_GetObjectByUIDName
X取对象从UID名称
# <翻译结束>


# <翻译开始>
func XC_GetObjectByName(pName
名称
# <翻译结束>

# <翻译开始>
func XC_GetObjectByName
X取对象从名称
# <翻译结束>


# <翻译开始>
func XC_SetPaintFrequency(nMilliseconds
重绘最小时间间隔
# <翻译结束>

# <翻译开始>
func XC_SetPaintFrequency
X置绘制频率
# <翻译结束>


# <翻译开始>
func XC_SetTextRenderingHint
X置文本渲染质量
# <翻译结束>


# <翻译开始>
func XC_EnableGdiDrawText(bEnable
是否启用
# <翻译结束>

# <翻译开始>
func XC_EnableGdiDrawText
X启用GDI绘制文本
# <翻译结束>


# <翻译开始>
func XC_RectInRect(pRect1 *RECT, pRect2
矩形2
# <翻译结束>

# <翻译开始>
func XC_RectInRect(pRect1
矩形1
# <翻译结束>

# <翻译开始>
func XC_RectInRect
X判断矩形相交
# <翻译结束>


# <翻译开始>
func XC_CombineRect(pDest *RECT, pSrc1 *RECT, pSrc2
源矩形2
# <翻译结束>

# <翻译开始>
func XC_CombineRect(pDest *RECT, pSrc1
源矩形1
# <翻译结束>

# <翻译开始>
func XC_CombineRect(pDest
新的矩形区域
# <翻译结束>

# <翻译开始>
func XC_CombineRect
X组合矩形
# <翻译结束>


# <翻译开始>
func XC_ShowLayoutFrame(bShow
是否显示
# <翻译结束>

# <翻译开始>
func XC_ShowLayoutFrame
X显示布局边界
# <翻译结束>


# <翻译开始>
func XC_EnableDebugFile(bEnable
是否启用
# <翻译结束>

# <翻译开始>
func XC_EnableDebugFile
X启用debug文件
# <翻译结束>


# <翻译开始>
func XC_EnableResMonitor(bEnable
是否启用
# <翻译结束>

# <翻译开始>
func XC_EnableResMonitor
X启用资源监视器
# <翻译结束>


# <翻译开始>
func XC_SetLayoutFrameColor(color
ABGR颜色值
# <翻译结束>

# <翻译开始>
func XC_SetLayoutFrameColor
X置布局边界颜色
# <翻译结束>


# <翻译开始>
func XC_EnableErrorMessageBox(bEnabel
是否启用
# <翻译结束>

# <翻译开始>
func XC_EnableErrorMessageBox
X启用错误弹窗
# <翻译结束>


# <翻译开始>
func XC_EnableAutoExitApp(bEnabel
是否启用
# <翻译结束>

# <翻译开始>
func XC_EnableAutoExitApp
X启用自动退出程序
# <翻译结束>


# <翻译开始>
func XC_GetTextSize(pString string, length int, hFontX int, pOutSize
接收返回大小
# <翻译结束>

# <翻译开始>
func XC_GetTextSize(pString string, length int, hFontX
字体
# <翻译结束>

# <翻译开始>
func XC_GetTextSize(pString string, length
字符串长度
# <翻译结束>

# <翻译开始>
func XC_GetTextSize(pString
字符串
# <翻译结束>

# <翻译开始>
func XC_GetTextSize
X取文本绘制大小
# <翻译结束>


# <翻译开始>
func XC_GetTextShowSize(pString string, length int, hFontX int, pOutSize
接收返回大小
# <翻译结束>

# <翻译开始>
func XC_GetTextShowSize(pString string, length int, hFontX
字体
# <翻译结束>

# <翻译开始>
func XC_GetTextShowSize(pString string, length
字符串长度
# <翻译结束>

# <翻译开始>
func XC_GetTextShowSize(pString
字符串
# <翻译结束>

# <翻译开始>
func XC_GetTextShowSize
X取文本显示大小
# <翻译结束>


# <翻译开始>
func XC_GetTextShowSizeEx(pString string, length int, hFontX int, nTextAlign xcc.TextFormatFlag_, pOutSize
接收返回大小
# <翻译结束>

# <翻译开始>
func XC_GetTextShowSizeEx(pString string, length int, hFontX int, nTextAlign
文本对齐方式
# <翻译结束>

# <翻译开始>
func XC_GetTextShowSizeEx(pString string, length int, hFontX
字体
# <翻译结束>

# <翻译开始>
func XC_GetTextShowSizeEx(pString string, length
字符串长度
# <翻译结束>

# <翻译开始>
func XC_GetTextShowSizeEx(pString
字符串
# <翻译结束>

# <翻译开始>
func XC_GetTextShowSizeEx
X取文本显示大小EX
# <翻译结束>


# <翻译开始>
func XC_GetTextShowRect(pString string, length int, hFontX int, nTextAlign xcc.TextFormatFlag_, width int, pOutSize
接收返回大小
# <翻译结束>

# <翻译开始>
func XC_GetTextShowRect(pString string, length int, hFontX int, nTextAlign xcc.TextFormatFlag_, width
最大宽度
# <翻译结束>

# <翻译开始>
func XC_GetTextShowRect(pString string, length int, hFontX int, nTextAlign
文本对齐
# <翻译结束>

# <翻译开始>
func XC_GetTextShowRect(pString string, length int, hFontX
字体
# <翻译结束>

# <翻译开始>
func XC_GetTextShowRect(pString string, length
字符串长度
# <翻译结束>

# <翻译开始>
func XC_GetTextShowRect(pString
字符串
# <翻译结束>

# <翻译开始>
func XC_GetTextShowRect
X炫彩_取文本显示矩形
# <翻译结束>


# <翻译开始>
func XC_SetDefaultFont(hFontX
炫彩字体句柄
# <翻译结束>

# <翻译开始>
func XC_SetDefaultFont
X置默认字体
# <翻译结束>


# <翻译开始>
func XC_AddFileSearchPath(pPath
文件夹
# <翻译结束>

# <翻译开始>
func XC_AddFileSearchPath
X添加搜索路径
# <翻译结束>


# <翻译开始>
func XC_InitFont(pFont *LOGFONTW, pName string, size int, bBold bool, bItalic bool, bUnderline bool, bStrikeOut
是否有删除线
# <翻译结束>

# <翻译开始>
func XC_InitFont(pFont *LOGFONTW, pName string, size int, bBold bool, bItalic bool, bUnderline
是否有下划线
# <翻译结束>

# <翻译开始>
func XC_InitFont(pFont *LOGFONTW, pName string, size int, bBold bool, bItalic
是否为斜体
# <翻译结束>

# <翻译开始>
func XC_InitFont(pFont *LOGFONTW, pName string, size int, bBold
是否为粗体
# <翻译结束>

# <翻译开始>
func XC_InitFont(pFont *LOGFONTW, pName string, size
字体大小
# <翻译结束>

# <翻译开始>
func XC_InitFont(pFont *LOGFONTW, pName
字体名称
# <翻译结束>

# <翻译开始>
func XC_InitFont(pFont
LOGFONTW结构体指针
# <翻译结束>

# <翻译开始>
func XC_InitFont
X初始化字体
# <翻译结束>


# <翻译开始>
func XC_Malloc(size
大小
# <翻译结束>

# <翻译开始>
func XC_Malloc
X分配内存
# <翻译结束>


# <翻译开始>
func XC_Free(p
内存首地址
# <翻译结束>

# <翻译开始>
func XC_Free
X释放内存
# <翻译结束>


# <翻译开始>
func XC_Alert(pTitle, pText
提示内容
# <翻译结束>

# <翻译开始>
func XC_Alert(pTitle
提示框标题
# <翻译结束>

# <翻译开始>
func XC_Alert
X弹框
# <翻译结束>


# <翻译开始>
func XC_Sys_ShellExecute(hwnd uintptr, lpOperation string, lpFile string, lpParameters string, lpDirectory string, nShowCmd
显示Cmd
# <翻译结束>

# <翻译开始>
func XC_Sys_ShellExecute(hwnd uintptr, lpOperation string, lpFile string, lpParameters string, lpDirectory
默认完整路径
# <翻译结束>

# <翻译开始>
func XC_Sys_ShellExecute(hwnd uintptr, lpOperation string, lpFile string, lpParameters
参数
# <翻译结束>

# <翻译开始>
func XC_Sys_ShellExecute(hwnd uintptr, lpOperation string, lpFile
文件名
# <翻译结束>

# <翻译开始>
func XC_Sys_ShellExecute(hwnd uintptr, lpOperation
操作类型
# <翻译结束>

# <翻译开始>
func XC_Sys_ShellExecute(hwnd
父窗口句柄
# <翻译结束>

# <翻译开始>
func XC_Sys_ShellExecute
X对指定文件执行操作
# <翻译结束>


# <翻译开始>
func XC_LoadLibrary(lpFileName
文件名
# <翻译结束>

# <翻译开始>
func XC_LoadLibrary
X载入动态库
# <翻译结束>


# <翻译开始>
func XC_GetProcAddress(hModule uintptr, lpProcName
函数名
# <翻译结束>

# <翻译开始>
func XC_GetProcAddress(hModule
动态库模块句柄
# <翻译结束>

# <翻译开始>
func XC_GetProcAddress
X取动态库中函数地址
# <翻译结束>


# <翻译开始>
func XC_FreeLibrary(hModule
动态库模块句柄
# <翻译结束>

# <翻译开始>
func XC_FreeLibrary
X释放动态库
# <翻译结束>


# <翻译开始>
func XC_LoadDll(pDllFileName
DLL文件名
# <翻译结束>

# <翻译开始>
func XC_LoadDll
X加载DLL
# <翻译结束>


# <翻译开始>
func XC_PostQuitMessage(nExitCode
退出码
# <翻译结束>

# <翻译开始>
func XC_PostQuitMessage
X发送WM_QUIT消息退出消息循环
# <翻译结束>


# <翻译开始>
func XC_SetD2dTextRenderingMode(mode
渲染模式
# <翻译结束>

# <翻译开始>
func XC_SetD2dTextRenderingMode
X置D2D文本渲染模式
# <翻译结束>


# <翻译开始>
func XC_IsEnableD2D
X是否启用了D2D
# <翻译结束>


# <翻译开始>
func XC_wtoa(pValue
参数
# <翻译结束>

# <翻译开始>
func XC_wtoa
W2A
# <翻译结束>


# <翻译开始>
func XC_atow(pValue
参数
# <翻译结束>

# <翻译开始>
func XC_atow
A2W
# <翻译结束>


# <翻译开始>
func XC_utf8tow(pUtf8
参数
# <翻译结束>

# <翻译开始>
func XC_utf8tow
UTF8到文本W
# <翻译结束>


# <翻译开始>
func XC_utf8towEx(pUtf8 uintptr, length
utf8字符串长度
# <翻译结束>

# <翻译开始>
func XC_utf8towEx(pUtf8
utf8字符串指针
# <翻译结束>

# <翻译开始>
func XC_utf8towEx
UTF8到文本WEX
# <翻译结束>


# <翻译开始>
func XC_utf8toa(pUtf8
utf8字符串指针
# <翻译结束>

# <翻译开始>
func XC_utf8toa
UTF8到文本A
# <翻译结束>


# <翻译开始>
func XC_atoutf8(pValue
参数
# <翻译结束>

# <翻译开始>
func XC_atoutf8
X文本A到UTF8
# <翻译结束>


# <翻译开始>
func XC_wtoutf8(pValue
字符串
# <翻译结束>

# <翻译开始>
func XC_wtoutf8
X文本W到UTF8
# <翻译结束>


# <翻译开始>
func XC_wtoutf8Ex(pValue string, length
字符串长度
# <翻译结束>

# <翻译开始>
func XC_wtoutf8Ex(pValue
字符串
# <翻译结束>

# <翻译开始>
func XC_wtoutf8Ex
X文本W到UTF8EX
# <翻译结束>


# <翻译开始>
func XC_UnicodeToAnsi(pIn string, inLen int, pOut uintptr, outLen
pOut缓冲区大小
# <翻译结束>

# <翻译开始>
func XC_UnicodeToAnsi(pIn string, inLen int, pOut
指向接收转换后的Ansi字符串缓冲区指针
# <翻译结束>

# <翻译开始>
func XC_UnicodeToAnsi(pIn string, inLen
pIn字符数量
# <翻译结束>

# <翻译开始>
func XC_UnicodeToAnsi(pIn
待转换的Unicode字符串
# <翻译结束>

# <翻译开始>
func XC_UnicodeToAnsi
U2A
# <翻译结束>


# <翻译开始>
func XC_AnsiToUnicode(pIn uintptr, inLen int, pOut *string, outLen
pOut缓冲区大小
# <翻译结束>

# <翻译开始>
func XC_AnsiToUnicode(pIn uintptr, inLen int, pOut
指向接收转换后的Unicode字符串缓冲区指针
# <翻译结束>

# <翻译开始>
func XC_AnsiToUnicode(pIn uintptr, inLen
pIn字符数量
# <翻译结束>

# <翻译开始>
func XC_AnsiToUnicode(pIn
指向待转换的Ansi字符串指针
# <翻译结束>

# <翻译开始>
func XC_AnsiToUnicode
A2U
# <翻译结束>


# <翻译开始>
func XDebug_Print(level int, pInfo
信息
# <翻译结束>

# <翻译开始>
func XDebug_Print(level
级别
# <翻译结束>

# <翻译开始>
func XDebug_Print
X打印调试信息
# <翻译结束>


# <翻译开始>
func XC_ShowSvgFrame(bShow
是否显示
# <翻译结束>

# <翻译开始>
func XC_ShowSvgFrame
X显示边界
# <翻译结束>


# <翻译开始>
func XC_EnableAutoDPI(bEnabel
是否启用
# <翻译结束>

# <翻译开始>
func XC_EnableAutoDPI
X启用自动DPI
# <翻译结束>


# <翻译开始>
func XC_EnableDPI(bEnabel
是否启用
# <翻译结束>

# <翻译开始>
func XC_EnableDPI
X启用DPI
# <翻译结束>


# <翻译开始>
func XC_SetWindowIcon(hImage
图标句柄
# <翻译结束>

# <翻译开始>
func XC_SetWindowIcon
X置窗口图标
# <翻译结束>

