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
# zz= 正则查找,配合前面/后面使用, 有设置正则查找,就不用设置上面的查找
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
# //zj:前面一行的代码,如果为空,追加到末尾行
# func (re *Regexp) X取文本() string { 
# re.F.String()
# }
# //zj:
# 备注结束

[func New(bD2D bool) *App {]
ff=创建
bD2D=是否启用D2D

[func (a *App) Run() int {]
ff=运行

[func (a *App) Exit() int {]
ff=退出

[func (a *App) ShowAndRun(hWindow int) {]
ff=显示窗口并运行
hWindow=炫彩窗口句柄

[func (a *App) DebugToFileInfo(pInfo string) {]
ff=输出调试信息到文件
pInfo=文本

[func (a *App) SetActivateTopWindow() bool {]
ff=激活窗口

[func (a *App) GetDefaultFont() int {]
ff=取默认字体

[func (a *App) MessageBox(pTitle, pText string, nFlags xcc.MessageBox_Flag_, hWndParent uintptr, XCStyle xcc.Window_Style_) xcc.MessageBox_Flag_ {]
ff=消息框
nFlags=标识
pText=内容文本
pTitle=标题

[func (a *App) Msg_Create(pTitle, pText string, nFlags xcc.MessageBox_Flag_, hWndParent uintptr, XCStyle xcc.Window_Style_) int {]
ff=消息框_创建
nFlags=标识
pText=内容文本
pTitle=标题

[func (a *App) Msg_CreateEx(dwExStyle int, dwStyle int, lpClassName string, pTitle, pText string, nFlags xcc.MessageBox_Flag_, hWndParent uintptr, XCStyle xcc.Window_Style_) int {]
ff=消息框_创建EX
nFlags=标识
pText=内容文本
pTitle=标题
lpClassName=窗口类名
dwStyle=窗口样式
dwExStyle=窗口样式EX

[func (a *App) SendMessage(hWindow int, msg uint32, wParam, lParam uint) uint {]
ff=发送窗口消息
lParam=参数2
wParam=参数1
msg=消息值
hWindow=窗口句柄

[func (a *App) PostMessage(hWindow int, msg uint32, wParam int32, lParam int32) bool {]
ff=投递窗口消息
lParam=参数2
wParam=参数1
msg=消息值
hWindow=窗口句柄

[func (a *App) CallUiThread(pCall func(data int) int, data int) int {]
ff=内部_调用界面线程
pCall=回调函数

[func (a *App) CallUiThreadEx(pCall func(data int) int, data int) int {]
ff=调用界面线程EX
pCall=回调函数

[func (a *App) CallUT(f func()) {]
ff=简易调用界面线程
f=回调函数

[func (a *App) CallUiThreader(u xc.UiThreader, data int) int {]
ff=调用界面线程

[func (a *App) IsHELE(hEle int) bool {]
ff=判断元素
hEle=元素句柄

[func (a *App) IsHWINDOW(hWindow int) bool {]
ff=判断窗口
hWindow=窗口句柄

[func (a *App) IsShape(hShape int) bool {]
ff=判断形状对象
hShape=形状对象句柄

[func (a *App) IsHXCGUI(hXCGUI int, nType xcc.XC_OBJECT_TYPE) bool {]
ff=判断句柄包含类型
nType=句柄类型
hXCGUI=炫彩句柄

[func (a *App) HWindowFromHWnd(hWnd uintptr) int {]
ff=转换HWND到HWINDOW
hWnd=窗口真实句柄HWND

[func (a *App) SetProperty(hXCGUI int, pName string, pValue string) bool {]
ff=置属性
pValue=属性值
pName=属性名
hXCGUI=对象句柄

[func (a *App) GetProperty(hXCGUI int, pName string) string {]
ff=取属性
pName=属性名
hXCGUI=对象句柄

[func (a *App) RegisterWindowClassName(pClassName string) bool {]
ff=注册窗口类名
pClassName=类名

[func (a *App) IsSViewExtend(hEle int) bool {]
ff=判断滚动视图扩展元素
hEle=元素句柄

[func (a *App) GetObjectType(hXCGUI int) xcc.XC_OBJECT_TYPE {]
ff=取对象类型
hXCGUI=炫彩对象句柄

[func (a *App) GetObjectByID(hWindow int, nID int) int {]
ff=取对象从ID
nID=ID值
hWindow=所属窗口句柄

[func (a *App) GetObjectByIDName(hWindow int, pName string) int {]
ff=取对象从ID名称
pName=ID名称
hWindow=所属窗口句柄

[func (a *App) GetObjectByUID(nUID int) int {]
ff=取对象从UID
nUID=UID值

[func (a *App) GetObjectByUIDName(pName string) int {]
ff=取对象从UID名称
pName=UID名称

[func (a *App) GetObjectByName(pName string) int {]
ff=取对象从名称
pName=name名称

[func (a *App) SetPaintFrequency(nMilliseconds int) {]
ff=置绘制频率
nMilliseconds=重绘最小时间间隔

[func (a *App) SetTextRenderingHint(nType int) {]
ff=置文本渲染质量

[func (a *App) EnableGdiDrawText(bEnable bool) {]
ff=启用GDI绘制文本
bEnable=是否启用

[func (a *App) RectInRect(pRect1 *xc.RECT, pRect2 *xc.RECT) bool {]
ff=判断矩形相交
pRect1=矩形1

[func (a *App) CombineRect(pDest *xc.RECT, pSrc1 *xc.RECT, pSrc2 *xc.RECT) {]
ff=组合矩形
pDest=新的矩形区域

[func (a *App) ShowLayoutFrame(bShow bool) {]
ff=显示布局边界
bShow=是否显示

[func (a *App) EnableDebugFile(bEnable bool) {]
ff=启用debug文件
bEnable=是否启用

[func (a *App) EnableResMonitor(bEnable bool) {]
ff=启用资源监视器
bEnable=是否启用

[func (a *App) SetLayoutFrameColor(color int) int {]
ff=置布局边界颜色
color=ABGR

[func (a *App) EnableErrorMessageBox(bEnabel bool) int {]
ff=启用错误弹窗
bEnabel=是否启用

[func (a *App) EnableAutoExitApp(bEnabel bool) int {]
ff=启用自动退出程序
bEnabel=是否启用

[func (a *App) GetTextSize(pString string, length int, hFontX int, pOutSize *xc.SIZE) int {]
ff=取文本绘制大小
pOutSize=接收返回大小
hFontX=字体
length=字符串长度
pString=字符串

[func (a *App) GetTextShowSize(pString string, length int, hFontX int, pOutSize *xc.SIZE) int {]
ff=取文本显示大小
pOutSize=接收返回大小
hFontX=字体
length=字符串长度
pString=字符串

[func (a *App) GetTextShowSizeEx(pString string, length int, hFontX int, nTextAlign xcc.TextFormatFlag_, pOutSize *xc.SIZE) int {]
ff=取文本显示大小EX
nTextAlign=文本对齐方式
hFontX=字体
length=字符串长度
pString=字符串

[func (a *App) GetTextShowRect(pString string, length int, hFontX int, nTextAlign xcc.TextFormatFlag_, width int, pOutSize *xc.SIZE) int {]
ff=取文本显示矩形
nTextAlign=文本对齐
hFontX=字体
length=字符串长度
pString=字符串

[func (a *App) SetDefaultFont(hFontX int) int {]
ff=置默认字体
hFontX=炫彩字体句柄

[func (a *App) AddFileSearchPath(pPath string) int {]
ff=添加搜索路径
pPath=文件夹

[func (a *App) InitFont(pFont *xc.LOGFONTW, pName string, size int, bBold bool, bItalic bool, bUnderline bool, bStrikeOut bool) int {]
ff=初始化字体
pFont=LOGFONTW结构体指针

[func (a *App) Malloc(size int) int {]
ff=分配内存
size=大小

[func (a *App) Free(p int) int {]
ff=释放内存
p=内存首地址

[func (a *App) Alert(pTitle string, pText string) int {]
ff=弹框
pText=提示内容
pTitle=提示框标题

[func (a *App) Sys_ShellExecute(hwnd uintptr, lpOperation string, lpFile string, lpParameters string, lpDirectory string, nShowCmd xcc.SW_) uintptr {]
ff=对指定文件执行操作

[func (a *App) LoadLibrary(lpFileName string) uintptr {]
ff=载入动态库
lpFileName=文件名

[func (a *App) GetProcAddress(hModule uintptr, lpProcName string) uintptr {]
ff=取动态库中函数地址
lpProcName=函数名
hModule=动态库模块句柄

[func (a *App) FreeLibrary(hModule uintptr) bool {]
ff=释放动态库
hModule=动态库模块句柄

[func (a *App) LoadDll(pDllFileName string) uintptr {]
ff=加载DLL
pDllFileName=DLL文件名

[func (a *App) PostQuitMessage(nExitCode int) int {]
ff=发送WM_QUIT消息退出消息循环
nExitCode=退出码

[func (a *App) LoadLayout(pFileName string, hParent int, hAttachWnd uintptr) int {]
ff=加载布局文件
hAttachWnd=附加窗口句柄
hParent=父对象句柄
pFileName=布局文件名

[func (a *App) LoadLayoutZip(pZipFileName string, pFileName string, pPassword string, hParent int, hAttachWnd uintptr) int {]
ff=加载布局文件ZIP
hAttachWnd=附加窗口句柄
hParent=父对象句柄
pPassword=zip密码
pFileName=布局文件名
pZipFileName=zip文件名

[func (a *App) LoadLayoutZipMem(data #左中括号##右中括号#byte, pFileName string, pPassword string, hParent int, hAttachWnd uintptr) int {]
ff=加载布局文件内存ZIP
hAttachWnd=附加窗口句柄
hParent=父对象句柄
pPassword=zip密码
pFileName=布局文件名
data=内存块指针

[func (a *App) LoadLayoutZipResEx(id int32, pFileName string, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd, hModule uintptr) int {]
ff=加载布局文件资源ZIPEX
hModule=模块句柄
hAttachWnd=附加窗口句柄
hParentWnd=父窗口句柄HWND
hParent=父对象句柄
pPrefixName=名称前缀
pPassword=zip密码
pFileName=布局文件名
id=RC资源ID

[func (a *App) LoadResourceZipRes(id int, pFileName string, pPassword string, hModule uintptr) bool {]
ff=加载资源文件资源ZIP
hModule=模块句柄
pPassword=zip压缩包密码
pFileName=资源文件名
id=RC资源ID

[func (a *App) LoadStyleZipRes(id int, pFileName string, pPassword string, hModule uintptr) bool {]
ff=加载样式文件从资源ZIP
hModule=模块句柄
pPassword=密码
pFileName=文件名
id=RC资源ID

[func (a *App) LoadLayoutFromStringW(pStringXML string, hParent int, hAttachWnd uintptr) int {]
ff=加载布局文件从字符串W
hAttachWnd=附加窗口句柄
hParent=父对象
pStringXML=字符串

[func (a *App) LoadLayoutEx(pFileName, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) int {]
ff=加载布局文件Ex
hAttachWnd=附加窗口句柄
hParentWnd=父窗口句柄HWND
hParent=父对象句柄
pPrefixName=名称
pFileName=布局文件名

[func (a *App) LoadLayoutZipEx(pZipFileName string, pFileName string, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) int {]
ff=加载布局文件ZIPEx
hAttachWnd=附加窗口句柄
hParentWnd=父窗口句柄HWND
hParent=父对象句柄
pPrefixName=名称
pPassword=zip密码
pFileName=布局文件名
pZipFileName=zip文件名

[func (a *App) LoadLayoutZipMemEx(data #左中括号##右中括号#byte, pFileName string, pPassword, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) int {]
ff=加载布局文件内存ZIPEx
hAttachWnd=附加窗口句柄
hParentWnd=父窗口句柄HWND
hParent=父对象句柄
pPrefixName=名称
pPassword=zip密码
pFileName=布局文件名
data=布局文件数据

[func (a *App) LoadLayoutFromStringWEx(pStringXML, pPrefixName string, hParent int, hParentWnd, hAttachWnd uintptr) int {]
ff=加载布局文件从字符串WEx
hAttachWnd=附加窗口句柄
hParentWnd=父窗口句柄HWND
hParent=父对象句柄
pPrefixName=名称
pStringXML=字符串

[func (a *App) LoadStyle(pFileName string) bool {]
ff=加载样式文件
pFileName=样式文件名称

[func (a *App) LoadStyleZip(pZipFile string, pFileName string, pPassword string) bool {]
ff=加载样式文件ZIP
pPassword=密码
pFileName=文件名
pZipFile=ZIP文件名

[func (a *App) LoadStyleZipMem(data #左中括号##右中括号#byte, pFileName string, pPassword string) bool {]
ff=加载样式文件从内存ZIP
pPassword=密码
pFileName=文件名
data=样式文件数据

[func (a *App) LoadResource(pFileName string) bool {]
ff=加载资源文件
pFileName=资源文件名

[func (a *App) LoadResourceZip(pZipFileName string, pFileName string, pPassword string) bool {]
ff=加载资源文件ZIP
pPassword=zip压缩包密码
pFileName=资源文件名
pZipFileName=zip文件名

[func (a *App) LoadResourceZipMem(data #左中括号##右中括号#byte, pFileName string, pPassword string) bool {]
ff=加载资源文件内存ZIP
pPassword=zip压缩包密码
pFileName=资源文件名
data=资源文件数据

[func (a *App) LoadResourceFromStringW(pStringXML string, pFileName string) bool {]
ff=加载资源文件从字符串W
pFileName=资源文件名
pStringXML=字符串

[func (a *App) WtoA(pValue string) uintptr {]
ff=W2A
pValue=参数

[func (a *App) AtoW(pValue uintptr) string {]
ff=A2W
pValue=参数

[func (a *App) Utf8toW(pUtf8 uintptr) string {]
ff=UTF8到文本W
pUtf8=参数

[func (a *App) Utf8toWEx(pUtf8 uintptr, length int) string {]
ff=UTF8到文本WEX
length=utf8字符串长度
pUtf8=utf8字符串指针

[func (a *App) Utf8toA(pUtf8 uintptr) uintptr {]
ff=UTF8到文本A
pUtf8=utf8字符串指针

[func (a *App) AtoUtf8(pValue uintptr) uintptr {]
ff=文本A到UTF8
pValue=参数

[func (a *App) WtoUtf8(pValue string) uintptr {]
ff=文本W到UTF8
pValue=字符串

[func (a *App) WtoUtf8Ex(pValue string, length int) uintptr {]
ff=文本W到UTF8EX
length=字符串长度
pValue=字符串

[func (a *App) UnicodeToAnsi(pIn string, inLen int, pOut uintptr, outLen int) int {]
ff=U2A
outLen=pOut缓冲区大小
pOut=接收转换后缓冲区指针
inLen=pIn字符数量
pIn=待转换Unicode文本

[func (a *App) AnsiToUnicode(pIn uintptr, inLen int, pOut *string, outLen int) int {]
ff=XA2U
outLen=pOut缓冲区大小
pOut=接收转换后缓冲区指针
inLen=pIn字符数量
pIn=待转换Ansi文本指针

[func (a *App) Print(level int, pInfo string) int {]
ff=打印调试信息
pInfo=信息
level=级别

[func (a *App) GetD2dFactory() int {]
ff=取D2D工厂

[func (a *App) GetDWriteFactory() int {]
ff=取DWrite工厂

[func (a *App) GetWicFactory() int {]
ff=取WIC工厂

[func (a *App) SetD2dTextRenderingMode(mode xcc.XC_DWRITE_RENDERING_MODE_) int {]
ff=置D2D文本渲染模式
mode=渲染模式

[func (a *App) IsEnableD2D() bool {]
ff=是否启用了D2D

[func (a *App) LoadStyleFromStringW(pFileName string, pString string) bool {]
ff=加载样式文件并按字符串W
pString=字符串
pFileName=样式文件名

[func (a *App) ShowSvgFrame(bShow bool) int {]
ff=显示边界
bShow=是否显示

[func (a *App) NotifyMsg_Popup(position xcc.Position_Flag_, pTitle, pText string, hIcon int, skin xcc.NotifyMsg_Skin_) int {]
ff=通知消息_弹出_未实现
position=位置

[func (a *App) NotifyMsg_PopupEx(position xcc.Position_Flag_, pTitle, pText string, hIcon int, skin xcc.NotifyMsg_Skin_, bBtnClose, bAutoClose bool, nWidth, nHeight int) int {]
ff=通知消息_弹出EX_未实现
position=位置

[func (a *App) NotifyMsg_SetDuration(hWindow, duration int) int {]
ff=通知消息_置持续时间
duration=持续时间
hWindow=窗口句柄

[func (a *App) NotifyMsg_SetParentMargin(hWindow, left, top, right, bottom int) int {]
ff=通知消息_置父边距
bottom=底部间隔
right=右侧间隔
top=顶部间隔
left=左侧间隔
hWindow=窗口句柄

[func (a *App) NotifyMsg_SetCaptionHeight(hWindow, nHeight int) int {]
ff=通知消息_置标题高度
nHeight=高度
hWindow=窗口句柄

[func (a *App) NotifyMsg_SetWidth(hWindow, nWidth int) int {]
ff=通知消息_置宽度
nWidth=宽度
hWindow=窗口句柄

[func (a *App) NotifyMsg_SetSpace(hWindow, nSpace int) int {]
ff=通知消息_置间隔
nSpace=间隔大小
hWindow=窗口句柄

[func (a *App) NotifyMsg_SetBorderSize(hWindow, left, top, right, bottom int) int {]
ff=通知消息_置边大小
bottom=底边
right=右边
top=顶边
left=左边
hWindow=窗口句柄

[func (a *App) EnableAutoDPI(bEnabel bool) int {]
ff=启用自动DPI
bEnabel=是否启用

[func (a *App) SetWindowIcon(hImage int) int {]
ff=置窗口图标
hImage=图标句柄

[func (a *App) EnableDPI(bEnabel bool) bool {]
ff=启用DPI
bEnabel=是否启用
