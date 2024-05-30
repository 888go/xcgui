# 备注开始
# **_方法.md 文件备注:
# ff= 方法,重命名方法名称
# 如://ff:取文本
#
# yx=true,此方法优先翻译
# 如: //yx=true

# **_package.md 文件备注:
# bm= 包名,更换新的包名称 
# 如: package gin //bm:gin类

# **_其他.md 文件备注:
# qm= 前面,跳转到前面进行重命名.文档内如果有多个相同的,会一起重命名.
# hm= 后面,跳转到后面进行重命名.文档内如果有多个相同的,会一起重命名.
# cz= 查找,配合前面/后面使用,
# 如: type Regexp struct {//qm:正则 cz:Regexp struct
#
# th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
# 如:
# type Regexp struct {//th:type Regexp222 struct
#
# cf= 重复,用于重命名多次,
# 如: 
# 一个文档内有2个"One(result interface{}) error"需要重命名.
# 但是要注意,多个新名称要保持一致. 如:"X取一条(result interface{})"

# **_追加.md 文件备注:
# 在代码内追加代码,如:
# //zj:
# func (re *Regexp) X取文本() string { 
# re.F.String()
# }
# //zj:
# 备注结束

[func XInitXCGUI(bD2D bool) bool {]
ff=初始化
bD2D=是否启用D2D

[func XRunXCGUI() int {]
ff=运行

[func XExitXCGUI() int {]
ff=退出

[func XC_DebugToFileInfo(pInfo string) {]
ff=输出调试信息到文件
pInfo=文本

[func XC_SetActivateTopWindow() bool {]
ff=激活窗口

[func XC_GetDefaultFont() int {]
ff=取默认字体

[func XC_MessageBox(pTitle, pText string, nFlags xcc.MessageBox_Flag_, hWndParent uintptr, XCStyle xcc.Window_Style_) xcc.MessageBox_Flag_ {]
ff=消息框
nFlags=标识
pText=内容文本
pTitle=标题

[func XMsg_Create(pTitle, pText string, nFlags xcc.MessageBox_Flag_, hWndParent uintptr, XCStyle xcc.Window_Style_) int {]
ff=消息框_创建
nFlags=标识
pText=内容文本
pTitle=标题

[func XMsg_CreateEx(dwExStyle int, dwStyle int, lpClassName, pTitle, pText string, nFlags xcc.MessageBox_Flag_, hWndParent uintptr, XCStyle xcc.Window_Style_) int {]
ff=消息框_创建EX
nFlags=标识
pText=内容文本
pTitle=标题
lpClassName=窗口类名
dwStyle=窗口样式
dwExStyle=窗口扩展样式

[func XC_SendMessage(hWindow int, msg uint32, wParam, lParam uint) uint {]
ff=发送窗口消息
lParam=参数2
wParam=参数1
msg=消息值
hWindow=窗口句柄

[func XC_PostMessage(hWindow int, msg uint32, wParam int32, lParam int32) bool {]
ff=投递窗口消息
lParam=参数2
wParam=参数1
msg=消息值
hWindow=窗口句柄

[func XC_CallUiThread(pCall func(data int) int, data int) int {]
ff=调用界面线程
pCall=回调函数

[func XC_IsHELE(hEle int) bool {]
ff=判断元素
hEle=元素句柄

[func XC_IsHWINDOW(hWindow int) bool {]
ff=判断窗口
hWindow=窗口句柄

[func XC_IsShape(hShape int) bool {]
ff=判断形状对象
hShape=形状对象句柄

[func XC_IsHXCGUI(hXCGUI int, nType xcc.XC_OBJECT_TYPE) bool {]
ff=判断句柄包含类型
nType=句柄类型
hXCGUI=炫彩句柄

[func XC_hWindowFromHWnd(hWnd uintptr) int {]
ff=转换HWND到HWINDOW
hWnd=窗口句柄HWND

[func XC_SetProperty(hXCGUI int, pName string, pValue string) bool {]
ff=置属性
pValue=属性值
pName=属性名
hXCGUI=对象句柄

[func XC_GetProperty(hXCGUI int, pName string) string {]
ff=取属性
pName=属性名
hXCGUI=对象句柄

[func XC_RegisterWindowClassName(pClassName string) bool {]
ff=注册窗口类名
pClassName=类名

[func XC_IsSViewExtend(hEle int) bool {]
ff=判断滚动视图EX元素
hEle=元素句柄

[func XC_GetObjectType(hXCGUI int) xcc.XC_OBJECT_TYPE {]
ff=取对象类型
hXCGUI=炫彩对象句柄

[func XC_GetObjectByID(hWindow int, nID int) int {]
ff=取对象从ID
nID=ID值
hWindow=所属窗口句柄

[func XC_GetObjectByIDName(hWindow int, pName string) int {]
ff=取对象从ID名称
pName=ID名称
hWindow=所属窗口句柄

[func XC_GetObjectByUID(nUID int) int {]
ff=取对象从UID
nUID=UID值

[func XC_GetObjectByUIDName(pName string) int {]
ff=取对象从UID名称
pName=UID名称

[func XC_GetObjectByName(pName string) int {]
ff=取对象从名称
pName=名称

[func XC_SetPaintFrequency(nMilliseconds int) {]
ff=置绘制频率
nMilliseconds=重绘最小时间间隔

[func XC_SetTextRenderingHint(nType int) {]
ff=置文本渲染质量

[func XC_EnableGdiDrawText(bEnable bool) {]
ff=启用GDI绘制文本
bEnable=是否启用

[func XC_RectInRect(pRect1 *RECT, pRect2 *RECT) bool {]
ff=判断矩形相交
pRect2=矩形2
pRect1=矩形1

[func XC_CombineRect(pDest *RECT, pSrc1 *RECT, pSrc2 *RECT) {]
ff=组合矩形
pSrc2=源矩形2
pSrc1=源矩形1
pDest=新的矩形区域

[func XC_ShowLayoutFrame(bShow bool) {]
ff=显示布局边界
bShow=是否显示

[func XC_EnableDebugFile(bEnable bool) {]
ff=启用debug文件
bEnable=是否启用

[func XC_EnableResMonitor(bEnable bool) {]
ff=启用资源监视器
bEnable=是否启用

[func XC_SetLayoutFrameColor(color int) int {]
ff=置布局边界颜色
color=ABGR颜色值

[func XC_EnableErrorMessageBox(bEnabel bool) int {]
ff=启用错误弹窗
bEnabel=是否启用

[func XC_EnableAutoExitApp(bEnabel bool) int {]
ff=启用自动退出程序
bEnabel=是否启用

[func XC_GetTextSize(pString string, length int, hFontX int, pOutSize *SIZE) int {]
ff=取文本绘制大小
pOutSize=接收返回大小
hFontX=字体
length=字符串长度
pString=字符串

[func XC_GetTextShowSize(pString string, length int, hFontX int, pOutSize *SIZE) int {]
ff=取文本显示大小
pOutSize=接收返回大小
hFontX=字体
length=字符串长度
pString=字符串

[func XC_GetTextShowSizeEx(pString string, length int, hFontX int, nTextAlign xcc.TextFormatFlag_, pOutSize *SIZE) int {]
ff=取文本显示大小EX
nTextAlign=文本对齐方式
hFontX=字体
length=字符串长度
pString=字符串

[func XC_GetTextShowRect(pString string, length int, hFontX int, nTextAlign xcc.TextFormatFlag_, width int, pOutSize *SIZE) int {]
ff=炫彩_取文本显示矩形
nTextAlign=文本对齐
hFontX=字体
length=字符串长度
pString=字符串

[func XC_SetDefaultFont(hFontX int) int {]
ff=置默认字体
hFontX=炫彩字体句柄

[func XC_AddFileSearchPath(pPath string) int {]
ff=添加搜索路径
pPath=文件夹

[func XC_InitFont(pFont *LOGFONTW, pName string, size int, bBold bool, bItalic bool, bUnderline bool, bStrikeOut bool) int {]
ff=初始化字体
bStrikeOut=是否有删除线
bUnderline=是否有下划线
bItalic=是否为斜体
bBold=是否为粗体
size=字体大小
pName=字体名称
pFont=LOGFONTW结构体指针

[func XC_Malloc(size int) int {]
ff=分配内存
size=大小

[func XC_Free(p int) int {]
ff=释放内存
p=内存首地址

[func XC_Alert(pTitle, pText string) int {]
ff=弹框
pText=提示内容
pTitle=提示框标题

[func XC_Sys_ShellExecute(hwnd uintptr, lpOperation string, lpFile string, lpParameters string, lpDirectory string, nShowCmd xcc.SW_) uintptr {]
ff=对指定文件执行操作
nShowCmd=显示Cmd
lpDirectory=默认完整路径
lpParameters=参数
lpFile=文件名
lpOperation=操作类型
hwnd=父窗口句柄

[func XC_LoadLibrary(lpFileName string) uintptr {]
ff=载入动态库
lpFileName=文件名

[func XC_GetProcAddress(hModule uintptr, lpProcName string) uintptr {]
ff=取动态库中函数地址
lpProcName=函数名
hModule=动态库模块句柄

[func XC_FreeLibrary(hModule uintptr) bool {]
ff=释放动态库
hModule=动态库模块句柄

[func XC_LoadDll(pDllFileName string) uintptr {]
ff=加载DLL
pDllFileName=DLL文件名

[func XC_PostQuitMessage(nExitCode int) int {]
ff=发送WM_QUIT消息退出消息循环
nExitCode=退出码

[func XC_SetD2dTextRenderingMode(mode xcc.XC_DWRITE_RENDERING_MODE_) int {]
ff=置D2D文本渲染模式
mode=渲染模式

[func XC_IsEnableD2D() bool {]
ff=是否启用了D2D

[func XC_wtoa(pValue string) uintptr {]
ff=W2A
pValue=参数

[func XC_atow(pValue uintptr) string {]
ff=A2W
pValue=参数

[func XC_utf8tow(pUtf8 uintptr) string {]
ff=UTF8到文本W
pUtf8=参数

[func XC_utf8towEx(pUtf8 uintptr, length int) string {]
ff=UTF8到文本WEX
length=utf8字符串长度
pUtf8=utf8字符串指针

[func XC_utf8toa(pUtf8 uintptr) uintptr {]
ff=UTF8到文本A
pUtf8=utf8字符串指针

[func XC_atoutf8(pValue uintptr) uintptr {]
ff=文本A到UTF8
pValue=参数

[func XC_wtoutf8(pValue string) uintptr {]
ff=文本W到UTF8
pValue=字符串

[func XC_wtoutf8Ex(pValue string, length int) uintptr {]
ff=文本W到UTF8EX
length=字符串长度
pValue=字符串

[func XC_UnicodeToAnsi(pIn string, inLen int, pOut uintptr, outLen int) int {]
ff=U2A
outLen=pOut缓冲区大小
pOut=转换后缓冲区指针
inLen=pIn字符数量
pIn=待转换的Unicode字符串

[func XC_AnsiToUnicode(pIn uintptr, inLen int, pOut *string, outLen int) int {]
ff=A2U
outLen=pOut缓冲区大小
pOut=转换后缓冲区指针
inLen=pIn字符数量
pIn=指向待转换的Ansi字符串指针

[func XDebug_Print(level int, pInfo string) int {]
ff=打印调试信息
pInfo=信息
level=级别

[func XC_ShowSvgFrame(bShow bool) int {]
ff=显示边界
bShow=是否显示

[func XC_EnableAutoDPI(bEnabel bool) int {]
ff=启用自动DPI
bEnabel=是否启用

[func XC_EnableDPI(bEnabel bool) bool {]
ff=启用DPI
bEnabel=是否启用

[func XC_SetWindowIcon(hImage int) int {]
ff=置窗口图标
hImage=图标句柄
