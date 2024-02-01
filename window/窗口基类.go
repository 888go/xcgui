package window

import (
	"e.coding.net/gogit/go/xcgui/objectbase"
	"e.coding.net/gogit/go/xcgui/wnd"
	"e.coding.net/gogit/go/xcgui/xc"
	"e.coding.net/gogit/go/xcgui/xcc"
)

// windowBase 窗口基类.
type windowBase struct {
	objectbase.UI
}

// I消息框 炫彩_消息框.
//	@param pTitle 标题.
//	@param pText 内容文本.
//	@param nFlags 标识: xcc.I常量_弹出消息框_.
//	@param XCStyle xcc.I常量_窗口样式_.
//	@return xcc.I常量_弹出消息框_. 返回: xcc.I常量_弹出消息框_确定按钮: 点击确定按钮退出. xcc.I常量_弹出消息框_取消按钮: 点击取消按钮退出. xcc.I常量_弹出消息框_其他: 其他方式退出.
//
func (w *windowBase) I消息框(pTitle, pText string, nFlags xcc.I常量_弹出消息框_, XCStyle xcc.I常量_窗口样式_) xcc.I常量_弹出消息框_ {
	return xc.XC_MessageBox(pTitle, pText, nFlags, w.I取HWND(), XCStyle)
}

// I消息框_创建 消息框_创建, 然后请调用 DoModal() 方法显示模态窗口.
//	@param pTitle 标题.
//	@param pText 内容文本.
//	@param nFlags 标识: xcc.I常量_弹出消息框_.
//	@param XCStyle xcc.I常量_窗口样式_.
//	@return *ModalWindow 模态窗口对象.
//
func (w *windowBase) I消息框_创建(pTitle, pText string, nFlags xcc.I常量_弹出消息框_, XCStyle xcc.I常量_窗口样式_) *ModalWindow {
	p := &ModalWindow{}
	p.SetHandle(xc.XMsg_Create(pTitle, pText, nFlags, w.I取HWND(), XCStyle))
	return p
}

// I消息框_创建扩展 消息框_创建扩展, 然后请调用 DoModal() 方法显示模态窗口.
//	@param dwExStyle 窗口扩展样式.
//	@param dwStyle 窗口样式.
//	@param lpClassName 窗口类名.
//	@param pTitle 标题.
//	@param pText 内容文本.
//	@param nFlags 标识: xcc.I常量_弹出消息框_.
//	@param XCStyle xcc.I常量_窗口样式_.
//	@return *ModalWindow 模态窗口对象.
//
func (w *windowBase) I消息框_创建扩展(dwExStyle, dwStyle int, lpClassName, pTitle, pText string, nFlags xcc.I常量_弹出消息框_, XCStyle xcc.I常量_窗口样式_) *ModalWindow {
	p := &ModalWindow{}
	p.SetHandle(xc.XMsg_CreateEx(dwExStyle, dwStyle, lpClassName, pTitle, pText, nFlags, w.I取HWND(), XCStyle))
	return p
}

// 炫彩_发送窗口消息.
//
// msg:.
//
// wParam:.
//
// lParam:.
func (w *windowBase) I发送窗口消息(msg int, wParam int, lParam int) int {
	return xc.XC_SendMessage(w.I句柄, msg, wParam, lParam)
}

// 炫彩_投递窗口消息.
//
// msg:.
//
// wParam:.
//
// lParam:.
func (w *windowBase) I投递窗口消息(msg int, wParam int, lParam int) bool {
	return xc.XC_PostMessage(w.I句柄, msg, wParam, lParam)
}

// 炫彩_判断窗口, 判断是否为窗口句柄.
func (w *windowBase) I判断窗口() bool {
	return xc.XC_IsHWINDOW(w.I句柄)
}

// 炫彩_取对象从ID, 通过ID获取对象句柄, 不包括窗口对象.
//
// nID: ID值.
func (w *windowBase) I取对象从ID(nID int) int {
	return xc.XC_GetObjectByID(w.I句柄, nID)
}

// 炫彩_取对象从ID名称, 通过ID名称获取对象句柄.
//
// pName: ID名称.
func (w *windowBase) I取对象从ID名称(pName string) int {
	return xc.XC_GetObjectByIDName(w.I句柄, pName)
}

// 窗口_显示.
//
// nCmdShow: 显示方式: xcc.I常量_窗口形式_.
func (w *windowBase) I显示(显示方式 xcc.I常量_窗口形式_) int {
	return xc.XWnd_ShowWindow(w.I句柄, 显示方式)
}

// 窗口_置顶.
func (w *windowBase) I置顶() int {
	return xc.XWnd_SetTop(w.I句柄)
}

// 窗口_注册事件C.
//
// nEvent: 事件类型: xcc.WM_, xcc.XWM_.
//
// pFun: 事件函数.
func (w *windowBase) I注册事件C(nEvent xcc.WM_, pFun interface{}) bool {
	return xc.XWnd_RegEventC(w.I句柄, nEvent, pFun)
}

// 窗口_注册事件C1.
//
// nEvent: 事件类型: xcc.WM_, xcc.XWM_.
//
// pFun: 事件函数.
func (w *windowBase) I注册事件C1(nEvent xcc.WM_, pFun interface{}) bool {
	return xc.XWnd_RegEventC1(w.I句柄, nEvent, pFun)
}

// 窗口_移除事件C.
//
// nEvent: 事件类型: xcc.WM_, xcc.XWM_.
//
// pFun: 事件函数.
func (w *windowBase) I移除事件C(nEvent xcc.WM_, pFun interface{}) bool {
	return xc.XWnd_RemoveEventC(w.I句柄, nEvent, pFun)
}

// 窗口_添加子对象.
//
// hChild: 要添加的对象句柄.
func (w *windowBase) I添加子对象(hChild int) bool {
	return xc.XWnd_AddChild(w.I句柄, hChild)
}

// 窗口_插入子对象.
//
// hChild: 要插入的对象句柄.
//
// index: 插入位置索引.
func (w *windowBase) I插入子对象(hChild int, index int) bool {
	return xc.XWnd_InsertChild(w.I句柄, hChild, index)
}

// 窗口_取HWND.
func (w *windowBase) I取HWND() int {
	return xc.XWnd_GetHWND(w.I句柄)
}

// 窗口_重绘.
//
// bImmediate: 是否立即重绘, 默认为否.
func (w *windowBase) I重绘(bImmediate bool) int {
	return xc.XWnd_Redraw(w.I句柄, bImmediate)
}

// 窗口_重绘指定区域.
//
// pRect: 需要重绘的区域坐标.
//
// bImmediate: TRUE立即重绘, FALSE放入消息队列延迟重绘.
func (w *windowBase) I重绘指定区域(pRect *xc.RECT, bImmediate bool) int {
	return xc.XWnd_RedrawRect(w.I句柄, pRect, bImmediate)
}

// 窗口_置坐标.
//
// pRect: 坐标.
func (w *windowBase) I置坐标(pRect *xc.RECT) bool {
	return xc.XWnd_SetRect(w.I句柄, pRect)
}

// 窗口_置焦点.
//
// hFocusEle: 将获得焦点的元素.
func (w *windowBase) I置焦点(hFocusEle int) bool {
	return xc.XWnd_SetFocusEle(w.I句柄, hFocusEle)
}

// 窗口_取焦点.
func (w *windowBase) I取焦点() int {
	return xc.XWnd_GetFocusEle(w.I句柄)
}

// 窗口_取鼠标停留元素.
func (w *windowBase) I取鼠标停留元素() int {
	return xc.XWnd_GetStayEle(w.I句柄)
}

// 窗口_绘制, 在自绘事件函数中,用户手动调用绘制窗口, 以便控制绘制顺序.
//
// hDraw: 图形绘制句柄.
func (w *windowBase) I绘制(hDraw int) int {
	return xc.XWnd_DrawWindow(w.I句柄, hDraw)
}

// 窗口_居中.
func (w *windowBase) I居中() int {
	return xc.XWnd_Center(w.I句柄)
}

// 窗口_居中扩展.
//
// width: 窗口宽度.
//
// height: 窗口高度.
func (w *windowBase) I居中扩展(width, height int) int {
	return xc.XWnd_CenterEx(w.I句柄, width, height)
}

// 窗口_置光标.
//
// hCursor: 鼠标光标句柄.
func (w *windowBase) I置光标(hCursor int) int {
	return xc.XWnd_SetCursor(w.I句柄, hCursor)
}

// 窗口_取光标.
func (w *windowBase) I取光标() int {
	return xc.XWnd_GetCursor(w.I句柄)
}

// 窗口_启用拖动边框.
//
// bEnable: 是否启用.
func (w *windowBase) I启用拖动边框(bEnable bool) int {
	return xc.XWnd_EnableDragBorder(w.I句柄, bEnable)
}

// 窗口_启用拖动窗口.
//
// bEnable: 是否启用.
func (w *windowBase) I启用拖动窗口(bEnable bool) int {
	return xc.XWnd_EnableDragWindow(w.I句柄, bEnable)
}

// 窗口_启用拖动标题栏.
//
// bEnable: 是否启用.
func (w *windowBase) I启用拖动标题栏(bEnable bool) int {
	return xc.XWnd_EnableDragCaption(w.I句柄, bEnable)
}

// 窗口_启用绘制背景.
//
// bEnable: 是否启用.
func (w *windowBase) I启用绘制背景(bEnable bool) int {
	return xc.XWnd_EnableDrawBk(w.I句柄, bEnable)
}

// 窗口_启用自动焦点.
//
// bEnable: 是否启用.
func (w *windowBase) I启用自动焦点(bEnable bool) int {
	return xc.XWnd_EnableAutoFocus(w.I句柄, bEnable)
}

// 窗口_启用允许最大化.
//
// bEnable: 是否启用.
func (w *windowBase) I启用允许最大化(bEnable bool) int {
	return xc.XWnd_EnableMaxWindow(w.I句柄, bEnable)
}

// 窗口_启用限制窗口大小.
//
// bEnable: 是否启用.
func (w *windowBase) I启用限制窗口大小(bEnable bool) int {
	return xc.XWnd_EnablemLimitWindowSize(w.I句柄, bEnable)
}

// 窗口_启用布局.
//
// bEnable: 是否启用.
func (w *windowBase) I启用布局(bEnable bool) int {
	return xc.XWnd_EnableLayout(w.I句柄, bEnable)
}

// 窗口_启用布局覆盖边框.
//
// bEnable: 是否启用.
func (w *windowBase) I启用布局覆盖边框(bEnable bool) int {
	return xc.XWnd_EnableLayoutOverlayBorder(w.I句柄, bEnable)
}

// 窗口_显示布局边界.
//
// bEnable: 是否启用.
func (w *windowBase) I显示布局边界(bEnable bool) int {
	return xc.XWnd_ShowLayoutFrame(w.I句柄, bEnable)
}

// 窗口_判断启用布局.
func (w *windowBase) I判断启用布局() bool {
	return xc.XWnd_IsEnableLayout(w.I句柄)
}

// 窗口_是否最大化.
func (w *windowBase) I是否最大化() bool {
	return xc.XWnd_IsMaxWindow(w.I句柄)
}

// 窗口_置鼠标捕获元素.
//
// hEle: 元素句柄.
func (w *windowBase) I置鼠标捕获元素(hEle int) int {
	return xc.XWnd_SetCaptureEle(w.I句柄, hEle)
}

// 窗口_取鼠标捕获元素.
func (w *windowBase) I取鼠标捕获元素() int {
	return xc.XWnd_GetCaptureEle(w.I句柄)
}

// 窗口_取绘制矩形.
//
// pRcPaint: 重绘区域坐标.
func (w *windowBase) I取绘制矩形(pRcPaint *xc.RECT) int {
	return xc.XWnd_GetDrawRect(w.I句柄, pRcPaint)
}

// 窗口_置系统光标.
//
// hCursor: 光标句柄.
func (w *windowBase) I置系统光标(hCursor int) int {
	return xc.XWnd_SetCursorSys(w.I句柄, hCursor)
}

// 窗口_置字体.
//
// hFontx: 炫彩字体句柄.
func (w *windowBase) I置字体(hFontx int) int {
	return xc.XWnd_SetFont(w.I句柄, hFontx)
}

// 窗口_置文本颜色.
//
// color: ABGR 颜色值.
func (w *windowBase) I置文本颜色(color int) int {
	return xc.XWnd_SetTextColor(w.I句柄, color)
}

// 窗口_取文本颜色, 返回ABGR 颜色.
func (w *windowBase) I取文本颜色() int {
	return xc.XWnd_GetTextColor(w.I句柄)
}

// 窗口_取文本颜色扩展, 返回ABGR 颜色.
func (w *windowBase) I取文本颜色扩展() int {
	return xc.XWnd_GetTextColorEx(w.I句柄)
}

// 窗口_置ID.
//
// nID: ID值.
func (w *windowBase) I置ID(nID int) int {
	return xc.XWnd_SetID(w.I句柄, nID)
}

// 窗口_取ID.
func (w *windowBase) I取ID() int {
	return xc.XWnd_GetID(w.I句柄)
}

// 窗口_置名称.
//
// pName: name值, 字符串.
func (w *windowBase) I置名称(pName string) int {
	return xc.XWnd_SetName(w.I句柄, pName)
}

// 窗口_取名称.
func (w *windowBase) I取名称() string {
	return xc.XWnd_GetName(w.I句柄)
}

// 窗口_置边大小.
//
// left: 窗口左边大小.
//
// top: 窗口上边大小.
//
// right: 窗口右边大小.
//
// bottom: 窗口底部大小.
func (w *windowBase) I置边大小(left, top, right, bottom int) int {
	return xc.XWnd_SetBorderSize(w.I句柄, left, top, right, bottom)
}

// 窗口_取边大小.
func (w *windowBase) I取边大小(pBorder *xc.RECT) int {
	return xc.XWnd_GetBorderSize(w.I句柄, pBorder)
}

// 窗口_置布局内填充大小.
//
// left: 左边大小.
//
// top: 上边大小.
//
// right: 右边大小.
//
// bottom: 下边大小.
func (w *windowBase) I置布局内填充大小(left, top, right, bottom int) int {
	return xc.XWnd_SetPadding(w.I句柄, left, top, right, bottom)
}

// 窗口_置拖动边框大小.
//
// left: 窗口左边大小.
//
// top: 窗口上边大小.
//
// right: 窗口右边大小.
//
// bottom: 窗口底边大小.
func (w *windowBase) I置拖动边框大小(left, top, right, bottom int) int {
	return xc.XWnd_SetDragBorderSize(w.I句柄, left, top, right, bottom)
}

// 窗口_取拖动边框大小.
//
// pSize: 拖动边框大小.
func (w *windowBase) I取拖动边框大小(pBorder *xc.RECT) int {
	return xc.XWnd_GetDragBorderSize(w.I句柄, pBorder)
}

// 窗口_置最小大小.
//
// width: 最小宽度.
//
// height: 最小高度.
func (w *windowBase) I置最小大小(width, height int) int {
	return xc.XWnd_SetMinimumSize(w.I句柄, width, height)
}

// 窗口_测试点击子元素.
//
// pPt: 左边点.
func (w *windowBase) I测试点击子元素(pPt *xc.POINT) int {
	return xc.XWnd_HitChildEle(w.I句柄, pPt)
}

// 窗口_取子对象数量.
func (w *windowBase) I取子对象数量() int {
	return xc.XWnd_GetChildCount(w.I句柄)
}

// 窗口_取子对象从索引.
//
// index: 元素索引.
func (w *windowBase) I取子对象从索引(index int) int {
	return xc.XWnd_GetChildByIndex(w.I句柄, index)
}

// 窗口_取子对象从ID.
//
// nID: 元素ID, ID必须大于0.
func (w *windowBase) I取子对象从ID(nID int) int {
	return xc.XWnd_GetChildByID(w.I句柄, nID)
}

// 窗口_取子对象.
//
// nID: 对象ID,ID必须大于0.
func (w *windowBase) I取子对象(nID int) int {
	return xc.XWnd_GetChild(w.I句柄, nID)
}

// 窗口_关闭.
func (w *windowBase) I关闭() int {
	return xc.XWnd_CloseWindow(w.I句柄)
}

// 窗口_调整布局.
func (w *windowBase) I调整布局() int {
	return xc.XWnd_AdjustLayout(w.I句柄)
}

// 窗口_调整布局扩展.
//
// nFlags: 调整布局标识位: xcc.AdjustLayout_.
func (w *windowBase) I调整布局扩展(nFlags xcc.AdjustLayout_) int {
	return xc.XWnd_AdjustLayoutEx(w.I句柄, nFlags)
}

// 窗口_创建插入符.
//
// hEle: 元素句柄.
//
// x: x坐标.
//
// y: y坐标.
//
// width: 宽度.
//
// height: 高度.
func (w *windowBase) I创建插入符(hEle, x, y, width, height int) int {
	return xc.XWnd_CreateCaret(w.I句柄, hEle, x, y, width, height)
}

// 窗口_置插入符位置, 设置插入符位置.
//
// x: x坐标.
//
// y: y坐标.
//
// width: 宽度.
//
// height: 高度.
//
// bUpdate: 是否立即更新UI.
func (w *windowBase) I置插入符位置(x, y, width, height int, bUpdate bool) int {
	return xc.XWnd_SetCaretPos(w.I句柄, x, y, width, height, bUpdate)
}

// 窗口_置插入符颜色.
//
// color: 颜色值, ABGR 颜色.
func (w *windowBase) I置插入符颜色(color int) int {
	return xc.XWnd_SetCaretColor(w.I句柄, color)
}

// 窗口_显示插入符.
//
// bShow: 是否显示.
func (w *windowBase) I显示插入符(bShow bool) int {
	return xc.XWnd_ShowCaret(w.I句柄, bShow)
}

// 窗口_销毁插入符.
func (w *windowBase) I销毁插入符() int {
	return xc.XWnd_DestroyCaret(w.I句柄)
}

// 窗口_取插入符元素.
func (w *windowBase) I取插入符元素() int {
	return xc.XWnd_GetCaretHELE(w.I句柄)
}

// 窗口_取客户区坐标.
//
// pRect: 坐标.
func (w *windowBase) I取客户区坐标(pRect *xc.RECT) int {
	return xc.XWnd_GetClientRect(w.I句柄, pRect)
}

// 窗口_取Body坐标.
//
// pRect: 坐标.
func (w *windowBase) I取Body坐标(pRect *xc.RECT) int {
	return xc.XWnd_GetBodyRect(w.I句柄, pRect)
}

// 窗口_取布局坐标.
//
// pRect: 接收返回坐标.
func (w *windowBase) I取布局坐标(pRect *xc.RECT) int {
	return xc.XWnd_GetLayoutRect(w.I句柄, pRect)
}

// 窗口_移动窗口.
//
// x: X坐标.
//
// y: Y坐标.
func (w *windowBase) I移动窗口(x, y int) int {
	return xc.XWnd_SetPosition(w.I句柄, x, y)
}

// 窗口_取坐标.
//
// pRect: 坐标.
func (w *windowBase) I取坐标(pRect *xc.RECT) int {
	return xc.XWnd_GetRect(w.I句柄, pRect)
}

// 窗口_最大化.
//
// bMaximize: 是否最大化.
func (w *windowBase) I最大化(bMaximize bool) int {
	return xc.XWnd_MaxWindow(w.I句柄, bMaximize)
}

// 窗口_置定时器.
//
// nIDEvent: 定时器ID.
//
// uElapse: 间隔值, 单位毫秒.
func (w *windowBase) I置定时器(nIDEvent, uElapse int) int {
	return xc.XWnd_SetTimer(w.I句柄, nIDEvent, uElapse)
}

// 窗口_关闭定时器.
//
// nIDEvent: 定时器ID.
func (w *windowBase) I关闭定时器(nIDEvent int) int {
	return xc.XWnd_KillTimer(w.I句柄, nIDEvent)
}

// 窗口_置炫彩定时器.
//
// nIDEvent: 定时器ID.
//
// uElapse: 间隔值, 单位毫秒.
func (w *windowBase) I置炫彩定时器(nIDEvent, uElapse int) int {
	return xc.XWnd_SetXCTimer(w.I句柄, nIDEvent, uElapse)
}

// 窗口_关闭炫彩定时器.
//
// nIDEvent: 定时器ID.
func (w *windowBase) I关闭炫彩定时器(nIDEvent int) int {
	return xc.XWnd_KillXCTimer(w.I句柄, nIDEvent)
}

// 窗口_取背景管理器.
func (w *windowBase) I取背景管理器() int {
	return xc.XWnd_GetBkManager(w.I句柄)
}

// 窗口_取背景管理器扩展.
func (w *windowBase) I取背景管理器扩展() int {
	return xc.XWnd_GetBkManagerEx(w.I句柄)
}

// 窗口_置背景管理器.
//
// hBkInfoM: 背景管理器.
func (w *windowBase) I置背景管理器(hBkInfoM int) int {
	return xc.XWnd_SetBkMagager(w.I句柄, hBkInfoM)
}

// 窗口_置透明类型.
//
// nType: 窗口透明类型: xcc.I常量_窗口透明_.
func (w *windowBase) I置透明类型(nType xcc.I常量_窗口透明_) int {
	return xc.XWnd_SetTransparentType(w.I句柄, nType)
}

// 窗口_置透明度.
//
// alpha: 窗口透明度, 范围0-255之间, 0透明, 255不透明.
func (w *windowBase) I置透明度(alpha byte) int {
	return xc.XWnd_SetTransparentAlpha(w.I句柄, alpha)
}

// 窗口_置透明色.
//
// color: 窗口透明色, ABGR 颜色.
func (w *windowBase) I置透明色(color int) int {
	return xc.XWnd_SetTransparentColor(w.I句柄, color)
}

// 窗口_置阴影信息.
//
// nSize: 阴影大小.
//
// nDepth: 阴影深度, 0-255.
//
// nAngeleSize: 圆角阴影内收大小.
//
// bRightAngle: 是否强制直角.
//
// color: 阴影颜色, ABGR 颜色.
func (w *windowBase) I置阴影信息(nSize int, nDepth byte, nAngeleSize int, bRightAngle bool, color int) int {
	return xc.XWnd_SetShadowInfo(w.I句柄, nSize, nDepth, nAngeleSize, bRightAngle, color)
}

// 窗口_取阴影信息.
//
// pnSize: 阴影大小.
//
// pnDepth: 阴影深度.
//
// pnAngeleSize: 圆角阴影内收大小 .
//
// pbRightAngle: 是否强制直角.
//
// pColor: 阴影颜色, ABGR 颜色.
func (w *windowBase) I取阴影信息(nSize *int, nDepth *byte, nAngeleSize *int, bRightAngle *bool, color *int) int {
	return xc.XWnd_GetShadowInfo(w.I句柄, nSize, nDepth, nAngeleSize, bRightAngle, color)
}

// 窗口_取透明类型, 返回: xcc.I常量_窗口透明_.
func (w *windowBase) I取透明类型() xcc.I常量_窗口透明_ {
	return xc.XWnd_GetTransparentType(w.I句柄)
}

// 窗口_启用拖放文件.
//
// bEnable: 是否启用.
func (w *windowBase) I启用拖放文件(bEnable bool) bool {
	return xc.XWnd_EnableDragFiles(w.I句柄, bEnable)
}

// 窗口_显示 显示隐藏窗口.
//
// bShow: 是否显示.
func (w *windowBase) I显示隐藏(bShow bool) int {
	return xc.XWnd_Show(w.I句柄, bShow)
}

// 窗口_取插入符信息, 获取插入符信息, 返回: 插入符元素句柄.
//
// hWindow: 窗口句柄.
//
// pX: 接收返回x坐标.
//
// pY: 接收返回y坐标.
//
// pWidth: 接收返回宽度.
//
// pHeight: 接收返回高度.
func (w *windowBase) I取插入符信息(pX, pY, pWidth, pHeight *int) int {
	return xc.XWnd_GetCaretInfo(w.I句柄, pX, pY, pWidth, pHeight)
}

// 窗口_置图标.
//
// hImage: 图标句柄.
func (w *windowBase) I置图标(hImage int) int {
	return xc.XWnd_SetIcon(w.I句柄, hImage)
}

// 窗口_置标题.
//
// pTitle: 标题文本.
func (w *windowBase) I置标题(pTitle string) int {
	return xc.XWnd_SetTitle(w.I句柄, pTitle)
}

// 窗口_置标题颜色.
//
// color: ABGR 颜色.
func (w *windowBase) I置标题颜色(color int) int {
	return xc.XWnd_SetTitleColor(w.I句柄, color)
}

// 窗口_取控制按钮, 返回按钮句柄.
//
// nFlag: 可用值: I常量_窗口样式_按钮_最小化 , I常量_窗口样式_按钮_最大化 , I常量_窗口样式_按钮_关闭 .
func (w *windowBase) I取控制按钮(nFlag int) int {
	return xc.XWnd_GetButton(w.I句柄, nFlag)
}

// 窗口_取图标, 返回图标句柄.
func (w *windowBase) I取图标() int {
	return xc.XWnd_GetIcon(w.I句柄)
}

// 窗口_取标题, 返回标题文本.
func (w *windowBase) I取标题() string {
	return xc.XWnd_GetTitle(w.I句柄)
}

// 窗口_取标题颜色, 返回ABGR 颜色.
func (w *windowBase) I取标题颜色() int {
	return xc.XWnd_GetTitleColor(w.I句柄)
}

// 窗口_添加背景边框.
//
// nState: 组合状态.
//
// color: ABGR 颜色.
//
// width: 线宽.
func (w *windowBase) I添加背景边框(nState xcc.Window_State_Flag_, color int, width int) int {
	return xc.XWnd_AddBkBorder(w.I句柄, nState, color, width)
}

// 窗口_添加背景填充.
//
// nState: 组合状态.
//
// color: ABGR 颜色.
func (w *windowBase) I添加背景填充(nState xcc.Window_State_Flag_, color int) int {
	return xc.XWnd_AddBkFill(w.I句柄, nState, color)
}

// 窗口_添加背景图片.
//
// nState: 组合状态.
//
// hImage: 图片句柄.
func (w *windowBase) I添加背景图片(nState xcc.Window_State_Flag_, hImage int) int {
	return xc.XWnd_AddBkImage(w.I句柄, nState, hImage)
}

// 窗口_取背景对象数量.
func (w *windowBase) I取背景对象数量() int {
	return xc.XWnd_GetBkInfoCount(w.I句柄)
}

// 窗口_清空背景对象.
func (w *windowBase) I清空背景对象() int {
	return xc.XWnd_ClearBkInfo(w.I句柄)
}

// 通知消息_窗口中弹出, 使用基础元素作为面板, 弹出一个通知消息, 返回元素句柄, 通过此句柄可对其操作.
//
// position: 位置, I常量_位置标识_.
//
// pTitle: 标题.
//
// pText: 内容.
//
// hIcon: 图标.
//
// skin: 外观类型, I常量_通知消息外观_.
func (w *windowBase) I通知消息_窗口中弹出(position xcc.I常量_位置标识_, pTitle, pText string, hIcon int, skin xcc.I常量_通知消息外观_) int {
	return xc.XNotifyMsg_WindowPopup(w.I句柄, position, pTitle, pText, hIcon, skin)
}

// 通知消息_窗口中弹出扩展, 使用基础元素作为面板, 弹出一个通知消息, 返回元素句柄, 通过此句柄可对其操作.
//
// position: 位置, I常量_位置标识_.
//
// pTitle: 标题.
//
// pText: 内容.
//
// hIcon: 图标.
//
// skin: 外观类型, I常量_通知消息外观_.
//
// bBtnClose: 是否启用关闭按钮.
//
// bAutoClose: 是否自动关闭.
//
// nWidth: 自定义宽度, -1(使用默认值).
//
// nHeight: 自定义高度, -1(使用默认值).
func (w *windowBase) I通知消息_窗口中弹出扩展(position xcc.I常量_位置标识_, pTitle, pText string, hIcon int, skin xcc.I常量_通知消息外观_, bBtnClose, bAutoClose bool, nWidth, nHeight int) int {
	return xc.XNotifyMsg_WindowPopupEx(w.I句柄, position, pTitle, pText, hIcon, skin, bBtnClose, bAutoClose, nWidth, nHeight)
}

// 通知消息_置持续时间.
//
// duration: 持续时间.
func (w *windowBase) I通知消息_置持续时间(duration int) int {
	return xc.XNotifyMsg_SetDuration(w.I句柄, duration)
}

// 通知消息_置父边距 设置通知消息与父对象的四边间隔.
//
// left: 左侧间隔, 未实现, 预留功能.
//
// top: 顶部间隔.
//
// right: 右侧间隔.
//
// bottom: 底部间隔, 未实现, 预留功能.
func (w *windowBase) I通知消息_置父边距(left, top, right, bottom int) int {
	return xc.XNotifyMsg_SetParentMargin(w.I句柄, left, top, right, bottom)
}

// 通知消息_置标题高度.
//
// nHeight: 高度.
func (w *windowBase) I通知消息_置标题高度(nHeight int) int {
	return xc.XNotifyMsg_SetCaptionHeight(w.I句柄, nHeight)
}

// 通知消息_置宽度, 设置默认宽度.
//
// nWidth: 宽度.
func (w *windowBase) I通知消息_置宽度(nWidth int) int {
	return xc.XNotifyMsg_SetWidth(w.I句柄, nWidth)
}

// 通知消息_置间隔.
//
// nSpace: 间隔大小.
func (w *windowBase) I通知消息_置间隔(nSpace int) int {
	return xc.XNotifyMsg_SetSpace(w.I句柄, nSpace)
}

// 通知消息_置边大小, 设置通知消息面板边大小.
//
// left: 左边.
//
// top: 顶边.
//
// right: 右边.
//
// bottom: 底边.
func (w *windowBase) I通知消息_置边大小(left, top, right, bottom int) int {
	return xc.XNotifyMsg_SetBorderSize(w.I句柄, left, top, right, bottom)
}

// 窗口_置背景, 返回设置的背景对象数量.
//
// pText: 背景内容字符串.
func (w *windowBase) I置背景(pText string) int {
	return xc.XWnd_SetBkInfo(w.I句柄, pText)
}

// 窗口_是否可拖动标题栏.
func (w *windowBase) I是否可拖动标题栏() bool {
	return xc.XWnd_IsDragCaption(w.I句柄)
}

// 窗口_是否可拖动窗口.
func (w *windowBase) I是否可拖动窗口() bool {
	return xc.XWnd_IsDragWindow(w.I句柄)
}

// 窗口_是否可拖动边框.
func (w *windowBase) I是否可拖动边框() bool {
	return xc.XWnd_IsDragBorder(w.I句柄)
}

// 窗口_置标题外间距, 设置标题内容(图标, 标题, 控制按钮)外间距.
//
// left: 左边间距.
//
// top: 上边间距.
//
// right: 右边间距.
//
// bottom: 下边间距.
func (w *windowBase) I置标题外间距(left int, top int, right int, bottom int) int {
	return xc.XWnd_SetCaptionMargin(w.I句柄, left, top, right, bottom)
}

// I窗口_置顶Ex 窗口_置顶Ex.
//	@param b 是否置顶.
//	@return bool
//
func (w *windowBase) I窗口_置顶Ex(b bool) bool {
	return wnd.SetTop(w.I取HWND(), b)
}

/*
下面都是事件
*/

type XWM_WINDPROC func(message uint, wParam int, lParam int, pbHandled *bool) int               // 窗口消息过程.
type XWM_WINDPROC1 func(hWindow int, message uint, wParam int, lParam int, pbHandled *bool) int // 窗口消息过程.
type XWM_XC_TIMER func(nTimerID uint, pbHandled *bool) int                                      // 炫彩定时器, 非系统定时器, 注册消息XWM_TIMER接收.
type XWM_XC_TIMER1 func(hWindow int, nTimerID uint, pbHandled *bool) int                        // 炫彩定时器, 非系统定时器, 注册消息XWM_TIMER接收.
type XWM_FLOAT_PANE func(hFloatWnd int, hPane int, pbHandled *bool) int                         // 浮动窗格.
type XWM_FLOAT_PANE1 func(hWindow int, hFloatWnd int, hPane int, pbHandled *bool) int           // 浮动窗格.
type XWM_PAINT_END func(hDraw int, pbHandled *bool) int                                         // 窗口绘制完成消息.
type XWM_PAINT_END1 func(hWindow int, hDraw int, pbHandled *bool) int                           // 窗口绘制完成消息.
type XWM_PAINT_DISPLAY func(pbHandled *bool) int                                                // 窗口绘制完成并且已经显示到屏幕.
type XWM_PAINT_DISPLAY1 func(hWindow int, pbHandled *bool) int                                  // 窗口绘制完成并且已经显示到屏幕.
type XWM_DOCK_POPUP func(hWindowDock, hPane int, pbHandled *bool) int                           // 框架窗口码头弹出窗格, 当用户点击码头上的按钮时, 显示对应的窗格, 当失去焦点时自动隐藏窗格.
type XWM_DOCK_POPUP1 func(hWindow int, hWindowDock, hPane int, pbHandled *bool) int             // 框架窗口码头弹出窗格, 当用户点击码头上的按钮时, 显示对应的窗格, 当失去焦点时自动隐藏窗格.
// 浮动窗口拖动, 用户拖动浮动窗口移动, 显示停靠提示.
//
// hFloatWnd: 拖动的浮动窗口句柄.
//
// hArray: HWINDOW array[6], 窗格停靠提示窗口句柄数组, 有6个成员, 分别为:[0]中间十字, [1]左侧, [2]顶部, [3]右侧, [4]底部, [5]停靠位置预览.
type XWM_FLOATWND_DRAG func(hFloatWnd int, hArray *[6]int, pbHandled *bool) int

// 浮动窗口拖动, 用户拖动浮动窗口移动, 显示停靠提示.
//
// hWindow: 传入的窗口资源句柄.
//
// hFloatWnd: 拖动的浮动窗口句柄.
//
// hArray: HWINDOW array[6], 窗格停靠提示窗口句柄数组, 有6个成员, 分别为:[0]中间十字, [1]左侧, [2]顶部, [3]右侧, [4]底部, [5]停靠位置预览.
type XWM_FLOATWND_DRAG1 func(hWindow int, hFloatWnd int, hArray *[6]int, pbHandled *bool) int

type WM_PAINT func(hDraw int, pbHandled *bool) int                                        // 窗口绘制消息.
type WM_PAINT1 func(hWindow int, hDraw int, pbHandled *bool) int                          // 窗口绘制消息.
type WM_CLOSE func(pbHandled *bool) int                                                   // 窗口关闭消息.
type WM_CLOSE1 func(hWindow int, pbHandled *bool) int                                     // 窗口关闭消息.
type WM_DESTROY func(pbHandled *bool) int                                                 // 窗口销毁消息.
type WM_DESTROY1 func(hWindow int, pbHandled *bool) int                                   // 窗口销毁消息.
type WM_NCDESTROY func(pbHandled *bool) int                                               // 窗口非客户区销毁消息.
type WM_NCDESTROY1 func(hWindow int, pbHandled *bool) int                                 // 窗口非客户区销毁消息.
type WM_MOUSEMOVE func(nFlags uint, pPt *xc.POINT, pbHandled *bool) int                   // 窗口鼠标移动消息.
type WM_MOUSEMOVE1 func(hWindow int, nFlags uint, pPt *xc.POINT, pbHandled *bool) int     // 窗口鼠标移动消息.
type WM_LBUTTONDOWN func(nFlags uint, pPt *xc.POINT, pbHandled *bool) int                 // 窗口鼠标左键按下消息.
type WM_LBUTTONDOWN1 func(hWindow int, nFlags uint, pPt *xc.POINT, pbHandled *bool) int   // 窗口鼠标左键按下消息.
type WM_LBUTTONUP func(nFlags uint, pPt *xc.POINT, pbHandled *bool) int                   // 窗口鼠标左键弹起消息.
type WM_LBUTTONUP1 func(hWindow int, nFlags uint, pPt *xc.POINT, pbHandled *bool) int     // 窗口鼠标左键弹起消息.
type WM_RBUTTONDOWN func(nFlags uint, pPt *xc.POINT, pbHandled *bool) int                 // 窗口鼠标右键按下消息.
type WM_RBUTTONDOWN1 func(hWindow int, nFlags uint, pPt *xc.POINT, pbHandled *bool) int   // 窗口鼠标右键按下消息.
type WM_RBUTTONUP func(nFlags uint, pPt *xc.POINT, pbHandled *bool) int                   // 窗口鼠标右键弹起消息.
type WM_RBUTTONUP1 func(hWindow int, nFlags uint, pPt *xc.POINT, pbHandled *bool) int     // 窗口鼠标右键弹起消息.
type WM_LBUTTONDBLCLK func(nFlags uint, pPt *xc.POINT, pbHandled *bool) int               // 窗口鼠标左键双击消息.
type WM_LBUTTONDBLCLK1 func(hWindow int, nFlags uint, pPt *xc.POINT, pbHandled *bool) int // 窗口鼠标左键双击消息.
type WM_RBUTTONDBLCLK func(nFlags uint, pPt *xc.POINT, pbHandled *bool) int               // 窗口鼠标右键双击消息.
type WM_RBUTTONDBLCLK1 func(hWindow int, nFlags uint, pPt *xc.POINT, pbHandled *bool) int // 窗口鼠标右键双击消息.
type WM_MOUSEWHEEL func(nFlags uint, pPt *xc.POINT, pbHandled *bool) int                  // 窗口鼠标滚轮滚动消息.
type WM_MOUSEWHEEL1 func(hWindow int, nFlags uint, pPt *xc.POINT, pbHandled *bool) int    // 窗口鼠标滚轮滚动消息.
type WM_EXITSIZEMOVE func(pbHandled *bool) int                                            // 窗口退出移动或调整大小模式循环改，详情参见MSDN.
type WM_EXITSIZEMOVE1 func(hWindow int, pbHandled *bool) int                              // 窗口退出移动或调整大小模式循环改，详情参见MSDN.
type WM_MOUSEHOVER func(nFlags uint, pPt *xc.POINT, pbHandled *bool) int                  // 窗口鼠标进入消息.
type WM_MOUSEHOVER1 func(hWindow int, nFlags uint, pPt *xc.POINT, pbHandled *bool) int    // 窗口鼠标进入消息.
type WM_MOUSELEAVE func(pbHandled *bool) int                                              // 窗口鼠标离开消息.
type WM_MOUSELEAVE1 func(hWindow int, pbHandled *bool) int                                // 窗口鼠标离开消息.
type WM_SIZE func(nFlags uint, pPt *xc.SIZE, pbHandled *bool) int                         // 窗口大小改变消息.
type WM_SIZE1 func(hWindow int, nFlags uint, pPt *xc.SIZE, pbHandled *bool) int           // 窗口大小改变消息.
type WM_TIMER func(nIDEvent uint, pbHandled *bool) int                                    // 窗口定时器消息.
type WM_TIMER1 func(hWindow int, nIDEvent uint, pbHandled *bool) int                      // 窗口定时器消息.
type WM_SETFOCUS func(pbHandled *bool) int                                                // 窗口获得焦点.
type WM_SETFOCUS1 func(hWindow int, pbHandled *bool) int                                  // 窗口获得焦点.
type WM_KILLFOCUS func(pbHandled *bool) int                                               // 窗口失去焦点.
type WM_KILLFOCUS1 func(hWindow int, pbHandled *bool) int                                 // 窗口失去焦点.
type WM_KEYDOWN func(wParam int, lParam int, pbHandled *bool) int                         // 窗口键盘按键消息.
type WM_KEYDOWN1 func(hWindow int, wParam int, lParam int, pbHandled *bool) int           // 窗口键盘按键消息.
type WM_CAPTURECHANGED func(hWnd int, pbHandled *bool) int                                // 窗口鼠标捕获改变消息.
type WM_CAPTURECHANGED1 func(hWindow int, hWnd int, pbHandled *bool) int                  // 窗口鼠标捕获改变消息.
type WM_SETCURSOR func(wParam int, lParam int, pbHandled *bool) int                       // 窗口设置鼠标光标.
type WM_SETCURSOR1 func(hWindow int, wParam int, lParam int, pbHandled *bool) int         // 窗口设置鼠标光标.
type WM_CHAR func(wParam int, lParam int, pbHandled *bool) int                            // 窗口字符消息.
type WM_CHAR1 func(hWindow int, wParam int, lParam int, pbHandled *bool) int              // 窗口字符消息.
type WM_DROPFILES func(hDropInfo int, pbHandled *bool) int                                // 拖动文件到窗口.
type WM_DROPFILES1 func(hWindow int, hDropInfo int, pbHandled *bool) int                  // 拖动文件到窗口.

// 窗口消息过程.
func (w *windowBase) I事件_消息过程(pFun XWM_WINDPROC) bool {
	return xc.XWnd_RegEventC(w.I句柄, xcc.XWM_WINDPROC, pFun)
}

// 窗口消息过程.
func (w *windowBase) I事件_消息过程1(pFun XWM_WINDPROC1) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.XWM_WINDPROC, pFun)
}

// 炫彩定时器, 非系统定时器, 注册消息XWM_TIMER接收.
func (w *windowBase) I事件_定时器(pFun XWM_XC_TIMER) bool {
	return xc.XWnd_RegEventC(w.I句柄, xcc.XWM_XC_TIMER, pFun)
}

// 炫彩定时器, 非系统定时器, 注册消息XWM_TIMER接收.
func (w *windowBase) I事件_定时器1(pFun XWM_XC_TIMER1) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.XWM_XC_TIMER, pFun)
}

// I事件_浮动窗格.
func (w *windowBase) I事件_浮动窗格(pFun XWM_FLOAT_PANE) bool {
	return xc.XWnd_RegEventC(w.I句柄, xcc.XWM_FLOAT_PANE, pFun)
}

// I事件_浮动窗格.
func (w *windowBase) I事件_浮动窗格1(pFun XWM_FLOAT_PANE1) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.XWM_FLOAT_PANE, pFun)
}

// I事件_窗口绘制完成 消息.
func (w *windowBase) I事件_窗口绘制完成(pFun XWM_PAINT_END) bool {
	return xc.XWnd_RegEventC(w.I句柄, xcc.XWM_PAINT_END, pFun)
}

// I事件_窗口绘制完成消息.
func (w *windowBase) I事件_窗口绘制完成1(pFun XWM_PAINT_END1) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.XWM_PAINT_END, pFun)
}

// I事件_窗口绘制完成并已显示.
func (w *windowBase) I事件_窗口绘制完成并已显示(pFun XWM_PAINT_DISPLAY) bool {
	return xc.XWnd_RegEventC(w.I句柄, xcc.XWM_PAINT_DISPLAY, pFun)
}

// I事件_窗口绘制完成并已显示.
func (w *windowBase) I事件_窗口绘制完成并已显示1(pFun XWM_PAINT_DISPLAY1) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.XWM_PAINT_DISPLAY, pFun)
}

// I事件_框架窗口码头弹出窗格 , 当用户点击码头上的按钮时, 显示对应的窗格, 当失去焦点时自动隐藏窗格.
func (w *windowBase) I事件_框架窗口码头弹出窗格(pFun XWM_DOCK_POPUP) bool {
	return xc.XWnd_RegEventC(w.I句柄, xcc.XWM_DOCK_POPUP, pFun)
}

// I事件_框架窗口码头弹出窗格 , 当用户点击码头上的按钮时, 显示对应的窗格, 当失去焦点时自动隐藏窗格.
func (w *windowBase) I事件_框架窗口码头弹出窗格1(pFun XWM_DOCK_POPUP1) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.XWM_DOCK_POPUP, pFun)
}

// I事件_浮动窗口拖动1 , 用户拖动浮动窗口移动, 显示停靠提示.
//
// hFloatWnd: 拖动的浮动窗口句柄.
//
// hArray: HWINDOW array[6], 窗格停靠提示窗口句柄数组, 有6个成员, 分别为:[0]中间十字, [1]左侧, [2]顶部, [3]右侧, [4]底部, [5]停靠位置预览.
func (w *windowBase) I事件_浮动窗口拖动(pFun XWM_FLOATWND_DRAG) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.XWM_FLOATWND_DRAG, pFun)
}

// I事件_浮动窗口拖动1 , 用户拖动浮动窗口移动, 显示停靠提示.
//
// hWindow: 传入的窗口资源句柄.
//
// hFloatWnd: 拖动的浮动窗口句柄.
//
// hArray: HWINDOW array[6], 窗格停靠提示窗口句柄数组, 有6个成员, 分别为:[0]中间十字, [1]左侧, [2]顶部, [3]右侧, [4]底部, [5]停靠位置预览.
func (w *windowBase) I事件_浮动窗口拖动1(pFun XWM_FLOATWND_DRAG1) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.XWM_FLOATWND_DRAG, pFun)
}

// I事件_窗口绘制   消息 .
func (w *windowBase) I事件_窗口绘制(pFun WM_PAINT) bool {
	return xc.XWnd_RegEventC(w.I句柄, xcc.WM_PAINT, pFun)
}

// I事件_窗口绘制   消息.
func (w *windowBase) Event_PAINT1(pFun WM_PAINT1) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.WM_PAINT, pFun)
}

// I事件_窗口关闭   消息.
func (w *windowBase) I事件_窗口关闭(pFun WM_CLOSE) bool {
	return xc.XWnd_RegEventC(w.I句柄, xcc.WM_CLOSE, pFun)
}

// I事件_窗口关闭   消息.
func (w *windowBase) I事件_窗口关闭1(pFun WM_CLOSE1) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.WM_CLOSE, pFun)
}

// I事件_窗口销毁   消息.
func (w *windowBase) I事件_窗口销毁(pFun WM_DESTROY) bool {
	return xc.XWnd_RegEventC(w.I句柄, xcc.WM_DESTROY, pFun)
}

// I事件_窗口销毁消息.
func (w *windowBase) I事件_窗口销毁1(pFun WM_DESTROY1) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.WM_DESTROY, pFun)
}

// 窗口非客户区销毁   消息.
func (w *windowBase) I事件_窗口非客户区销毁(pFun WM_NCDESTROY) bool {
	return xc.XWnd_RegEventC(w.I句柄, xcc.WM_NCDESTROY, pFun)
}

// I事件_窗口非客户区销毁   消息.
func (w *windowBase) I事件_窗口非客户区销毁1(pFun WM_NCDESTROY1) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.WM_NCDESTROY, pFun)
}

// I事件_窗口鼠标移动   消息.
func (w *windowBase) I事件_窗口鼠标移动(pFun WM_MOUSEMOVE) bool {
	return xc.XWnd_RegEventC(w.I句柄, xcc.WM_MOUSEMOVE, pFun)
}

// I事件_窗口鼠标移动  消息.
func (w *windowBase) I事件_窗口鼠标移动1(pFun WM_MOUSEMOVE1) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.WM_MOUSEMOVE, pFun)
}

// I事件_窗口鼠标左键按下   消息.
func (w *windowBase) I事件_窗口鼠标左键按下(pFun WM_LBUTTONDOWN) bool {
	return xc.XWnd_RegEventC(w.I句柄, xcc.WM_LBUTTONDOWN, pFun)
}

// I事件_窗口鼠标左键按下   消息.
func (w *windowBase) I事件_窗口鼠标左键按下1(pFun WM_LBUTTONDOWN1) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.WM_LBUTTONDOWN, pFun)
}

// I事件_窗口鼠标左键弹起   消息.
func (w *windowBase) I事件_窗口鼠标左键弹起(pFun WM_LBUTTONUP) bool {
	return xc.XWnd_RegEventC(w.I句柄, xcc.WM_LBUTTONUP, pFun)
}

// I事件_窗口鼠标左键弹起   消息.
func (w *windowBase) I事件_窗口鼠标左键弹起1(pFun WM_LBUTTONUP1) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.WM_LBUTTONUP, pFun)
}

// I事件_窗口鼠标右键按下  消息.
func (w *windowBase) I事件_窗口鼠标右键按下(pFun WM_RBUTTONDOWN) bool {
	return xc.XWnd_RegEventC(w.I句柄, xcc.WM_RBUTTONDOWN, pFun)
}

// I事件_窗口鼠标右键按下  消息.
func (w *windowBase) I事件_窗口鼠标右键按下1(pFun WM_RBUTTONDOWN1) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.WM_RBUTTONDOWN, pFun)
}

// I事件_窗口鼠标右键弹起  消息.
func (w *windowBase) I事件_窗口鼠标右键弹起(pFun WM_RBUTTONUP) bool {
	return xc.XWnd_RegEventC(w.I句柄, xcc.WM_RBUTTONUP, pFun)
}

// I事件_窗口鼠标右键弹起   消息.
func (w *windowBase) I事件_窗口鼠标右键弹起1(pFun WM_RBUTTONUP1) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.WM_RBUTTONUP, pFun)
}

// I事件_窗口鼠标左键双击   消息.
func (w *windowBase) I事件_窗口鼠标左键双击(pFun WM_LBUTTONDBLCLK) bool {
	return xc.XWnd_RegEventC(w.I句柄, xcc.WM_LBUTTONDBLCLK, pFun)
}

// I事件_窗口鼠标左键双击   消息.
func (w *windowBase) I事件_窗口鼠标左键双击1(pFun WM_LBUTTONDBLCLK1) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.WM_LBUTTONDBLCLK, pFun)
}

// I事件_窗口鼠标右键双击   消息.
func (w *windowBase) I事件_窗口鼠标右键双击(pFun WM_RBUTTONDBLCLK) bool {
	return xc.XWnd_RegEventC(w.I句柄, xcc.WM_RBUTTONDBLCLK, pFun)
}

// I事件_窗口鼠标右键双击   消息.
func (w *windowBase) I事件_窗口鼠标右键双击1(pFun WM_RBUTTONDBLCLK1) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.WM_RBUTTONDBLCLK, pFun)
}

// I事件_窗口鼠标滚轮滚动   消息.
func (w *windowBase) I事件_窗口鼠标滚轮滚动(pFun WM_MOUSEWHEEL) bool {
	return xc.XWnd_RegEventC(w.I句柄, xcc.WM_MOUSEWHEEL, pFun)
}

// I事件_窗口鼠标滚轮滚动   消息.
func (w *windowBase) I事件_窗口鼠标滚轮滚动1(pFun WM_MOUSEWHEEL1) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.WM_MOUSEWHEEL, pFun)
}

// I事件_窗口退出移动   或调整大小模式循环改，详情参见MSDN.
func (w *windowBase) I事件_窗口退出移动(pFun WM_EXITSIZEMOVE) bool {
	return xc.XWnd_RegEventC(w.I句柄, xcc.WM_EXITSIZEMOVE, pFun)
}

// I事件_窗口退出移动   或调整大小模式循环改，详情参见MSDN.
func (w *windowBase) I事件_窗口退出移动1(pFun WM_EXITSIZEMOVE1) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.WM_EXITSIZEMOVE, pFun)
}

// I事件_窗口鼠标进入  消息.
func (w *windowBase) I事件_窗口鼠标进入(pFun WM_MOUSEHOVER) bool {
	return xc.XWnd_RegEventC(w.I句柄, xcc.WM_MOUSEHOVER, pFun)
}

// I事件_窗口鼠标进入  消息.
func (w *windowBase) I事件_窗口鼠标进入1(pFun WM_MOUSEHOVER1) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.WM_MOUSEHOVER, pFun)
}

// I事件_窗口鼠标离开  消息.
func (w *windowBase) I事件_窗口鼠标离开(pFun WM_MOUSELEAVE) bool {
	return xc.XWnd_RegEventC(w.I句柄, xcc.WM_MOUSELEAVE, pFun)
}

// I事件_窗口鼠标离开  消息.
func (w *windowBase) I事件_窗口鼠标离开1(pFun WM_MOUSELEAVE1) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.WM_MOUSELEAVE, pFun)
}

// I事件_窗口大小改变  消息.
func (w *windowBase) I事件_窗口大小改变(pFun WM_SIZE) bool {
	return xc.XWnd_RegEventC(w.I句柄, xcc.WM_SIZE, pFun)
}

// I事件_窗口大小改变  消息.
func (w *windowBase) I事件_窗口大小改变1(pFun WM_SIZE1) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.WM_SIZE, pFun)
}

// I事件_窗口定时器  消息.
func (w *windowBase) I事件_窗口定时器(pFun WM_TIMER) bool {
	return xc.XWnd_RegEventC(w.I句柄, xcc.WM_TIMER, pFun)
}

// I事件_窗口定时器  消息.
func (w *windowBase) I事件_窗口定时器1(pFun WM_TIMER1) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.WM_TIMER, pFun)
}

// I事件_窗口获得焦点.
func (w *windowBase) I事件_窗口获得焦点(pFun WM_SETFOCUS) bool {
	return xc.XWnd_RegEventC(w.I句柄, xcc.WM_SETFOCUS, pFun)
}

// I事件_窗口获得焦点.
func (w *windowBase) I事件_窗口获得焦点1(pFun WM_SETFOCUS1) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.WM_SETFOCUS, pFun)
}

// I事件_窗口失去焦点.
func (w *windowBase) I事件_窗口失去焦点(pFun WM_KILLFOCUS) bool {
	return xc.XWnd_RegEventC(w.I句柄, xcc.WM_KILLFOCUS, pFun)
}

// I事件_窗口失去焦点.
func (w *windowBase) I事件_窗口失去焦点1(pFun WM_KILLFOCUS1) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.WM_KILLFOCUS, pFun)
}

// I事件_窗口键盘按键消息.
func (w *windowBase) I事件_窗口键盘按键(pFun WM_KEYDOWN) bool {
	return xc.XWnd_RegEventC(w.I句柄, xcc.WM_KEYDOWN, pFun)
}

// I事件_窗口键盘按键 消息.
func (w *windowBase) I事件_窗口键盘按键1(pFun WM_KEYDOWN1) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.WM_KEYDOWN, pFun)
}

// I事件_窗口鼠标捕获  改变消息.
func (w *windowBase) I事件_窗口鼠标捕获改变(pFun WM_CAPTURECHANGED) bool {
	return xc.XWnd_RegEventC(w.I句柄, xcc.WM_CAPTURECHANGED, pFun)
}

// I事件_窗口鼠标捕获改变  消息.
func (w *windowBase) I事件_窗口鼠标捕获改变1(pFun WM_CAPTURECHANGED1) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.WM_CAPTURECHANGED, pFun)
}

// I事件_窗口设置鼠标光标.
func (w *windowBase) I事件_窗口设置鼠标光标(pFun WM_SETCURSOR) bool {
	return xc.XWnd_RegEventC(w.I句柄, xcc.WM_SETCURSOR, pFun)
}

// I事件_窗口设置鼠标光标.
func (w *windowBase) Event_SETCURSOR1(pFun WM_SETCURSOR1) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.WM_SETCURSOR, pFun)
}

// I事件_窗口字符.
func (w *windowBase) I事件_窗口字符(pFun WM_CHAR) bool {
	return xc.XWnd_RegEventC(w.I句柄, xcc.WM_CHAR, pFun)
}

// I事件_窗口字符.
func (w *windowBase) I事件_窗口字符1(pFun WM_CHAR1) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.WM_CHAR, pFun)
}

// I事件_拖动文件到窗口1.
func (w *windowBase) I事件_拖动文件到窗口(pFun WM_DROPFILES) bool {
	return xc.XWnd_RegEventC(w.I句柄, xcc.WM_DROPFILES, pFun)
}

// I事件_拖动文件到窗口1.
func (w *windowBase) I事件_拖动文件到窗口1(pFun WM_DROPFILES1) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.WM_DROPFILES, pFun)
}

type XWM_MENU_POPUP func(hMenu int, pbHandled *bool) int                                           // 菜单弹出.
type XWM_MENU_POPUP_WND func(hMenu int, pInfo *xc.Menu_PopupWnd_, pbHandled *bool) int             // 菜单弹出窗口.
type XWM_MENU_SELECT func(nID int, pbHandled *bool) int                                            // 菜单选择.
type XWM_MENU_EXIT func(pbHandled *bool) int                                                       // 菜单退出.
type XWM_MENU_DRAW_BACKGROUND func(hDraw int, pInfo *xc.Menu_DrawBackground_, pbHandled *bool) int // 绘制菜单背景, 启用该功能需要调用XMenu_EnableDrawBackground().
type XWM_MENU_DRAWITEM func(hDraw int, pInfo *xc.Menu_DrawItem_, pbHandled *bool) int              // 绘制菜单项事件, 启用该功能需要调用XMenu_EnableDrawItem().

// I事件_菜单弹出.
func (w *windowBase) I事件_菜单弹出(pFun XWM_MENU_POPUP) bool {
	return xc.XWnd_RegEventC(w.I句柄, xcc.XWM_MENU_POPUP, pFun)
}

// I事件_菜单弹出窗口.
func (w *windowBase) I事件_菜单弹出窗口(pFun XWM_MENU_POPUP_WND) bool {
	return xc.XWnd_RegEventC(w.I句柄, xcc.XWM_MENU_POPUP_WND, pFun)
}

// I事件_菜单选择.
func (w *windowBase) I事件_菜单选择(pFun XWM_MENU_SELECT) bool {
	return xc.XWnd_RegEventC(w.I句柄, xcc.XWM_MENU_SELECT, pFun)
}

// I事件_菜单退出.
func (w *windowBase) Event_MENU_EXIT(pFun XWM_MENU_EXIT) bool {
	return xc.XWnd_RegEventC(w.I句柄, xcc.XWM_MENU_EXIT, pFun)
}

// I事件_绘制菜单背景, 启用该功能需要调用XMenu_EnableDrawBackground().
func (w *windowBase) I事件_绘制菜单背景(pFun XWM_MENU_DRAW_BACKGROUND) bool {
	return xc.XWnd_RegEventC(w.I句柄, xcc.XWM_MENU_DRAW_BACKGROUND, pFun)
}

// I事件_绘制菜单项  事件, 启用该功能需要调用XMenu_EnableDrawItem().
func (w *windowBase) I事件_绘制菜单项(pFun XWM_MENU_DRAWITEM) bool {
	return xc.XWnd_RegEventC(w.I句柄, xcc.XWM_MENU_DRAWITEM, pFun)
}

/* type XWM_MENU_POPUP1 func(hWindow int, hMenu int, pbHandled *bool) int                                           // 菜单弹出.
type XWM_MENU_POPUP_WND1 func(hWindow int, hMenu int, pInfo *xc.Menu_PopupWnd_, pbHandled *bool) int             // 菜单弹出窗口.
type XWM_MENU_SELECT1 func(hWindow int, nID int, pbHandled *bool) int                                            // 菜单选择.
type XWM_MENU_EXIT1 func(hWindow int, pbHandled *bool) int                                                       // 菜单退出.
type XWM_MENU_DRAW_BACKGROUND1 func(hWindow int, hDraw int, pInfo *xc.Menu_DrawBackground_, pbHandled *bool) int // 绘制菜单背景, 启用该功能需要调用XMenu_EnableDrawBackground().
type XWM_MENU_DRAWITEM1 func(hWindow int, hDraw int, pInfo *xc.Menu_DrawItem_, pbHandled *bool) int              // 绘制菜单项事件, 启用该功能需要调用XMenu_EnableDrawItem().

// 菜单弹出.
func (w *windowBase) Event_MENU_POPUP1(pFun XWM_MENU_POPUP1) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.XWM_MENU_POPUP, pFun)
}

// 菜单弹出窗口.
func (w *windowBase) Event_MENU_POPUP_WND1(pFun XWM_MENU_POPUP_WND1) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.XWM_MENU_POPUP_WND, pFun)
}

// 菜单选择.
func (w *windowBase) Event_MENU_SELECT1(pFun XWM_MENU_SELECT1) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.XWM_MENU_SELECT, pFun)
}

// 菜单退出.
func (w *windowBase) Event_MENU_EXIT1(pFun XWM_MENU_EXIT1) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.XWM_MENU_EXIT, pFun)
}

// 绘制菜单背景, 启用该功能需要调用XMenu_EnableDrawBackground().
func (w *windowBase) Event_MENU_DRAW_BACKGROUND1(pFun XWM_MENU_DRAW_BACKGROUND1) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.XWM_MENU_DRAW_BACKGROUND, pFun)
}

// 绘制菜单项事件, 启用该功能需要调用XMenu_EnableDrawItem().
func (w *windowBase) Event_MENU_DRAWITEM1(pFun XWM_MENU_DRAWITEM1) bool {
	return xc.XWnd_RegEventC1(w.I句柄, xcc.XWM_MENU_DRAWITEM, pFun)
} */
