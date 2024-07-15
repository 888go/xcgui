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

[func XEle_Create(x, y, cx, cy int32, hParent int) int {]
ff=元素_创建
hParent=父窗口句柄或元素句柄
cy=高度
cx=宽度
y=y坐标
x=x坐标

[func XEle_RegEventC(hEle int, nEvent xcc.XE_, pFun interface{}) bool {]
ff=元素_注册事件C
nEvent=事件类型
hEle=元素句柄

[func XEle_RegEventC1(hEle int, nEvent xcc.XE_, pFun interface{}) bool {]
ff=元素_注册事件C1
nEvent=事件类型
hEle=元素句柄

[func XEle_RemoveEventC(hEle int, nEvent xcc.XE_, pFun interface{}) bool {]
ff=元素_移除事件C
nEvent=事件类型
hEle=元素句柄

[func XEle_RegEventCEx(hEle int, nEvent xcc.XE_, pFun uintptr) bool {]
ff=元素_注册事件CEx
nEvent=事件类型
hEle=元素句柄

[func XEle_RegEventC1Ex(hEle int, nEvent xcc.XE_, pFun uintptr) bool {]
ff=元素_注册事件C1Ex
nEvent=事件类型
hEle=元素句柄

[func XEle_RemoveEventCEx(hEle int, nEvent xcc.XE_, pFun uintptr) bool {]
ff=元素_移除事件CEx
nEvent=事件类型
hEle=元素句柄

[func XEle_SendEvent(hEle int, nEvent xcc.XE_, wParam, lParam uint) int {]
ff=元素_发送事件
nEvent=事件类型
hEle=元素句柄

[func XEle_PostEvent(hEle int, nEvent xcc.XE_, wParam, lParam uint) int {]
ff=元素_投递事件
nEvent=事件类型
hEle=元素句柄

[func XEle_GetRect(hEle int, pRect *RECT) int {]
ff=元素_取坐标
pRect=坐标
hEle=元素句柄

[func XEle_GetRectLogic(hEle int, pRect *RECT) int {]
ff=元素_取逻辑坐标
pRect=坐标
hEle=元素句柄

[func XEle_GetClientRect(hEle int, pRect *RECT) int {]
ff=元素_取客户区坐标
pRect=坐标
hEle=元素句柄

[func XEle_SetWidth(hEle int, nWidth int) int {]
ff=元素_置宽度
nWidth=宽度
hEle=元素句柄

[func XEle_SetHeight(hEle int, nHeight int) int {]
ff=元素_置高度
nHeight=高度
hEle=元素句柄

[func XEle_GetWidth(hEle int) int {]
ff=元素_取宽度
hEle=元素句柄

[func XEle_GetHeight(hEle int) int {]
ff=元素_取高度
hEle=元素句柄

[func XEle_RectWndClientToEleClient(hEle int, pRect *RECT) int {]
ff=元素_窗口客户区坐标到元素客户区坐标
pRect=坐标
hEle=元素句柄

[func XEle_PointWndClientToEleClient(hEle int, pPt *POINT) int {]
ff=元素_窗口客户区点到元素客户区
pPt=坐标
hEle=元素句柄

[func XEle_RectClientToWndClient(hEle int, pRect *RECT) int {]
ff=元素_客户区坐标到窗口客户区
pRect=坐标
hEle=元素句柄

[func XEle_PointClientToWndClient(hEle int, pPt *POINT) int {]
ff=元素_客户区点到窗口客户区
pPt=坐标
hEle=元素句柄

[func XEle_GetWndClientRect(hEle int, pRect *RECT) int {]
ff=元素_基于窗口客户区坐标
pRect=坐标
hEle=元素句柄

[func XEle_GetCursor(hEle int) int {]
ff=元素_取光标
hEle=元素句柄

[func XEle_SetCursor(hEle int, hCursor int) int {]
ff=元素_置光标
hCursor=光标句柄
hEle=元素句柄

[func XEle_AddChild(hEle int, hChild int) bool {]
ff=元素_添加子对象
hChild=要添加的子元素句柄或形状对象句柄
hEle=元素句柄

[func XEle_InsertChild(hEle int, hChild int, index int) bool {]
ff=元素_插入子对象
index=插入位置索引
hChild=要插入的元素句柄或形状对象句柄
hEle=元素句柄

[func XEle_SetRect(hEle int, pRect *RECT, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int {]
ff=元素_置坐标
nFlags=调整布局标识位
bRedraw=是否重绘
pRect=坐标
hEle=元素句柄

[func XEle_SetRectEx(hEle int, x int, y int, cx int, cy int, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int {]
ff=元素_置坐标EX
nFlags=调整布局标识位
bRedraw=是否重绘
cy=高度
cx=宽度
y=Y坐标
x=坐标
hEle=元素句柄

[func XEle_SetRectLogic(hEle int, pRect *RECT, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int {]
ff=元素_置逻辑坐标
nFlags=调整布局标识位
bRedraw=是否重绘
pRect=坐标
hEle=元素句柄

[func XEle_SetPosition(hEle int, x, y int32, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int {]
ff=元素_移动
nFlags=调整布局标识位
bRedraw=是否重绘
y=Y坐标
x=坐标
hEle=元素句柄

[func XEle_SetPositionLogic(hEle int, x, y int32, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int {]
ff=元素_移动逻辑坐标
nFlags=调整布局标识位
bRedraw=是否重绘
y=Y坐标
x=坐标
hEle=元素句柄

[func XEle_IsDrawFocus(hEle int) bool {]
ff=元素_判断绘制焦点
hEle=元素句柄

[func XEle_IsEnable(hEle int) bool {]
ff=元素_判断启用
hEle=元素句柄

[func XEle_IsEnableFocus(hEle int) bool {]
ff=元素_判断启用焦点
hEle=元素句柄

[func XEle_IsMouseThrough(hEle int) bool {]
ff=元素_判断鼠标穿透
hEle=元素句柄

[func XEle_HitChildEle(hEle int, pPt *POINT) int {]
ff=元素_测试点击元素
pPt=坐标点
hEle=元素句柄

[func XEle_IsBkTransparent(hEle int) bool {]
ff=元素_判断背景透明
hEle=元素句柄

[func XEle_IsEnableEvent_XE_PAINT_END(hEle int) bool {]
ff=元素_判断启用事件_XE_PAINT_END
hEle=元素句柄

[func XEle_IsKeyTab(hEle int) bool {]
ff=元素_判断接受TAB
hEle=元素句柄

[func XEle_IsSwitchFocus(hEle int) bool {]
ff=元素_判断接受切换焦点
hEle=元素句柄

[func XEle_IsEnable_XE_MOUSEWHEEL(hEle int) bool {]
ff=元素_判断启用_XE_MOUSEWHEEL
hEle=元素句柄

[func XEle_IsChildEle(hEle int, hChildEle int) bool {]
ff=元素_判断为子元素
hChildEle=子元素句柄
hEle=元素句柄

[func XEle_IsEnableCanvas(hEle int) bool {]
ff=元素_判断启用画布
hEle=元素句柄

[func XEle_IsFocus(hEle int) bool {]
ff=元素_判断焦点
hEle=元素句柄

[func XEle_IsFocusEx(hEle int) bool {]
ff=元素_判断焦点EX
hEle=元素句柄

[func XEle_Enable(hEle int, bEnable bool) int {]
ff=元素_启用
bEnable=是否启用
hEle=元素句柄

[func XEle_EnableFocus(hEle int, bEnable bool) int {]
ff=元素_启用焦点
bEnable=是否启用
hEle=元素句柄

[func XEle_EnableDrawFocus(hEle int, bEnable bool) int {]
ff=元素_启用绘制焦点
bEnable=是否启用
hEle=元素句柄

[func XEle_EnableDrawBorder(hEle int, bEnable bool) int {]
ff=元素_启用绘制边框
bEnable=是否启用
hEle=元素句柄

[func XEle_EnableCanvas(hEle int, bEnable bool) int {]
ff=元素_启用画布
bEnable=是否启用
hEle=元素句柄

[func XEle_EnableEvent_XE_PAINT_END(hEle int, bEnable bool) int {]
ff=元素_启用事件_XE_PAINT_END
bEnable=是否启用
hEle=元素句柄

[func XEle_EnableBkTransparent(hEle int, bEnable bool) int {]
ff=元素_启用背景透明
bEnable=是否启用
hEle=元素句柄

[func XEle_EnableMouseThrough(hEle int, bEnable bool) int {]
ff=元素_启用鼠标穿透
bEnable=是否启用
hEle=元素句柄

[func XEle_EnableKeyTab(hEle int, bEnable bool) int {]
ff=元素_启用接收TAB
bEnable=是否启用
hEle=元素句柄

[func XEle_EnableSwitchFocus(hEle int, bEnable bool) int {]
ff=元素_启用切换焦点
bEnable=是否启用
hEle=元素句柄

[func XEle_EnableEvent_XE_MOUSEWHEEL(hEle int, bEnable bool) int {]
ff=元素_启用事件_XE_MOUSEWHEEL
bEnable=是否启用
hEle=元素句柄

[func XEle_Remove(hEle int) int {]
ff=元素_移除
hEle=元素句柄

[func XEle_SetZOrder(hEle int, index int) bool {]
ff=元素_置Z序
index=位置索引
hEle=元素句柄

[func XEle_SetZOrderEx(hEle int, hDestEle int, nType xcc.Zorder_) bool {]
ff=元素_置Z序EX
nType=类型
hDestEle=目标元素
hEle=元素句柄

[func XEle_GetZOrder(hEle int) int {]
ff=元素_取Z序
hEle=元素句柄

[func XEle_EnableTopmost(hEle int, bTopmost bool) bool {]
ff=元素_启用置顶
bTopmost=是否置顶显示
hEle=元素句柄

[func XEle_Redraw(hEle int, bImmediate bool) int {]
ff=元素_重绘
bImmediate=是否立即重绘
hEle=元素句柄

[func XEle_RedrawRect(hEle int, pRect *RECT, bImmediate bool) int {]
ff=元素_重绘指定区域
bImmediate=是否立即重绘
pRect=相对于元素客户区坐标
hEle=元素句柄

[func XEle_GetChildCount(hEle int) int {]
ff=元素_取子对象数量
hEle=元素句柄

[func XEle_GetChildByIndex(hEle int, index int) int {]
ff=元素_取子对象从索引
index=索引
hEle=元素句柄

[func XEle_GetChildByID(hEle int, nID int) int {]
ff=元素_取子对象从ID
nID=元素ID
hEle=元素句柄

[func XEle_SetBorderSize(hEle int, left int, top int, right int, bottom int) int {]
ff=元素_置边框大小
bottom=下
right=右
top=上
left=左
hEle=元素句柄

[func XEle_GetBorderSize(hEle int, pBorder *RECT) int {]
ff=元素_取边框大小
pBorder=大小
hEle=元素句柄

[func XEle_SetPadding(hEle int, left int, top int, right int, bottom int) int {]
ff=元素_置内填充大小
bottom=下
right=右
top=上
left=左
hEle=元素句柄

[func XEle_GetPadding(hEle int, pPadding *RECT) int {]
ff=元素_取内填充大小
pPadding=大小
hEle=元素句柄

[func XEle_SetDragBorder(hEle int, nFlags xcc.Element_Position_) int {]
ff=元素_置拖动边框
nFlags=边框位置组合
hEle=元素句柄

[func XEle_SetDragBorderBindEle(hEle int, nFlags xcc.Element_Position_, hBindEle int, nSpace int) int {]
ff=元素_置拖动边框绑定元素
nFlags=边框位置标识
hEle=元素句柄

[func XEle_SetMinSize(hEle int, nWidth int, nHeight int) int {]
ff=元素_置最小大小
nHeight=最小高度
nWidth=最小宽度
hEle=元素句柄

[func XEle_SetMaxSize(hEle int, nWidth int, nHeight int) int {]
ff=元素_置最大大小
nHeight=最大高度
nWidth=最大宽度
hEle=元素句柄

[func XEle_SetLockScroll(hEle int, bHorizon bool, bVertical bool) int {]
ff=元素_置锁定滚动
bVertical=是否锁定垂直滚动
bHorizon=是否锁定水平滚动
hEle=元素句柄

[func XEle_SetTextColor(hEle int, color int) int {]
ff=元素_置文本颜色
color=ABGR颜色值
hEle=元素句柄

[func XEle_GetTextColor(hEle int) int {]
ff=元素_取文本颜色
hEle=元素句柄

[func XEle_GetTextColorEx(hEle int) int {]
ff=元素_取文本颜色EX
hEle=元素句柄

[func XEle_SetFocusBorderColor(hEle int, color int) int {]
ff=元素_置焦点边框颜色
color=ABGR颜色值
hEle=元素句柄

[func XEle_GetFocusBorderColor(hEle int) int {]
ff=元素_取焦点边框颜色
hEle=元素句柄

[func XEle_SetFont(hEle int, hFontx int) int {]
ff=元素_置字体
hFontx=炫彩字体
hEle=元素句柄

[func XEle_GetFont(hEle int) int {]
ff=元素_取字体
hEle=元素句柄

[func XEle_GetFontEx(hEle int) int {]
ff=元素_取字体EX
hEle=元素句柄

[func XEle_SetAlpha(hEle int, alpha uint8) int {]
ff=元素_置透明度
hEle=元素句柄

[func XEle_Destroy(hEle int) int {]
ff=元素_销毁
hEle=元素句柄

[func XEle_AddBkBorder(hEle int, nState xcc.CombinedState, color int, width int) int {]
ff=元素_添加背景边框
nState=组合状态
hEle=元素句柄

[func XEle_AddBkFill(hEle int, nState xcc.CombinedState, color int) int {]
ff=元素_添加背景填充
nState=组合状态
hEle=元素句柄

[func XEle_AddBkImage(hEle int, nState xcc.CombinedState, hImage int) int {]
ff=元素_添加背景图片
nState=组合状态
hEle=元素句柄

[func XEle_GetBkInfoCount(hEle int) int {]
ff=元素_取背景对象数量
hEle=元素句柄

[func XEle_ClearBkInfo(hEle int) int {]
ff=元素_清空背景对象
hEle=元素句柄

[func XEle_GetBkManager(hEle int) int {]
ff=元素_取背景管理器
hEle=元素句柄

[func XEle_GetBkManagerEx(hEle int) int {]
ff=元素_取背景管理器EX
hEle=元素句柄

[func XEle_SetBkManager(hEle int, hBkInfoM int) int {]
ff=元素_置背景管理器
hBkInfoM=背景管理器
hEle=元素句柄

[func XEle_GetStateFlags(hEle int) xcc.CombinedState {]
ff=元素_取状态
hEle=元素句柄

[func XEle_DrawFocus(hEle int, hDraw int, pRect *RECT) bool {]
ff=元素_绘制焦点
pRect=区域坐标
hDraw=图形绘制句柄
hEle=元素句柄

[func XEle_DrawEle(hEle int, hDraw int) int {]
ff=元素_绘制
hDraw=图形绘制句柄
hEle=元素句柄

[func XEle_SetUserData(hEle int, nData int) int {]
ff=元素_置用户数据
nData=用户数据
hEle=元素句柄

[func XEle_GetUserData(hEle int) int {]
ff=元素_取用户数据
hEle=元素句柄

[func XEle_GetContentSize(hEle int, bHorizon bool, cx int, cy int, pSize *SIZE) int {]
ff=元素_取内容大小
pSize=返回大小
cy=高度
cx=宽度
bHorizon=水平或垂直
hEle=元素句柄

[func XEle_SetCapture(hEle int, b bool) int {]
ff=元素_置鼠标捕获
b=启用或关闭
hEle=元素句柄

[func XEle_EnableTransparentChannel(hEle int, bEnable bool) int {]
ff=元素_启用透明通道
bEnable=是否启用
hEle=元素句柄

[func XEle_SetXCTimer(hEle int, nIDEvent int, uElapse int) bool {]
ff=元素_置炫彩定时器
uElapse=延时毫秒
nIDEvent=事件ID
hEle=元素句柄

[func XEle_KillXCTimer(hEle int, nIDEvent int) bool {]
ff=元素_关闭炫彩定时器
nIDEvent=事件ID
hEle=元素句柄

[func XEle_SetToolTip(hEle int, pText string) int {]
ff=元素_置工具提示
pText=工具提示内容
hEle=元素句柄

[func XEle_SetToolTipEx(hEle int, pText string, nTextAlign xcc.TextFormatFlag_) int {]
ff=元素_置工具提示EX
nTextAlign=文本对齐方式
pText=工具提示内容
hEle=元素句柄

[func XEle_GetToolTip(hEle int) int {]
ff=元素_取工具提示
hEle=元素句柄

[func XEle_PopupToolTip(hEle int, x int, y int) int {]
ff=元素_弹出工具提示
y=Y坐标
x=坐标
hEle=元素句柄

[func XEle_AdjustLayout(hEle int, nAdjustNo uint32) int {]
ff=元素_调整布局
nAdjustNo=调整布局流水号
hEle=元素句柄

[func XEle_AdjustLayoutEx(hEle int, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int {]
ff=元素_调整布局EX
nFlags=调整标识
hEle=元素句柄

[func XEle_GetAlpha(hEle int) byte {]
ff=元素_取透明度
hEle=元素句柄

[func XEle_GetPosition(hEle int, pOutX, pOutY *int32) int {]
ff=元素_取位置
pOutY=返回Y坐标
pOutX=返回X坐标
hEle=元素句柄

[func XEle_SetSize(hEle int, nWidth, nHeight int32, bRedraw bool, nFlags xcc.AdjustLayout_, nAdjustNo uint32) int {]
ff=元素_置大小
nFlags=调整布局标识位
bRedraw=是否重绘
nHeight=高度
nWidth=宽度
hEle=元素句柄

[func XEle_GetSize(hEle int, pOutWidth, pOutHeight *int32) int {]
ff=元素_取大小
pOutHeight=返回高度
pOutWidth=返回宽度
hEle=元素句柄

[func XEle_SetBkInfo(hEle int, pText string) int {]
ff=元素_置背景
pText=背景内容字符串
hEle=元素句柄

[func XEle_GetWndClientRectDPI(hEle int, pRect *RECT) int {]
ff=元素_取窗口客户区坐标DPI
pRect=接收返回坐标
hEle=元素句柄

[func XEle_PointClientToWndClientDPI(hEle int, pPt *POINT) int {]
ff=元素_取窗口客户区坐标点DPI
pPt=接收返回坐标点
hEle=元素句柄

[func XEle_RectClientToWndClientDPI(hEle int, pRect *RECT) int {]
ff=元素_客户区坐标到窗口客户区DPI
pRect=接收返回坐标
hEle=元素句柄
