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

[func XWnd_Create(x, y, cx, cy int32, pTitle string, hWndParent uintptr, XCStyle xcc.Window_Style_) int {]
ff=窗口_创建
XCStyle=GUI库窗口样式
hWndParent=父窗口
pTitle=窗口标题
cy=高度
cx=宽度
y=左上角y坐标
x=左上角x坐标

[func XWnd_CreateEx(dwExStyle int, dwStyle int, lpClassName string, x int, y int, cx int, cy int, pTitle string, hWndParent uintptr, XCStyle xcc.Window_Style_) int {]
ff=窗口_创建EX
XCStyle=GUI库窗口样式
hWndParent=父窗口
pTitle=窗口名
cy=高度
cx=宽度
y=左上角y坐标
x=左上角x坐标
lpClassName=类名
dwStyle=样式
dwExStyle=扩展样式

[func XWnd_ShowWindow(hWindow int, nCmdShow xcc.SW_) int {]
ff=窗口_显示方式
nCmdShow=显示方式
hWindow=窗口句柄

[func XWnd_SetTop(hWindow int) int {]
ff=窗口_置顶
hWindow=窗口句柄

[func XWnd_RegEventC(hWindow int, nEvent xcc.WM_, pFun interface{}) bool {]
ff=窗口_注册事件C
nEvent=事件类型
hWindow=窗口句柄

[func XWnd_RegEventC1(hWindow int, nEvent xcc.WM_, pFun interface{}) bool {]
ff=窗口_注册事件C1
nEvent=事件类型
hWindow=窗口句柄

[func XWnd_RemoveEventC(hWindow int, nEvent xcc.WM_, pFun interface{}) bool {]
ff=窗口_移除事件C
nEvent=事件类型
hWindow=窗口句柄

[func XWnd_RegEventCEx(hWindow int, nEvent xcc.WM_, pFun uintptr) bool {]
ff=窗口_注册事件CEx
nEvent=事件类型
hWindow=窗口句柄

[func XWnd_RegEventC1Ex(hWindow int, nEvent xcc.WM_, pFun uintptr) bool {]
ff=窗口_注册事件C1Ex
nEvent=事件类型
hWindow=窗口句柄

[func XWnd_RemoveEventCEx(hWindow int, nEvent xcc.WM_, pFun uintptr) bool {]
ff=窗口_移除事件CEx
nEvent=事件类型
hWindow=窗口句柄

[func XWnd_AddChild(hWindow int, hChild int) bool {]
ff=窗口_添加子对象
hChild=要添加的对象句柄
hWindow=窗口句柄

[func XWnd_InsertChild(hWindow int, hChild int, index int) bool {]
ff=窗口_插入子对象
index=插入位置索引
hChild=要插入的对象句柄
hWindow=窗口句柄

[func XWnd_GetHWND(hWindow int) uintptr {]
ff=窗口_取HWND
hWindow=窗口句柄

[func XWnd_Redraw(hWindow int, bImmediate bool) int {]
ff=窗口_重绘
bImmediate=是否立即重绘
hWindow=窗口资源句柄

[func XWnd_RedrawRect(hWindow int, pRect *RECT, bImmediate bool) int {]
ff=窗口_重绘指定区域
bImmediate=TRUE立即重绘
pRect=区域坐标
hWindow=窗口资源句柄

[func XWnd_SetRect(hWindow int, pRect *RECT) bool {]
ff=窗口_置坐标
pRect=坐标
hWindow=窗口句柄

[func XWnd_SetFocusEle(hWindow int, hFocusEle int) bool {]
ff=窗口_置焦点
hFocusEle=将获得焦点的元素
hWindow=窗口资源句柄

[func XWnd_GetFocusEle(hWindow int) int {]
ff=窗口_取焦点
hWindow=窗口资源句柄

[func XWnd_GetStayEle(hWindow int) int {]
ff=窗口_取鼠标停留元素
hWindow=窗口资源句柄

[func XWnd_DrawWindow(hWindow int, hDraw int) int {]
ff=窗口_绘制
hDraw=图形绘制句柄
hWindow=窗口资源句柄

[func XWnd_Center(hWindow int) int {]
ff=窗口_居中
hWindow=窗口资源句柄

[func XWnd_CenterEx(hWindow, width, height int) int {]
ff=窗口_居中EX
height=窗口高度
width=窗口宽度
hWindow=窗口资源句柄

[func XWnd_SetCursor(hWindow int, hCursor int) int {]
ff=窗口_置光标
hCursor=鼠标光标句柄
hWindow=窗口句柄

[func XWnd_GetCursor(hWindow int) int {]
ff=窗口_取光标
hWindow=窗口句柄

[func XWnd_EnableDragBorder(hWindow int, bEnable bool) int {]
ff=窗口_启用拖动边框
bEnable=是否启用
hWindow=窗口句柄

[func XWnd_EnableDragWindow(hWindow int, bEnable bool) int {]
ff=窗口_启用拖动窗口
bEnable=是否启用
hWindow=窗口句柄

[func XWnd_EnableDragCaption(hWindow int, bEnable bool) int {]
ff=窗口_启用拖动标题栏
bEnable=是否启用
hWindow=窗口句柄

[func XWnd_EnableDrawBk(hWindow int, bEnable bool) int {]
ff=窗口_启用绘制背景
bEnable=是否启用
hWindow=窗口句柄

[func XWnd_EnableAutoFocus(hWindow int, bEnable bool) int {]
ff=窗口_启用自动焦点
bEnable=是否启用
hWindow=窗口句柄

[func XWnd_EnableMaxWindow(hWindow int, bEnable bool) int {]
ff=窗口_启用允许最大化
bEnable=是否启用
hWindow=窗口句柄

[func XWnd_EnableLimitWindowSize(hWindow int, bEnable bool) int {]
ff=窗口_启用限制窗口大小
bEnable=是否启用
hWindow=窗口句柄

[func XWnd_EnableLayout(hWindow int, bEnable bool) int {]
ff=窗口_启用布局
bEnable=是否启用
hWindow=窗口句柄

[func XWnd_EnableLayoutOverlayBorder(hWindow int, bEnable bool) int {]
ff=窗口_启用布局覆盖边框
bEnable=是否启用
hWindow=窗口句柄

[func XWnd_ShowLayoutFrame(hWindow int, bEnable bool) int {]
ff=窗口_显示布局边界
bEnable=是否启用
hWindow=窗口句柄

[func XWnd_IsEnableLayout(hWindow int) bool {]
ff=窗口_判断启用布局
hWindow=窗口句柄

[func XWnd_IsMaxWindow(hWindow int) bool {]
ff=窗口_是否最大化
hWindow=窗口句柄

[func XWnd_SetCaptureEle(hWindow, hEle int) int {]
ff=窗口_置鼠标捕获元素
hEle=元素句柄
hWindow=窗口句柄

[func XWnd_GetCaptureEle(hWindow int) int {]
ff=窗口_取鼠标捕获元素
hWindow=窗口句柄

[func XWnd_GetDrawRect(hWindow int, pRcPaint *RECT) int {]
ff=窗口_取绘制矩形
pRcPaint=重绘区域坐标
hWindow=窗口句柄

[func XWnd_SetCursorSys(hWindow, hCursor int) int {]
ff=窗口_置系统光标
hCursor=光标句柄
hWindow=窗口句柄

[func XWnd_SetFont(hWindow, hFontx int) int {]
ff=窗口_置字体
hFontx=炫彩字体句柄
hWindow=窗口句柄

[func XWnd_SetTextColor(hWindow, color int) int {]
ff=窗口_置文本颜色
color=ABGR颜色值
hWindow=窗口句柄

[func XWnd_GetTextColor(hWindow int) int {]
ff=窗口_取文本颜色
hWindow=窗口句柄

[func XWnd_GetTextColorEx(hWindow int) int {]
ff=窗口_取文本颜色EX
hWindow=窗口句柄

[func XWnd_SetID(hWindow, nID int) int {]
ff=窗口_置ID
nID=ID值
hWindow=窗口句柄

[func XWnd_GetID(hWindow int) int {]
ff=窗口_取ID
hWindow=窗口句柄

[func XWnd_SetName(hWindow int, pName string) int {]
ff=窗口_置名称
pName=名称
hWindow=窗口句柄

[func XWnd_GetName(hWindow int) string {]
ff=窗口_取名称
hWindow=窗口句柄

[func XWnd_SetBorderSize(hWindow, left, top, right, bottom int) int {]
ff=窗口_置边大小
bottom=底部
right=右边
top=上边
left=左边
hWindow=窗口句柄

[func XWnd_GetBorderSize(hWindow int, pBorder *RECT) int {]
ff=窗口_取边大小
hWindow=窗口句柄

[func XWnd_SetPadding(hWindow, left, top, right, bottom int) int {]
ff=窗口_置布局内填充大小
bottom=底部
right=右边
top=上边
left=左边
hWindow=窗口句柄

[func XWnd_SetDragBorderSize(hWindow, left, top, right, bottom int) int {]
ff=窗口_置拖动边框大小
bottom=底部
right=右边
top=上边
left=左边
hWindow=窗口句柄

[func XWnd_GetDragBorderSize(hWindow int, pBorder *RECT) int {]
ff=窗口_取拖动边框大小
hWindow=窗口句柄

[func XWnd_SetMinimumSize(hWindow, width, height int) int {]
ff=窗口_置最小大小
height=最小高度
width=最小宽度
hWindow=窗口句柄

[func XWnd_HitChildEle(hWindow int, pPt *POINT) int {]
ff=窗口_测试点击子元素
pPt=左边点
hWindow=窗口句柄

[func XWnd_GetChildCount(hWindow int) int {]
ff=窗口_取子对象数量
hWindow=窗口句柄

[func XWnd_GetChildByIndex(hWindow, index int) int {]
ff=窗口_取子对象从索引
index=元素索引
hWindow=窗口句柄

[func XWnd_GetChildByID(hWindow, nID int) int {]
ff=窗口_取子对象从ID
nID=元素ID
hWindow=窗口句柄

[func XWnd_GetChild(hWindow, nID int) int {]
ff=窗口_取子对象
nID=对象ID
hWindow=窗口句柄

[func XWnd_CloseWindow(hWindow int) int {]
ff=窗口_关闭
hWindow=窗口句柄

[func XWnd_AdjustLayout(hWindow int) int {]
ff=窗口_调整布局
hWindow=窗口句柄

[func XWnd_AdjustLayoutEx(hWindow int, nFlags xcc.AdjustLayout_) int {]
ff=窗口_调整布局EX
nFlags=调整布局标识位
hWindow=窗口句柄

[func XWnd_CreateCaret(hWindow, hEle, x, y, width, height int) int {]
ff=窗口_创建插入符
height=高度
width=宽度
y=y坐标
x=x坐标
hEle=元素句柄
hWindow=窗口句柄

[func XWnd_SetCaretPos(hWindow, x, y, width, height int, bUpdate bool) int {]
ff=窗口_置插入符位置
bUpdate=是否立即更新UI
height=高度
width=宽度
y=y坐标
x=x坐标
hWindow=窗口句柄

[func XWnd_SetCaretColor(hWindow, color int) int {]
ff=窗口_置插入符颜色
color=颜色值
hWindow=窗口句柄

[func XWnd_ShowCaret(hWindow int, bShow bool) int {]
ff=窗口_显示插入符
bShow=是否显示
hWindow=窗口句柄

[func XWnd_DestroyCaret(hWindow int) int {]
ff=窗口_销毁插入符
hWindow=窗口句柄

[func XWnd_GetCaretHELE(hWindow int) int {]
ff=窗口_取插入符元素
hWindow=窗口句柄

[func XWnd_GetClientRect(hWindow int, pRect *RECT) int {]
ff=窗口_取客户区坐标
pRect=坐标
hWindow=窗口句柄

[func XWnd_GetBodyRect(hWindow int, pRect *RECT) int {]
ff=窗口_取Body坐标
pRect=坐标
hWindow=窗口句柄

[func XWnd_GetLayoutRect(hWindow int, pRect *RECT) int {]
ff=窗口_取布局坐标
pRect=接收返回坐标
hWindow=窗口句柄

[func XWnd_SetPosition(hWindow int, x, y int32) {]
ff=窗口_移动窗口
y=Y坐标
x=坐标
hWindow=窗口句柄

[func XWnd_GetRect(hWindow int, pRect *RECT) int {]
ff=窗口_取坐标
pRect=坐标
hWindow=窗口句柄

[func XWnd_MaxWindow(hWindow int, bMaximize bool) int {]
ff=窗口_最大化
bMaximize=是否最大化
hWindow=窗口句柄

[func XWnd_SetTimer(hWindow, nIDEvent, uElapse int) int {]
ff=窗口_置定时器
uElapse=间隔值
nIDEvent=定时器ID
hWindow=窗口句柄

[func XWnd_KillTimer(hWindow, nIDEvent int) int {]
ff=窗口_关闭定时器
nIDEvent=定时器ID
hWindow=窗口句柄

[func XWnd_SetXCTimer(hWindow, nIDEvent, uElapse int) int {]
ff=窗口_置炫彩定时器
uElapse=间隔值
nIDEvent=定时器ID
hWindow=窗口句柄

[func XWnd_KillXCTimer(hWindow, nIDEvent int) int {]
ff=窗口_关闭炫彩定时器
nIDEvent=定时器ID
hWindow=窗口句柄

[func XWnd_GetBkManager(hWindow int) int {]
ff=窗口_取背景管理器
hWindow=窗口句柄

[func XWnd_GetBkManagerEx(hWindow int) int {]
ff=窗口_取背景管理器EX
hWindow=窗口句柄

[func XWnd_SetBkMagager(hWindow, hBkInfoM int) int {]
ff=窗口_置背景管理器
hBkInfoM=背景管理器
hWindow=窗口句柄

[func XWnd_SetTransparentType(hWindow int, nType xcc.Window_Transparent_) int {]
ff=窗口_置透明类型
nType=透明类型
hWindow=窗口句柄

[func XWnd_SetTransparentAlpha(hWindow int, alpha byte) int {]
ff=窗口_置透明度
alpha=透明度
hWindow=窗口句柄

[func XWnd_SetTransparentColor(hWindow, color int) int {]
ff=窗口_置透明色
color=透明色
hWindow=窗口句柄

[func XWnd_SetShadowInfo(hWindow, nSize int, nDepth byte, nAngeleSize int, bRightAngle bool, color int) int {]
ff=窗口_置阴影信息
color=阴影颜色
bRightAngle=是否强制直角
nAngeleSize=圆角阴影内收大小
nDepth=阴影深度
nSize=阴影大小
hWindow=窗口句柄

[func XWnd_GetShadowInfo(hWindow int, pnSize, pnDepth, pnAngeleSize *int32, pbRightAngle *bool, pColor *int) int {]
ff=窗口_取阴影信息
pColor=阴影颜色
pbRightAngle=是否强制直角
pnAngeleSize=圆角阴影内收大小
pnDepth=阴影深度
pnSize=阴影大小
hWindow=窗口句柄

[func XWnd_GetTransparentType(hWindow int) xcc.Window_Transparent_ {]
ff=窗口_取透明类型
hWindow=窗口句柄

[func XWnd_Attach(hWnd uintptr, XCStyle xcc.Window_Style_) int {]
ff=窗口_附加窗口
XCStyle=窗口样式
hWnd=外部窗口句柄

[func XWnd_EnableDragFiles(hWindow int, bEnable bool) bool {]
ff=窗口_启用拖放文件
bEnable=是否启用
hWindow=窗口句柄

[func XWnd_Show(hWindow int, bShow bool) int {]
ff=窗口_显示
bShow=是否显示
hWindow=窗口句柄

[func XWnd_GetCaretInfo(hWindow int, pX, pY, pWidth, pHeight *int32) int {]
ff=窗口_取插入符信息
pHeight=接收返回高度
pWidth=接收返回宽度
pY=接收返回y坐标
pX=接收返回x坐标
hWindow=窗口句柄

[func XWnd_SetIcon(hWindow, hImage int) int {]
ff=窗口_置图标
hImage=图片句柄
hWindow=窗口句柄

[func XWnd_SetTitle(hWindow int, pTitle string) int {]
ff=窗口_置标题
pTitle=标题文本
hWindow=窗口句柄

[func XWnd_SetTitleColor(hWindow, color int) int {]
ff=窗口_置标题颜色
color=ABGR颜色
hWindow=窗口句柄

[func XWnd_GetButton(hWindow int, nFlag xcc.Window_Style_) int {]
ff=窗口_取控制按钮
nFlag=标志
hWindow=窗口句柄

[func XWnd_GetIcon(hWindow int) int {]
ff=窗口_取图标
hWindow=窗口句柄

[func XWnd_GetTitle(hWindow int) string {]
ff=窗口_取标题
hWindow=窗口句柄

[func XWnd_GetTitleColor(hWindow int) int {]
ff=窗口_取标题颜色
hWindow=窗口句柄

[func XWnd_AddBkBorder(hWindow int, nState xcc.Window_State_Flag_, color int, width int) int {]
ff=窗口_添加背景边框
nState=组合状态
hWindow=窗口句柄

[func XWnd_AddBkFill(hWindow int, nState xcc.Window_State_Flag_, color int) int {]
ff=窗口_添加背景填充
nState=组合状态
hWindow=窗口句柄

[func XWnd_AddBkImage(hWindow int, nState xcc.Window_State_Flag_, hImage int) int {]
ff=窗口_添加背景图片
nState=组合状态
hWindow=窗口句柄

[func XWnd_GetBkInfoCount(hWindow int) int {]
ff=窗口_取背景对象数量
hWindow=窗口句柄

[func XWnd_ClearBkInfo(hWindow int) int {]
ff=窗口_清空背景对象
hWindow=窗口句柄

[func XWnd_SetBkInfo(hWindow int, pText string) int {]
ff=窗口_置背景
pText=背景内容字符串
hWindow=窗口句柄

[func XWnd_IsDragCaption(hWindow int) bool {]
ff=窗口_是否可拖动标题栏
hWindow=窗口句柄

[func XWnd_IsDragWindow(hWindow int) bool {]
ff=窗口_是否可拖动窗口
hWindow=窗口句柄

[func XWnd_IsDragBorder(hWindow int) bool {]
ff=窗口_是否可拖动边框
hWindow=窗口句柄

[func XWnd_SetCaptionMargin(hWindow int, left int, top int, right int, bottom int) int {]
ff=窗口_置标题外间距
bottom=下边间距
right=右边间距
top=上边间距
left=左边间距
hWindow=窗口句柄

[func XWnd_SetWindowPos(hWindow int, hWndInsertAfter xcc.HWND_, x, y, cx, cy int32, uFlags xcc.SWP_) int {]
ff=窗口_置窗口位置
hWndInsertAfter=置顶方式
hWindow=窗口句柄

[func XWnd_GetDPI(hWindow int) int {]
ff=窗口_取DPI
hWindow=窗口句柄

[func XWnd_RectToDPI(hWindow int, pRect *RECT) int {]
ff=窗口_坐标转换DPI
pRect=接收返回坐标
hWindow=窗口句柄

[func XWnd_PointToDPI(hWindow int, pPt *POINT) int {]
ff=窗口_坐标点转换DPI
pPt=接收返回坐标点
hWindow=窗口句柄

[func XWnd_GetCursorPos(hWindow int, pPt *POINT) bool {]
ff=窗口_取光标位置
pPt=接收返回坐标点
hWindow=窗口句柄

[func XWnd_ClientToScreen(hWindow int, pPt *POINT) bool {]
ff=窗口_客户区坐标点到屏幕坐标点
pPt=接收返回坐标点
hWindow=窗口句柄

[func XWnd_ScreenToClient(hWindow int, pPt *POINT) bool {]
ff=窗口_屏幕坐标点到客户区坐标点
pPt=接收返回坐标点
hWindow=窗口句柄

[func XWnd_SetDPI(hWindow int, nDPI int) int {]
ff=窗口_置DPI
nDPI=DPI值
hWindow=窗口句柄
