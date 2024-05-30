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

[func NewElement(x, y, cx, cy int32, hParent int) *Element {]
ff=创建基础元素
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=y坐标
x=x坐标

[func NewElementByHandle(handle int) *Element {]
ff=创建基础元素并按句柄

[func NewElementByName(name string) *Element {]
ff=创建基础元素并按名称

[func NewElementByUID(nUID int) *Element {]
ff=创建基础元素并按UID

[func NewElementByUIDName(name string) *Element {]
ff=创建基础元素并按UID名称

[func (e *Element) RegEventC(nEvent xcc.XE_, pFun interface{}) bool {]
ff=注册事件C
nEvent=事件类型

[func (e *Element) RegEventC1(nEvent xcc.XE_, pFun interface{}) bool {]
ff=注册事件C1
nEvent=事件类型

[func (e *Element) RemoveEventC(nEvent xcc.XE_, pFun interface{}) bool {]
ff=移除事件C
nEvent=事件类型

[func (e *Element) RegEventCEx(nEvent xcc.XE_, pFun uintptr) bool {]
ff=注册事件CEx
nEvent=事件类型

[func (e *Element) RegEventC1Ex(nEvent xcc.XE_, pFun uintptr) bool {]
ff=注册事件C1Ex
nEvent=事件类型

[func (e *Element) RemoveEventCEx(nEvent xcc.XE_, pFun uintptr) bool {]
ff=移除事件CEx
nEvent=事件类型

[func (e *Element) SendEvent(nEvent xcc.XE_, wParam, lParam uint) int {]
ff=发送事件
nEvent=事件类型

[func (e *Element) PostEvent(nEvent xcc.XE_, wParam, lParam uint) int {]
ff=投递事件
nEvent=事件类型

[func (e *Element) GetRect(pRect *xc.RECT) int {]
ff=取坐标
pRect=坐标

[func (e *Element) GetRectLogic(pRect *xc.RECT) int {]
ff=取逻辑坐标
pRect=坐标

[func (e *Element) GetClientRect(pRect *xc.RECT) int {]
ff=取客户区坐标
pRect=坐标

[func (e *Element) SetWidth(nWidth int) int {]
ff=置宽度
nWidth=宽度

[func (e *Element) SetHeight(nHeight int) int {]
ff=置高度
nHeight=高度

[func (e *Element) GetWidth() int {]
ff=取宽度

[func (e *Element) GetHeight() int {]
ff=取高度

[func (e *Element) RectWndClientToEleClient(pRect *xc.RECT) int {]
ff=窗口客户区坐标到元素客户区坐标
pRect=坐标

[func (e *Element) PointWndClientToEleClient(pPt *xc.POINT) int {]
ff=窗口客户区点到元素客户区
pPt=坐标

[func (e *Element) RectClientToWndClient(pRect *xc.RECT) int {]
ff=客户区坐标到窗口客户区
pRect=坐标

[func (e *Element) PointClientToWndClient(pPt *xc.POINT) int {]
ff=客户区点到窗口客户区
pPt=坐标

[func (e *Element) GetWndClientRect(pRect *xc.RECT) int {]
ff=取窗口客户区坐标
pRect=坐标

[func (e *Element) GetCursor() int {]
ff=取光标

[func (e *Element) SetCursor(hCursor int) int {]
ff=置光标
hCursor=光标句柄

[func (e *Element) AddChild(hChild int) bool {]
ff=添加子对象
hChild=子元素句柄或形状对象句柄

[func (e *Element) InsertChild(hChild int, index int) bool {]
ff=插入子对象
index=插入位置索引
hChild=元素句柄或形状对象句柄

[func (e *Element) SetRect(pRect *xc.RECT, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int {]
ff=置坐标
pRect=坐标

[func (e *Element) SetRectEx(x int, y int, cx int, cy int, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int {]
ff=置坐标EX
nFlags=标识
bRedraw=是否重绘
cy=高度
cx=宽度
y=Y坐标
x=坐标

[func (e *Element) SetRectLogic(pRect *xc.RECT, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int {]
ff=置逻辑坐标
pRect=坐标

[func (e *Element) SetPosition(x, y int32, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int {]
ff=移动
nFlags=标识
bRedraw=是否重绘
y=Y坐标
x=坐标

[func (e *Element) SetPositionLogic(x, y int32, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int {]
ff=移动逻辑坐标
nFlags=标识
bRedraw=是否重绘
y=Y坐标
x=坐标

[func (e *Element) IsDrawFocus() bool {]
ff=判断绘制焦点

[func (e *Element) IsEnable() bool {]
ff=判断启用

[func (e *Element) IsEnableFocus() bool {]
ff=判断启用焦点

[func (e *Element) IsMouseThrough() bool {]
ff=判断鼠标穿透

[func (e *Element) HitChildEle(pPt *xc.POINT) int {]
ff=测试点击元素
pPt=坐标点

[func (e *Element) IsBkTransparent() bool {]
ff=判断背景透明

[func (e *Element) IsEnableEvent_XE_PAINT_END() bool {]
ff=判断启用事件_XE_PAINT_END

[func (e *Element) IsKeyTab() bool {]
ff=判断接受TAB

[func (e *Element) IsSwitchFocus() bool {]
ff=判断接受切换焦点

[func (e *Element) IsEnable_XE_MOUSEWHEEL() bool {]
ff=判断启用_XE_MOUSEWHEEL

[func (e *Element) IsChildEle(hChildEle int) bool {]
ff=判断为子元素
hChildEle=子元素句柄

[func (e *Element) IsEnableCanvas() bool {]
ff=判断启用画布

[func (e *Element) IsFocus() bool {]
ff=判断焦点

[func (e *Element) IsFocusEx() bool {]
ff=判断焦点EX

[func (e *Element) Enable(bEnable bool) int {]
ff=启用
bEnable=启用或禁用

[func (e *Element) EnableFocus(bEnable bool) int {]
ff=启用焦点
bEnable=是否启用

[func (e *Element) EnableDrawFocus(bEnable bool) int {]
ff=启用绘制焦点
bEnable=是否启用

[func (e *Element) EnableDrawBorder(bEnable bool) int {]
ff=启用绘制边框
bEnable=是否启用

[func (e *Element) EnableCanvas(bEnable bool) int {]
ff=启用画布
bEnable=是否启用

[func (e *Element) EnableEvent_XE_PAINT_END(bEnable bool) int {]
ff=启用事件_XE_PAINT_END
bEnable=是否启用

[func (e *Element) EnableBkTransparent(bEnable bool) int {]
ff=启用背景透明
bEnable=是否启用

[func (e *Element) EnableMouseThrough(bEnable bool) int {]
ff=启用鼠标穿透
bEnable=是否启用

[func (e *Element) EnableKeyTab(bEnable bool) int {]
ff=启用接收TAB
bEnable=是否启用

[func (e *Element) EnableSwitchFocus(bEnable bool) int {]
ff=启用切换焦点
bEnable=是否启用

[func (e *Element) EnableEvent_XE_MOUSEWHEEL(bEnable bool) int {]
ff=启用事件_XE_MOUSEWHEEL
bEnable=是否启用

[func (e *Element) Remove() int {]
ff=移除

[func (e *Element) SetZOrder(index int) bool {]
ff=置Z序
index=位置索引

[func (e *Element) SetZOrderEx(hDestEle int, nType xcc.Zorder_) bool {]
ff=置Z序EX
nType=类型
hDestEle=目标元素

[func (e *Element) GetZOrder() int {]
ff=取Z序

[func (e *Element) EnableTopmost(bTopmost bool) bool {]
ff=启用置顶
bTopmost=是否置顶显示

[func (e *Element) Redraw(bImmediate bool) int {]
ff=重绘
bImmediate=是否立即重绘

[func (e *Element) RedrawRect(pRect *xc.RECT, bImmediate bool) int {]
ff=重绘指定区域
pRect=坐标

[func (e *Element) GetChildCount() int {]
ff=取子对象数量

[func (e *Element) GetChildByIndex(index int) int {]
ff=取子对象并按索引
index=索引

[func (e *Element) GetChildByID(nID int) int {]
ff=取子对象并按ID
nID=元素ID

[func (e *Element) SetBorderSize(left int, top int, right int, bottom int) int {]
ff=置边框大小
bottom=下边
right=右边
top=上边
left=左边

[func (e *Element) GetBorderSize(pBorder *xc.RECT) int {]
ff=取边框大小
pBorder=大小

[func (e *Element) SetPadding(left int, top int, right int, bottom int) int {]
ff=置内填充大小
bottom=下边
right=右边
top=上边
left=左边

[func (e *Element) GetPadding(pPadding *xc.RECT) int {]
ff=取内填充大小
pPadding=大小

[func (e *Element) SetDragBorder(nFlags xcc.Element_Position_) int {]
ff=置拖动边框
nFlags=边框位置组合

[func (e *Element) SetDragBorderBindEle(nFlags xcc.Element_Position_, hBindEle int, nSpace int) int {]
ff=置拖动边框绑定元素
nFlags=边框位置标识

[func (e *Element) SetMinSize(nWidth int, nHeight int) int {]
ff=置最小大小
nHeight=最小高度
nWidth=最小宽度

[func (e *Element) SetMaxSize(nWidth int, nHeight int) int {]
ff=置最大大小
nHeight=最大高度
nWidth=最大宽度

[func (e *Element) SetLockScroll(bHorizon bool, bVertical bool) int {]
ff=置锁定滚动
bVertical=是否锁定垂直
bHorizon=是否锁定水平

[func (e *Element) SetTextColor(color int) int {]
ff=置文本颜色
color=ABGR颜色值

[func (e *Element) GetTextColor() int {]
ff=取文本颜色

[func (e *Element) GetTextColorEx() int {]
ff=取文本颜色EX

[func (e *Element) SetFocusBorderColor(color int) int {]
ff=置焦点边框颜色
color=ABGR颜色值

[func (e *Element) GetFocusBorderColor() int {]
ff=取焦点边框颜色

[func (e *Element) SetFont(hFontx int) int {]
ff=置字体
hFontx=炫彩字体

[func (e *Element) GetFont() int {]
ff=取字体

[func (e *Element) GetFontEx() int {]
ff=取字体EX

[func (e *Element) SetAlpha(alpha uint8) int {]
ff=置透明度
alpha=透明度

[func (e *Element) Destroy() int {]
ff=销毁

[func (e *Element) AddBkBorder(nState xcc.CombinedState, color int, width int) int {]
ff=添加背景边框
nState=组合状态

[func (e *Element) AddBkFill(nState xcc.CombinedState, color int) int {]
ff=添加背景填充
nState=组合状态

[func (e *Element) AddBkImage(nState xcc.CombinedState, hImage int) int {]
ff=添加背景图片
nState=组合状态

[func (e *Element) GetBkInfoCount() int {]
ff=取背景对象数量

[func (e *Element) ClearBkInfo() int {]
ff=清空背景对象

[func (e *Element) GetBkManager() int {]
ff=取背景管理器

[func (e *Element) GetBkManagerEx() int {]
ff=取背景管理器EX

[func (e *Element) SetBkManager(hBkInfoM int) int {]
ff=置背景管理器
hBkInfoM=背景管理器

[func (e *Element) GetStateFlags() xcc.CombinedState {]
ff=取状态

[func (e *Element) DrawFocus(hDraw int, pRect *xc.RECT) bool {]
ff=绘制焦点
pRect=区域坐标
hDraw=图形绘制句柄

[func (e *Element) DrawEle(hDraw int) int {]
ff=绘制
hDraw=图形绘制句柄

[func (e *Element) SetUserData(nData int) int {]
ff=置用户数据
nData=用户数据

[func (e *Element) GetUserData() int {]
ff=取用户数据

[func (e *Element) GetContentSize(bHorizon bool, cx int, cy int, pSize *xc.SIZE) int {]
ff=取内容大小
pSize=返回大小
cy=高度
cx=宽度
bHorizon=水平或垂直

[func (e *Element) SetCapture(b bool) int {]
ff=置鼠标捕获
b=开启

[func (e *Element) EnableTransparentChannel(bEnable bool) int {]
ff=启用透明通道
bEnable=启用或关闭

[func (e *Element) SetXCTimer(nIDEvent int, uElapse int) bool {]
ff=置炫彩定时器
uElapse=延时毫秒
nIDEvent=事件ID

[func (e *Element) KillXCTimer(nIDEvent int) bool {]
ff=关闭炫彩定时器
nIDEvent=事件ID

[func (e *Element) SetToolTip(pText string) int {]
ff=置工具提示
pText=内容

[func (e *Element) SetToolTipEx(pText string, nTextAlign xcc.TextFormatFlag_) int {]
ff=置工具提示EX
nTextAlign=对齐方式
pText=内容

[func (e *Element) GetToolTip() int {]
ff=取工具提示

[func (e *Element) PopupToolTip(x int, y int) int {]
ff=弹出工具提示
y=Y坐标
x=坐标

[func (e *Element) AdjustLayout(nAdjustNo uint32) int {]
ff=调整布局
nAdjustNo=流水号

[func (e *Element) AdjustLayoutEx(nFlags xcc.AdjustLayout_, nAdjustNo uint32) int {]
ff=调整布局EX
nFlags=标识

[func (e *Element) GetAlpha() byte {]
ff=取透明度

[func (e *Element) GetPosition(pOutX, pOutY *int32) int {]
ff=取位置
pOutY=返回Y坐标
pOutX=返回X坐标

[func (e *Element) SetSize(nWidth, nHeight int32, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int {]
ff=置大小
nFlags=标识
bRedraw=是否重绘
nHeight=高度
nWidth=宽度

[func (e *Element) GetSize(pOutWidth, pOutHeight *int32) int {]
ff=取大小
pOutHeight=返回高度
pOutWidth=返回宽度

[func (e *Element) SetBkInfo(pText string) int {]
ff=置背景
pText=字符串

[func (e *Element) GetWndClientRectDPI(pRect *xc.RECT) int {]
ff=取窗口客户区坐标DPI
pRect=接收返回坐标

[func (e *Element) PointClientToWndClientDPI(pPt *xc.POINT) int {]
ff=取窗口客户区坐标点DPI
pPt=接收返回坐标点

[func (e *Element) RectClientToWndClientDPI(pRect *xc.RECT) int {]
ff=客户区坐标到窗口客户区DPI
pRect=接收返回坐标

[func (e *Element) SetFocus() bool {]
ff=置焦点

[func (e *Element) GetLeft() int32 {]
ff=取左边

[func (e *Element) GetTop() int32 {]
ff=取顶边

[func (e *Element) GetRight() int32 {]
ff=取右边

[func (e *Element) GetBottom() int32 {]
ff=取底边

[func (e *Element) SetLeft(x int32, bRedraw bool) bool {]
ff=置左边
bRedraw=是否重绘
x=左边x坐标

[func (e *Element) SetTop(y int32, bRedraw bool) bool {]
ff=置顶边
bRedraw=是否重绘
y=顶边y坐标

[func (e *Element) Event_ELEPROCE(pFun XE_ELEPROCE) bool {]
ff=事件_处理过程

[func (e *Element) Event_ELEPROCE1(pFun XE_ELEPROCE1) bool {]
ff=事件_处理过程1

[func (e *Element) Event_PAINT(pFun XE_PAINT) bool {]
ff=事件_绘制事件

[func (e *Element) Event_PAINT1(pFun XE_PAINT1) bool {]
ff=事件_绘制事件1

[func (e *Element) Event_PAINT_END(pFun XE_PAINT_END) bool {]
ff=事件_绘制完成

[func (e *Element) Event_PAINT_END1(pFun XE_PAINT_END1) bool {]
ff=事件_绘制完成1

[func (e *Element) Event_PAINT_SCROLLVIEW(pFun XE_PAINT_SCROLLVIEW) bool {]
ff=事件_滚动视图

[func (e *Element) Event_PAINT_SCROLLVIEW1(pFun XE_PAINT_SCROLLVIEW1) bool {]
ff=事件_滚动视图1

[func (e *Element) Event_MOUSEMOVE(pFun XE_MOUSEMOVE) bool {]
ff=事件_鼠标移动

[func (e *Element) Event_MOUSEMOVE1(pFun XE_MOUSEMOVE1) bool {]
ff=事件_鼠标移动1

[func (e *Element) Event_MOUSESTAY(pFun XE_MOUSESTAY) bool {]
ff=事件_鼠标进入

[func (e *Element) Event_MOUSESTAY1(pFun XE_MOUSESTAY1) bool {]
ff=事件_鼠标进入1

[func (e *Element) Event_MOUSEHOVER(pFun XE_MOUSEHOVER) bool {]
ff=事件_鼠标悬停

[func (e *Element) Event_MOUSEHOVER1(pFun XE_MOUSEHOVER1) bool {]
ff=事件_鼠标悬停1

[func (e *Element) Event_MOUSELEAVE(pFun XE_MOUSELEAVE) bool {]
ff=事件_鼠标离开

[func (e *Element) Event_MOUSELEAVE1(pFun XE_MOUSELEAVE1) bool {]
ff=事件_鼠标离开1

[func (e *Element) Event_MOUSEWHEEL(pFun XE_MOUSEWHEEL) bool {]
ff=事件_鼠标滚轮滚动

[func (e *Element) Event_MOUSEWHEEL1(pFun XE_MOUSEWHEEL1) bool {]
ff=事件_鼠标滚轮滚动1

[func (e *Element) Event_LBUTTONDOWN(pFun XE_LBUTTONDOWN) bool {]
ff=事件_左键按下

[func (e *Element) Event_LBUTTONDOWN1(pFun XE_LBUTTONDOWN1) bool {]
ff=事件_左键按下1

[func (e *Element) Event_LBUTTONUP(pFun XE_LBUTTONUP) bool {]
ff=事件_左键弹起

[func (e *Element) Event_LBUTTONUP1(pFun XE_LBUTTONUP1) bool {]
ff=事件_左键弹起1

[func (e *Element) Event_RBUTTONDOWN(pFun XE_RBUTTONDOWN) bool {]
ff=事件_右键按下

[func (e *Element) Event_RBUTTONDOWN1(pFun XE_RBUTTONDOWN1) bool {]
ff=事件_右键按下1

[func (e *Element) Event_RBUTTONUP(pFun XE_RBUTTONUP) bool {]
ff=事件_右键弹起

[func (e *Element) Event_RBUTTONUP1(pFun XE_RBUTTONUP1) bool {]
ff=事件_右键弹起1

[func (e *Element) Event_LBUTTONDBCLICK(pFun XE_LBUTTONDBCLICK) bool {]
ff=事件_左键双击

[func (e *Element) Event_LBUTTONDBCLICK1(pFun XE_LBUTTONDBCLICK1) bool {]
ff=事件_左键双击1

[func (e *Element) Event_XC_TIMER(pFun XE_XC_TIMER) bool {]
ff=事件_炫彩定时器

[func (e *Element) Event_XC_TIMER1(pFun XE_XC_TIMER1) bool {]
ff=事件_炫彩定时器1

[func (e *Element) Event_ADJUSTLAYOUT(pFun XE_ADJUSTLAYOUT) bool {]
ff=事件_调整布局

[func (e *Element) Event_ADJUSTLAYOUT1(pFun XE_ADJUSTLAYOUT1) bool {]
ff=事件_调整布局1

[func (e *Element) Event_ADJUSTLAYOUT_END(pFun XE_ADJUSTLAYOUT_END) bool {]
ff=事件_调整布局完成

[func (e *Element) Event_ADJUSTLAYOUT_END1(pFun XE_ADJUSTLAYOUT_END1) bool {]
ff=事件_调整布局完成1

[func (e *Element) Event_SETFOCUS(pFun XE_SETFOCUS) bool {]
ff=事件_获得焦点

[func (e *Element) Event_SETFOCUS1(pFun XE_SETFOCUS1) bool {]
ff=事件_获得焦点1

[func (e *Element) Event_KILLFOCUS(pFun XE_KILLFOCUS) bool {]
ff=事件_失去焦点

[func (e *Element) Event_KILLFOCUS1(pFun XE_KILLFOCUS1) bool {]
ff=事件_失去焦点1

[func (e *Element) Event_DESTROY(pFun XE_DESTROY) bool {]
ff=事件_即将销毁

[func (e *Element) Event_DESTROY1(pFun XE_DESTROY1) bool {]
ff=事件_即将销毁1

[func (e *Element) Event_DESTROY_END(pFun XE_DESTROY_END) bool {]
ff=事件_销毁完成

[func (e *Element) Event_DESTROY_END1(pFun XE_DESTROY_END1) bool {]
ff=事件_销毁完成1

[func (e *Element) Event_SIZE(pFun XE_SIZE) bool {]
ff=事件_大小改变事件

[func (e *Element) Event_SIZE1(pFun XE_SIZE1) bool {]
ff=事件_大小改变事件1

[func (e *Element) Event_SHOW(pFun XE_SHOW) bool {]
ff=事件_显示隐藏事件

[func (e *Element) Event_SHOW1(pFun XE_SHOW1) bool {]
ff=事件_显示隐藏事件1

[func (e *Element) Event_SETFONT(pFun XE_SETFONT) bool {]
ff=事件_设置字体

[func (e *Element) Event_SETFONT1(pFun XE_SETFONT1) bool {]
ff=事件_设置字体1

[func (e *Element) Event_KEYDOWN(pFun XE_KEYDOWN) bool {]
ff=事件_按键按下

[func (e *Element) Event_KEYDOWN1(pFun XE_KEYDOWN1) bool {]
ff=事件_按键按下1

[func (e *Element) Event_KEYUP(pFun XE_KEYUP) bool {]
ff=事件_按键弹起

[func (e *Element) Event_KEYUP1(pFun XE_KEYUP1) bool {]
ff=事件_按键弹起1

[func (e *Element) Event_CHAR(pFun XE_CHAR) bool {]
ff=事件_CHAR

[func (e *Element) Event_CHAR1(pFun XE_CHAR1) bool {]
ff=事件_CHAR1

[func (e *Element) Event_SETCAPTURE(pFun XE_SETCAPTURE) bool {]
ff=事件_设置鼠标捕获

[func (e *Element) Event_SETCAPTURE1(pFun XE_SETCAPTURE1) bool {]
ff=事件_设置鼠标捕获1

[func (e *Element) Event_KILLCAPTURE(pFun XE_KILLCAPTURE) bool {]
ff=事件_失去鼠标捕获

[func (e *Element) Event_KILLCAPTURE1(pFun XE_KILLCAPTURE1) bool {]
ff=事件_失去鼠标捕获1

[func (e *Element) Event_SETCURSOR(pFun XE_SETCURSOR) bool {]
ff=事件_设置鼠标光标

[func (e *Element) Event_SETCURSOR1(pFun XE_SETCURSOR1) bool {]
ff=事件_设置鼠标光标1

[func (e *Element) Event_DROPFILES(pFun XE_DROPFILES) bool {]
ff=事件_文件拖放

[func (e *Element) Event_DROPFILES1(pFun XE_DROPFILES1) bool {]
ff=事件_文件拖放事件1

[func (e *Element) Event_TOOLTIP_POPUP(pFun XE_TOOLTIP_POPUP) bool {]
ff=事件_工具提示弹出

[func (e *Element) Event_TOOLTIP_POPUP1(pFun XE_TOOLTIP_POPUP1) bool {]
ff=事件_工具提示弹出1

[func (e *Element) Event_MENU_SELECT(pFun XE_MENU_SELECT) bool {]
ff=事件_弹出菜单项被选择

[func (e *Element) Event_MENU_POPUP(pFun XE_MENU_POPUP) bool {]
ff=事件_菜单弹出

[func (e *Element) Event_MENU_EXIT(pFun XE_MENU_EXIT) bool {]
ff=事件_菜单退出

[func (e *Element) Event_MENU_POPUP_WND(pFun XE_MENU_POPUP_WND) bool {]
ff=事件_菜单弹出窗口

[func (e *Element) Event_MENU_DRAW_BACKGROUND(pFun XE_MENU_DRAW_BACKGROUND) bool {]
ff=事件_绘制菜单背景

[func (e *Element) Event_MENU_DRAWITEM(pFun XE_MENU_DRAWITEM) bool {]
ff=事件_绘制菜单项事件

[func (e *Element) Event_MENU_SELECT1(pFun XE_MENU_SELECT1) bool {]
ff=事件_弹出菜单项被选择1

[func (e *Element) Event_MENU_POPUP1(pFun XE_MENU_POPUP1) bool {]
ff=事件_菜单弹出1

[func (e *Element) Event_MENU_EXIT1(pFun XE_MENU_EXIT1) bool {]
ff=事件_菜单退出1

[func (e *Element) Event_MENU_POPUP_WND1(pFun XE_MENU_POPUP_WND1) bool {]
ff=菜单弹出窗口1

[func (e *Element) Event_MENU_DRAW_BACKGROUND1(pFun XE_MENU_DRAW_BACKGROUND1) bool {]
ff=事件_绘制菜单背景1

[func (e *Element) Event_MENU_DRAWITEM1(pFun XE_MENU_DRAWITEM1) bool {]
ff=事件_绘制菜单项1
