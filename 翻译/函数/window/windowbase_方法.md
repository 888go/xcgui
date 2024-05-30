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

[func (w *windowBase) MessageBox(pTitle, pText string, nFlags xcc.MessageBox_Flag_, XCStyle xcc.Window_Style_) xcc.MessageBox_Flag_ {]
ff=消息框
nFlags=标识
pText=内容文本
pTitle=标题

[func (w *windowBase) Msg_Create(pTitle, pText string, nFlags xcc.MessageBox_Flag_, XCStyle xcc.Window_Style_) *ModalWindow {]
ff=消息框创建
nFlags=标识
pText=内容文本

[func (w *windowBase) Msg_CreateEx(dwExStyle, dwStyle int, lpClassName, pTitle, pText string, nFlags xcc.MessageBox_Flag_, XCStyle xcc.Window_Style_) *ModalWindow {]
ff=消息框创建EX
nFlags=标识
pText=内容文本
pTitle=标题
lpClassName=窗口类名
dwStyle=窗口样式

[func (w *windowBase) SendMessage(msg uint32, wParam, lParam uint) uint {]
ff=发送窗口消息
lParam=参数2
wParam=参数1
msg=消息值

[func (w *windowBase) PostMessage(msg uint32, wParam int32, lParam int32) bool {]
ff=投递窗口消息
lParam=参数2
wParam=参数1
msg=消息值

[func (w *windowBase) IsHWINDOW() bool {]
ff=判断窗口

[func (w *windowBase) GetObjectByID(nID int) int {]
ff=取对象并按ID
nID=ID值

[func (w *windowBase) GetObjectByIDName(pName string) int {]
ff=取对象并按ID名称
pName=ID名称

[func (w *windowBase) ShowWindow(nCmdShow xcc.SW_) int {]
ff=显示方式
nCmdShow=显示方式

[func (w *windowBase) SetTop() int {]
ff=置顶

[func (w *windowBase) RegEventC(nEvent xcc.WM_, pFun interface{}) bool {]
ff=注册事件C
nEvent=事件类型

[func (w *windowBase) RegEventC1(nEvent xcc.WM_, pFun interface{}) bool {]
ff=注册事件C1
nEvent=事件类型

[func (w *windowBase) RemoveEventC(nEvent xcc.WM_, pFun interface{}) bool {]
ff=移除事件C
nEvent=事件类型

[func (w *windowBase) RemoveEventCEx(nEvent xcc.WM_, pFun uintptr) bool {]
ff=移除事件CEx
nEvent=事件类型

[func (w *windowBase) AddChild(hChild int) bool {]
ff=添加子对象
hChild=要添加的对象句柄

[func (w *windowBase) InsertChild(hChild int, index int) bool {]
ff=插入子对象
index=插入位置索引
hChild=要插入的对象句柄

[func (w *windowBase) GetHWND() uintptr {]
ff=取HWND

[func (w *windowBase) Redraw(bImmediate bool) int {]
ff=重绘
bImmediate=是否立即重绘

[func (w *windowBase) RedrawRect(pRect *xc.RECT, bImmediate bool) int {]
ff=重绘指定区域
pRect=需要重绘的区域坐标

[func (w *windowBase) SetRect(pRect *xc.RECT) bool {]
ff=置坐标
pRect=坐标

[func (w *windowBase) SetFocusEle(hFocusEle int) bool {]
ff=置焦点
hFocusEle=元素

[func (w *windowBase) GetFocusEle() int {]
ff=取焦点

[func (w *windowBase) GetStayEle() int {]
ff=取鼠标停留元素

[func (w *windowBase) DrawWindow(hDraw int) int {]
ff=绘制
hDraw=图形绘制句柄

[func (w *windowBase) Center() int {]
ff=居中

[func (w *windowBase) CenterEx(width, height int) int {]
ff=居中EX
height=窗口高度
width=窗口宽度

[func (w *windowBase) SetCursor(hCursor int) int {]
ff=置光标
hCursor=鼠标光标句柄

[func (w *windowBase) GetCursor() int {]
ff=取光标

[func (w *windowBase) EnableDragBorder(bEnable bool) int {]
ff=启用拖动边框
bEnable=是否启用

[func (w *windowBase) EnableDragWindow(bEnable bool) int {]
ff=启用拖动窗口
bEnable=是否启用

[func (w *windowBase) EnableDragCaption(bEnable bool) int {]
ff=启用拖动标题栏
bEnable=是否启用

[func (w *windowBase) EnableDrawBk(bEnable bool) int {]
ff=启用绘制背景
bEnable=是否启用

[func (w *windowBase) EnableAutoFocus(bEnable bool) int {]
ff=启用自动焦点
bEnable=是否启用

[func (w *windowBase) EnableMaxWindow(bEnable bool) int {]
ff=启用允许最大化
bEnable=是否启用

[func (w *windowBase) EnableLimitWindowSize(bEnable bool) int {]
ff=启用限制窗口大小
bEnable=是否启用

[func (w *windowBase) EnableLayout(bEnable bool) int {]
ff=启用布局
bEnable=是否启用

[func (w *windowBase) EnableLayoutOverlayBorder(bEnable bool) int {]
ff=启用布局覆盖边框
bEnable=是否启用

[func (w *windowBase) ShowLayoutFrame(bEnable bool) int {]
ff=显示布局边界
bEnable=是否启用

[func (w *windowBase) IsEnableLayout() bool {]
ff=判断启用布局

[func (w *windowBase) IsMaxWindow() bool {]
ff=是否最大化

[func (w *windowBase) SetCaptureEle(hEle int) int {]
ff=置鼠标捕获元素
hEle=元素句柄

[func (w *windowBase) GetCaptureEle() int {]
ff=取鼠标捕获元素

[func (w *windowBase) GetDrawRect(pRcPaint *xc.RECT) int {]
ff=取绘制矩形
pRcPaint=重绘区域坐标

[func (w *windowBase) SetCursorSys(hCursor int) int {]
ff=置系统光标
hCursor=光标句柄

[func (w *windowBase) SetFont(hFontx int) int {]
ff=置字体
hFontx=炫彩字体句柄

[func (w *windowBase) SetTextColor(color int) int {]
ff=置文本颜色
color=ABGR颜色值

[func (w *windowBase) GetTextColor() int {]
ff=取文本颜色

[func (w *windowBase) GetTextColorEx() int {]
ff=取文本颜色EX

[func (w *windowBase) SetID(nID int) int {]
ff=置ID
nID=ID值

[func (w *windowBase) GetID() int {]
ff=取ID

[func (w *windowBase) SetName(pName string) int {]
ff=置名称
pName=名称

[func (w *windowBase) GetName() string {]
ff=取名称

[func (w *windowBase) SetBorderSize(left, top, right, bottom int) int {]
ff=置边大小
bottom=底部
right=右边
top=上边
left=左边

[func (w *windowBase) GetBorderSize(pBorder *xc.RECT) int {]
ff=取边大小

[func (w *windowBase) SetPadding(left, top, right, bottom int) int {]
ff=置布局内填充大小
bottom=下边
right=右边
top=上边
left=左边

[func (w *windowBase) SetDragBorderSize(left, top, right, bottom int) int {]
ff=置拖动边框大小
bottom=底边
right=右边
top=上边
left=左边

[func (w *windowBase) GetDragBorderSize(pBorder *xc.RECT) int {]
ff=取拖动边框大小

[func (w *windowBase) SetMinimumSize(width, height int) int {]
ff=置最小大小
height=最小高度
width=最小宽度

[func (w *windowBase) HitChildEle(pPt *xc.POINT) int {]
ff=测试点击子元素
pPt=左边点

[func (w *windowBase) GetChildCount() int {]
ff=取子对象数量

[func (w *windowBase) GetChildByIndex(index int) int {]
ff=取子对象并按索引
index=元素索引

[func (w *windowBase) GetChildByID(nID int) int {]
ff=取子对象并按ID
nID=元素ID

[func (w *windowBase) GetChild(nID int) int {]
ff=取子对象
nID=对象ID

[func (w *windowBase) CloseWindow() int {]
ff=关闭

[func (w *windowBase) AdjustLayout() int {]
ff=调整布局

[func (w *windowBase) AdjustLayoutEx(nFlags xcc.AdjustLayout_) int {]
ff=调整布局EX
nFlags=调整布局标识位

[func (w *windowBase) CreateCaret(hEle, x, y, width, height int) int {]
ff=创建插入符
height=高度
width=宽度
y=y坐标
x=x坐标
hEle=元素句柄

[func (w *windowBase) SetCaretPos(x, y, width, height int, bUpdate bool) int {]
ff=置插入符位置
bUpdate=是否立即更新UI
height=高度
width=宽度
y=y坐标
x=x坐标

[func (w *windowBase) SetCaretColor(color int) int {]
ff=置插入符颜色
color=颜色值

[func (w *windowBase) ShowCaret(bShow bool) int {]
ff=显示插入符
bShow=是否显示

[func (w *windowBase) DestroyCaret() int {]
ff=销毁插入符

[func (w *windowBase) GetCaretHELE() int {]
ff=取插入符元素

[func (w *windowBase) GetClientRect(pRect *xc.RECT) int {]
ff=取客户区坐标
pRect=坐标

[func (w *windowBase) GetBodyRect(pRect *xc.RECT) int {]
ff=取Body坐标
pRect=坐标

[func (w *windowBase) GetLayoutRect(pRect *xc.RECT) int {]
ff=取布局坐标
pRect=接收返回坐标

[func (w *windowBase) SetPosition(x, y int32) {]
ff=移动窗口
y=Y坐标
x=坐标

[func (w *windowBase) GetRect(pRect *xc.RECT) int {]
ff=取坐标
pRect=坐标

[func (w *windowBase) MaxWindow(bMaximize bool) int {]
ff=最大化
bMaximize=是否最大化

[func (w *windowBase) SetTimer(nIDEvent, uElapse int) int {]
ff=置定时器
uElapse=间隔值
nIDEvent=定时器ID

[func (w *windowBase) KillTimer(nIDEvent int) int {]
ff=关闭定时器
nIDEvent=定时器ID

[func (w *windowBase) SetXCTimer(nIDEvent, uElapse int) int {]
ff=置炫彩定时器
uElapse=间隔值
nIDEvent=定时器ID

[func (w *windowBase) KillXCTimer(nIDEvent int) int {]
ff=关闭炫彩定时器
nIDEvent=定时器ID

[func (w *windowBase) GetBkManager() int {]
ff=取背景管理器

[func (w *windowBase) GetBkManagerEx() int {]
ff=取背景管理器EX

[func (w *windowBase) SetBkMagager(hBkInfoM int) int {]
ff=置背景管理器
hBkInfoM=背景管理器

[func (w *windowBase) SetTransparentType(nType xcc.Window_Transparent_) int {]
ff=置透明类型
nType=窗口透明类型

[func (w *windowBase) SetTransparentAlpha(alpha byte) int {]
ff=置透明度
alpha=窗口透明度

[func (w *windowBase) SetTransparentColor(color int) int {]
ff=置透明色
color=窗口透明色

[func (w *windowBase) SetShadowInfo(nSize int, nDepth byte, nAngeleSize int, bRightAngle bool, color int) int {]
ff=置阴影信息
color=阴影颜色
bRightAngle=是否强制直角
nAngeleSize=圆角阴影内收大小
nDepth=阴影深度
nSize=阴影大小

[func (w *windowBase) GetShadowInfo(pnSize, pnDepth, pnAngeleSize *int32, pbRightAngle *bool, pColor *int) int {]
ff=取阴影信息
pColor=阴影颜色
pbRightAngle=是否强制直角
pnAngeleSize=圆角阴影内收大小
pnDepth=阴影深度
pnSize=阴影大小

[func (w *windowBase) GetTransparentType() xcc.Window_Transparent_ {]
ff=取透明类型

[func (w *windowBase) EnableDragFiles(bEnable bool) bool {]
ff=启用拖放文件
bEnable=是否启用

[func (w *windowBase) Show(bShow bool) int {]
ff=显示
bShow=是否显示

[func (w *windowBase) GetCaretInfo(pX, pY, pWidth, pHeight *int32) int {]
ff=取插入符信息
pHeight=接收返回高度
pWidth=接收返回宽度
pY=接收返回y坐标
pX=接收返回x坐标

[func (w *windowBase) SetIcon(hImage int) int {]
ff=置图标
hImage=图标句柄

[func (w *windowBase) SetTitle(pTitle string) int {]
ff=置标题
pTitle=标题文本

[func (w *windowBase) SetTitleColor(color int) int {]
ff=置标题颜色
color=ABGR颜色

[func (w *windowBase) GetButton(nFlag xcc.Window_Style_) int {]
ff=取控制按钮
nFlag=xcc

[func (w *windowBase) GetIcon() int {]
ff=取图标

[func (w *windowBase) GetTitle() string {]
ff=取标题

[func (w *windowBase) GetTitleColor() int {]
ff=取标题颜色

[func (w *windowBase) AddBkBorder(nState xcc.Window_State_Flag_, color int, width int) int {]
ff=添加背景边框
nState=组合状态

[func (w *windowBase) AddBkFill(nState xcc.Window_State_Flag_, color int) int {]
ff=添加背景填充
nState=组合状态

[func (w *windowBase) AddBkImage(nState xcc.Window_State_Flag_, hImage int) int {]
ff=添加背景图片
nState=组合状态

[func (w *windowBase) GetBkInfoCount() int {]
ff=取背景对象数量

[func (w *windowBase) ClearBkInfo() int {]
ff=清空背景对象

[func (w *windowBase) NotifyMsg_WindowPopup(position xcc.Position_Flag_, pTitle, pText string, hIcon int, skin xcc.NotifyMsg_Skin_) int {]
ff=通知消息_窗口中弹出
position=位置

[func (w *windowBase) NotifyMsg_WindowPopupEx(position xcc.Position_Flag_, pTitle, pText string, hIcon int, skin xcc.NotifyMsg_Skin_, bBtnClose, bAutoClose bool, nWidth, nHeight int) int {]
ff=通知消息_窗口中弹出EX
position=位置

[func (w *windowBase) NotifyMsg_SetDuration(duration int) int {]
ff=通知消息_置持续时间
duration=持续时间

[func (w *windowBase) NotifyMsg_SetParentMargin(left, top, right, bottom int) int {]
ff=通知消息_置父边距
bottom=底部间隔
right=右侧间隔
top=顶部间隔
left=左侧间隔

[func (w *windowBase) NotifyMsg_SetCaptionHeight(nHeight int) int {]
ff=通知消息_置标题高度
nHeight=高度

[func (w *windowBase) NotifyMsg_SetWidth(nWidth int) int {]
ff=通知消息_置宽度
nWidth=宽度

[func (w *windowBase) NotifyMsg_SetSpace(nSpace int) int {]
ff=通知消息_置间隔
nSpace=间隔大小

[func (w *windowBase) NotifyMsg_SetBorderSize(left, top, right, bottom int) int {]
ff=通知消息_置边大小
bottom=底边
right=右边
top=顶边
left=左边

[func (w *windowBase) SetBkInfo(pText string) int {]
ff=置背景文本
pText=背景内容字符串

[func (w *windowBase) IsDragCaption() bool {]
ff=是否可拖动标题栏

[func (w *windowBase) IsDragWindow() bool {]
ff=是否可拖动窗口

[func (w *windowBase) IsDragBorder() bool {]
ff=是否可拖动边框

[func (w *windowBase) SetCaptionMargin(left int, top int, right int, bottom int) int {]
ff=置标题外间距
bottom=下
right=右
top=上
left=左

[func (w *windowBase) SetTopEx(b bool) bool {]
ff=XSetTopEx
b=是否置顶

[func (w *windowBase) SetWindowPos(hWndInsertAfter xcc.HWND_, x, y, cx, cy int32, uFlags xcc.SWP_) int {]
ff=置窗口位置
hWndInsertAfter=置顶方式

[func (w *windowBase) GetDPI() int {]
ff=取DPI

[func (w *windowBase) RectToDPI(pRect *xc.RECT) int {]
ff=坐标转换DPI
pRect=接收返回坐标

[func (w *windowBase) PointToDPI(pPt *xc.POINT) int {]
ff=坐标点转换DPI
pPt=接收返回坐标点

[func (w *windowBase) GetCursorPos(pPt *xc.POINT) bool {]
ff=取光标位置
pPt=接收返回坐标点

[func (w *windowBase) ClientToScreen(pPt *xc.POINT) bool {]
ff=客户区坐标点到屏幕坐标点
pPt=接收返回坐标点

[func (w *windowBase) ScreenToClient(pPt *xc.POINT) bool {]
ff=屏幕坐标点到客户区坐标点
pPt=接收返回坐标点

[func (w *windowBase) SetDPI(nDPI int) int {]
ff=置DPI
nDPI=DPI值

[func (w *windowBase) SetSize(width, height int32) bool {]
ff=置大小
height=高
width=宽

[func (w *windowBase) SetWidth(width int32) bool {]
ff=置宽度
width=宽

[func (w *windowBase) SetHeight(height int32) bool {]
ff=置高度
height=高

[func (w *windowBase) GetWidth() int32 {]
ff=取宽度

[func (w *windowBase) GetHeight() int32 {]
ff=取高度

[func (w *windowBase) GetLeft() int32 {]
ff=取左边

[func (w *windowBase) GetTop() int32 {]
ff=取顶边

[func (w *windowBase) GetRight() int32 {]
ff=取右边

[func (w *windowBase) GetBottom() int32 {]
ff=取底边

[func (w *windowBase) SetLeft(x int32) {]
ff=置左边
x=x坐标

[func (w *windowBase) SetTopEdge(y int32) {]
ff=置顶边
y=y坐标

[func (w *windowBase) EnableHorizon(bEnable bool) int {]
ff=启用水平
bEnable=是否启用

[func (w *windowBase) EnableAutoWrap(bEnable bool) int {]
ff=启用自动换行
bEnable=是否启用

[func (w *windowBase) EnableOverflowHide(bEnable bool) int {]
ff=启用溢出隐藏
bEnable=是否启用

[func (w *windowBase) SetAlignH(nAlign xcc.Layout_Align_) int {]
ff=置水平对齐
nAlign=对齐方式

[func (w *windowBase) SetAlignV(nAlign xcc.Layout_Align_) int {]
ff=置垂直对齐
nAlign=对齐方式

[func (w *windowBase) SetAlignBaseline(nAlign xcc.Layout_Align_Axis_) int {]
ff=置对齐基线
nAlign=对齐方式

[func (w *windowBase) SetSpace(nSpace int) int {]
ff=置间距
nSpace=项间距大小

[func (w *windowBase) SetSpaceRow(nSpace int) int {]
ff=置行距
nSpace=行间距大小

[func (w *windowBase) Event_TRAYICON(pFun XWM_TRAYICON) bool {]
ff=线程_托盘图标事件

[func (w *windowBase) Event_WINDPROC(pFun XWM_WINDPROC) bool {]
ff=线程_消息过程

[func (w *windowBase) Event_WINDPROC1(pFun XWM_WINDPROC1) bool {]
ff=线程_消息过程1

[func (w *windowBase) Event_XC_TIMER(pFun XWM_XC_TIMER) bool {]
ff=线程_炫彩定时器

[func (w *windowBase) Event_XC_TIMER1(pFun XWM_XC_TIMER1) bool {]
ff=线程_炫彩定时器1

[func (w *windowBase) Event_SETFOCUS_ELE(pFun XWM_SETFOCUS_ELE) bool {]
ff=事件_置焦点元素

[func (w *windowBase) Event_SETFOCUS_ELE1(pFun XWM_SETFOCUS_ELE1) bool {]
ff=事件_置焦点元素1

[func (w *windowBase) Event_FLOAT_PANE(pFun XWM_FLOAT_PANE) bool {]
ff=线程_浮动窗格

[func (w *windowBase) Event_FLOAT_PANE1(pFun XWM_FLOAT_PANE1) bool {]
ff=线程_浮动窗格1

[func (w *windowBase) Event_PAINT_END(pFun XWM_PAINT_END) bool {]
ff=线程_绘制完成消息

[func (w *windowBase) Event_PAINT_END1(pFun XWM_PAINT_END1) bool {]
ff=线程_绘制完成消息1

[func (w *windowBase) Event_PAINT_DISPLAY(pFun XWM_PAINT_DISPLAY) bool {]
ff=线程_绘制完成并且已经显示到屏幕

[func (w *windowBase) Event_PAINT_DISPLAY1(pFun XWM_PAINT_DISPLAY1) bool {]
ff=线程_绘制完成并且已经显示到屏幕1

[func (w *windowBase) Event_DOCK_POPUP(pFun XWM_DOCK_POPUP) bool {]
ff=线程_框架窗口码头弹出窗格

[func (w *windowBase) Event_DOCK_POPUP1(pFun XWM_DOCK_POPUP1) bool {]
ff=线程_框架窗口码头弹出窗格1

[func (w *windowBase) Event_BODYVIEW_RECT(pFun XWM_BODYVIEW_RECT) bool {]
ff=线程_框架窗口主视图坐标改变

[func (w *windowBase) Event_BODYVIEW_RECT1(pFun XWM_BODYVIEW_RECT) bool {]
ff=线程_框架窗口主视图坐标改变1

[func (w *windowBase) Event_FLOATWND_DRAG(pFun XWM_FLOATWND_DRAG) bool {]
ff=线程_浮动窗口拖动

[func (w *windowBase) Event_FLOATWND_DRAG1(pFun XWM_FLOATWND_DRAG1) bool {]
ff=线程_浮动窗口拖动1

[func (w *windowBase) Event_PAINT(pFun WM_PAINT) bool {]
ff=线程_绘制消息

[func (w *windowBase) Event_PAINT1(pFun WM_PAINT1) bool {]
ff=线程_绘制消息1

[func (w *windowBase) Event_CLOSE(pFun WM_CLOSE) bool {]
ff=线程_关闭消息

[func (w *windowBase) Event_CLOSE1(pFun WM_CLOSE1) bool {]
ff=线程_关闭消息1

[func (w *windowBase) Event_DESTROY(pFun WM_DESTROY) bool {]
ff=线程_销毁消息

[func (w *windowBase) Event_DESTROY1(pFun WM_DESTROY1) bool {]
ff=线程_销毁消息1

[func (w *windowBase) Event_NCDESTROY(pFun WM_NCDESTROY) bool {]
ff=线程_非客户区销毁消息

[func (w *windowBase) Event_NCDESTROY1(pFun WM_NCDESTROY1) bool {]
ff=线程_非客户区销毁消息1

[func (w *windowBase) Event_MOUSEMOVE(pFun WM_MOUSEMOVE) bool {]
ff=线程_鼠标移动消息

[func (w *windowBase) Event_MOUSEMOVE1(pFun WM_MOUSEMOVE1) bool {]
ff=线程_鼠标移动消息1

[func (w *windowBase) Event_LBUTTONDOWN(pFun WM_LBUTTONDOWN) bool {]
ff=线程_鼠标左键按下消息

[func (w *windowBase) Event_LBUTTONDOWN1(pFun WM_LBUTTONDOWN1) bool {]
ff=线程_鼠标左键按下消息1

[func (w *windowBase) Event_LBUTTONUP(pFun WM_LBUTTONUP) bool {]
ff=线程_鼠标左键弹起消息

[func (w *windowBase) Event_LBUTTONUP1(pFun WM_LBUTTONUP1) bool {]
ff=线程_鼠标左键弹起消息1

[func (w *windowBase) Event_RBUTTONDOWN(pFun WM_RBUTTONDOWN) bool {]
ff=线程_鼠标右键按下消息

[func (w *windowBase) Event_RBUTTONDOWN1(pFun WM_RBUTTONDOWN1) bool {]
ff=线程_鼠标右键按下消息1

[func (w *windowBase) Event_RBUTTONUP(pFun WM_RBUTTONUP) bool {]
ff=线程_鼠标右键弹起消息

[func (w *windowBase) Event_RBUTTONUP1(pFun WM_RBUTTONUP1) bool {]
ff=线程_鼠标右键弹起消息1

[func (w *windowBase) Event_LBUTTONDBLCLK(pFun WM_LBUTTONDBLCLK) bool {]
ff=线程_鼠标左键双击消息

[func (w *windowBase) Event_LBUTTONDBLCLK1(pFun WM_LBUTTONDBLCLK1) bool {]
ff=线程_鼠标左键双击消息1

[func (w *windowBase) Event_RBUTTONDBLCLK(pFun WM_RBUTTONDBLCLK) bool {]
ff=线程_鼠标右键双击消息

[func (w *windowBase) Event_RBUTTONDBLCLK1(pFun WM_RBUTTONDBLCLK1) bool {]
ff=线程_鼠标右键双击消息1

[func (w *windowBase) Event_MOUSEWHEEL(pFun WM_MOUSEWHEEL) bool {]
ff=线程_鼠标滚轮滚动消息

[func (w *windowBase) Event_MOUSEWHEEL1(pFun WM_MOUSEWHEEL1) bool {]
ff=线程_鼠标滚轮滚动消息1

[func (w *windowBase) Event_EXITSIZEMOVE(pFun WM_EXITSIZEMOVE) bool {]
ff=线程_退出移动或调整大小模式循环改

[func (w *windowBase) Event_EXITSIZEMOVE1(pFun WM_EXITSIZEMOVE1) bool {]
ff=线程_退出移动或调整大小模式循环改1

[func (w *windowBase) Event_MOUSEHOVER(pFun WM_MOUSEHOVER) bool {]
ff=线程_鼠标进入消息

[func (w *windowBase) Event_MOUSEHOVER1(pFun WM_MOUSEHOVER1) bool {]
ff=线程_鼠标进入消息1

[func (w *windowBase) Event_MOUSELEAVE(pFun WM_MOUSELEAVE) bool {]
ff=线程_鼠标离开消息

[func (w *windowBase) Event_MOUSELEAVE1(pFun WM_MOUSELEAVE1) bool {]
ff=线程_鼠标离开消息1

[func (w *windowBase) Event_SIZE(pFun WM_SIZE) bool {]
ff=线程_大小改变消息

[func (w *windowBase) Event_SIZE1(pFun WM_SIZE1) bool {]
ff=线程_大小改变消息1

[func (w *windowBase) Event_TIMER(pFun WM_TIMER) bool {]
ff=线程_定时器消息

[func (w *windowBase) Event_TIMER1(pFun WM_TIMER1) bool {]
ff=线程_定时器消息1

[func (w *windowBase) Event_SETFOCUS(pFun WM_SETFOCUS) bool {]
ff=线程_获得焦点

[func (w *windowBase) Event_SETFOCUS1(pFun WM_SETFOCUS1) bool {]
ff=线程_获得焦点1

[func (w *windowBase) Event_KILLFOCUS(pFun WM_KILLFOCUS) bool {]
ff=线程_失去焦点

[func (w *windowBase) Event_KILLFOCUS1(pFun WM_KILLFOCUS1) bool {]
ff=线程_失去焦点1

[func (w *windowBase) Event_KEYDOWN(pFun WM_KEYDOWN) bool {]
ff=线程_键盘按键消息

[func (w *windowBase) Event_KEYDOWN1(pFun WM_KEYDOWN1) bool {]
ff=线程_键盘按键消息1

[func (w *windowBase) Event_CAPTURECHANGED(pFun WM_CAPTURECHANGED) bool {]
ff=线程_鼠标捕获改变消息

[func (w *windowBase) Event_CAPTURECHANGED1(pFun WM_CAPTURECHANGED1) bool {]
ff=线程_鼠标捕获改变消息1

[func (w *windowBase) Event_SETCURSOR(pFun WM_SETCURSOR) bool {]
ff=线程_设置鼠标光标

[func (w *windowBase) Event_SETCURSOR1(pFun WM_SETCURSOR1) bool {]
ff=线程_设置鼠标光标1

[func (w *windowBase) Event_CHAR(pFun WM_CHAR) bool {]
ff=线程_字符消息

[func (w *windowBase) Event_CHAR1(pFun WM_CHAR1) bool {]
ff=线程_字符消息1

[func (w *windowBase) Event_DROPFILES(pFun WM_DROPFILES) bool {]
ff=线程_拖动文件到窗口

[func (w *windowBase) Event_DROPFILES1(pFun WM_DROPFILES1) bool {]
ff=线程_拖动文件到窗口1

[func (w *windowBase) Event_MENU_POPUP(pFun XWM_MENU_POPUP) bool {]
ff=线程_菜单弹出

[func (w *windowBase) Event_MENU_POPUP1(pFun XWM_MENU_POPUP1) bool {]
ff=线程_菜单弹出1

[func (w *windowBase) Event_MENU_POPUP_WND(pFun XWM_MENU_POPUP_WND) bool {]
ff=线程_菜单弹出窗口

[func (w *windowBase) Event_MENU_POPUP_WND1(pFun XWM_MENU_POPUP_WND1) bool {]
ff=线程_菜单弹出窗口1

[func (w *windowBase) Event_MENU_SELECT(pFun XWM_MENU_SELECT) bool {]
ff=线程_菜单选择

[func (w *windowBase) Event_MENU_SELECT1(pFun XWM_MENU_SELECT1) bool {]
ff=线程_菜单选择1

[func (w *windowBase) Event_MENU_EXIT(pFun XWM_MENU_EXIT) bool {]
ff=线程_菜单退出

[func (w *windowBase) Event_MENU_EXIT1(pFun XWM_MENU_EXIT1) bool {]
ff=线程_菜单退出1

[func (w *windowBase) Event_MENU_DRAW_BACKGROUND(pFun XWM_MENU_DRAW_BACKGROUND) bool {]
ff=线程_绘制菜单背景

[func (w *windowBase) Event_MENU_DRAW_BACKGROUND1(pFun XWM_MENU_DRAW_BACKGROUND1) bool {]
ff=线程_绘制菜单背景1

[func (w *windowBase) Event_MENU_DRAWITEM(pFun XWM_MENU_DRAWITEM) bool {]
ff=线程_绘制菜单项事件
[func (w *windowBase) Event_MENU_DRAWITEM1(pFun XWM_MENU_DRAWITEM1) bool {]
ff=线程_绘制菜单项事件1
